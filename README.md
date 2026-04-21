# Project 21: Simplified Log Replication System Architecture

## 1. System Overview

A distributed log replication cluster implementing RAFT consensus. A single leader serializes client writes, replicates entries to followers, and commits entries after receiving majority (n/2 + 1) acknowledgment. Automatic leader election handles node failures under a fail-stop model. Each node runs in an isolated Docker container simulating multiple servers within docker-compose. SQLite handles persistent storage. An in-memory state machine and metrics collector use `sync.Mutex` for thread-safe concurrent access.

---

## 2. Component Architecture

| Component | Responsibility | Exact Tech Mapping |
|-----------|----------------|-------------------|
| **ReplicationEngine** | Coordinates log replication, commit logic, and client request handling. Drives the single-goroutine RAFT event loop. | Go struct with channels for RPC events. `time.Ticker` for heartbeats. Sequential state machine calls. |
| **ElectionManager** | Handles term tracking, vote requests, vote responses, and leader election transitions. | Go struct with randomized timeout logic. Persists `currentTerm` and `votedFor` to SQLite. |
| **Log** | Stores and retrieves RAFT log entries. Provides append, lookup, and truncate operations. | SQLite-backed store via `database/sql`. `PRAGMA journal_mode=WAL`. `PRAGMA synchronous=FULL`. |
| **StateMachine** | Applies committed log entries deterministically. Guarantees replicated state convergence. | `map[string]string` protected by `sync.Mutex`. Sequential apply loop invoked by ReplicationEngine. |
| **gRPC/Protobuf Layer** | Handles inter-node RPCs (`RequestVote`, `AppendEntries`) and client APIs (`SubmitEntry`, `QueryLog`, `GetMetrics`). | `protoc` generates Go structs. `grpc-go` implements server and client. Strict protobuf message schemas. |
| **Metrics Collector** | Records submission, commit, and apply timestamps. Calculates replication lag and apply delay. | Struct with `map[int64]time.Time` and `map[int64]time.Duration`. Protected by `sync.Mutex`. |

---

## 2.1 Node Internal Module Structure

Each container runs one RAFT node instance. The instance contains four modules that interact through Go channels and mutex-protected data structures:

- **Log Module**: Persists log entries to SQLite. Provides `Append(term, command)`, `Get(index)`, and `Truncate(fromIndex)` operations. Flushes to disk before acknowledging writes.
- **StateMachine Module**: Applies committed entries to an in-memory `map[string]string`. Protected by `sync.Mutex`. Updated sequentially by the ReplicationEngine after majority commit.
- **ReplicationEngine Module**: Accepts client writes, broadcasts `AppendEntries` RPCs, tracks `nextIndex[]` and `matchIndex[]`, and advances `commitIndex` on majority ACK. Invokes StateMachine apply and Metrics updates.
- **ElectionManager Module**: Manages election timeouts, sends `RequestVote` RPCs, processes vote responses, and updates `currentTerm` and `votedFor`. Persists voting state to SQLite before responding to RPCs.

Module interaction flow:
```
Client Request
     |
     V
ReplicationEngine --> Log (append entry)
     |
     V
ReplicationEngine --> ElectionManager (check term validity)
     |
     V
ReplicationEngine --> gRPC Layer (broadcast AppendEntries)
     |
     V
Majority ACK received
     |
     V
ReplicationEngine --> StateMachine (apply committed entry)
     |
     V
ReplicationEngine --> Metrics Collector (record timestamps)
     |
     V
Respond to Client
```

---

## 3. Data Flow & Consistency Model

```
Client --[Submit]--> Leader --[Append to SQLite]--> Broadcast AppendEntries --> Followers
                          |                                   ^
                          V                                   |
                  Majority ACK? --Yes--> Update commitIndex ──┘
                          |
                          V
                  Apply to State Machine --> Respond to Client
```

- **Strong Consistency:** All writes route through the leader. Entries commit only after majority replication. Reads query the leader. Followers serve reads only after applying `commitIndex`.
- **Log Matching:** `prevLogIndex` and `prevLogTerm` checks guarantee identical history across nodes. Candidates require up-to-date logs to win elections.
- **Crash Recovery:** Node startup loads `currentTerm`, `votedFor`, and full log from SQLite. Volatile tracking (`nextIndex[]`, `matchIndex[]`) resets to defaults. Leader detects mismatches and backfills missing entries through `AppendEntries` retries.

---

## 4. Deployment Architecture

The cluster runs three identical containers on a dedicated Docker bridge network. Each container exposes a unique host port mapped to the internal gRPC port. Docker DNS resolves peer addresses using container names. Host directories mount to `/data` inside each container to persist SQLite files. Container stop and start commands simulate fail-stop crashes and recoveries. Peer addresses initialize from environment variables at startup.

Docker compose is not a part of the architecture, and is only used to simplify the testing process. In production conditions, each node should be deployed on a separate server.

---

## 5. Validation Checklist Mapping

| Checklist Item | Implementation Steps |
|----------------|----------------------|
| **Submit & compare logs** | Client CLI submits entries via gRPC. `QueryLog` RPC returns SQLite contents. Validation script fetches logs from all nodes. Script diffs `(index, term, command)` sequences. Identical sequences confirm correct replication. |
| **Crash/restart continuity** | Execution stops target container. Execution restarts target container. Node reloads SQLite state. Node rejoins cluster. Leader detects index mismatch and backfills entries. Validation script verifies continuous index sequence and consistent terms. |
| **Measure delay & lag** | Leader records `t_commit` timestamp after majority ACK. Followers record `t_apply` timestamp after state machine update. Client CLI fetches timestamps via `GetMetrics` RPC. CLI computes `replication_lag = leader.commitIndex - follower.lastApplied`. CLI computes `apply_delay = t_apply - t_commit`. CLI outputs numeric results. |

---

## 6. Go Implementation

- **RAFT Goroutine:** One event loop owns all state transitions, timer ticks, and log writes. This design eliminates data races.
- **Channel Routing:** gRPC handlers send requests to unbuffered channels. RAFT loop consumes requests synchronously.
- **State Machine:** RAFT loop acquires `sync.Mutex`. Loop applies committed entries to `map[string]string`. Loop releases mutex. Query handlers acquire mutex for safe concurrent reads.
- **Metrics Collector:** RAFT loop and gRPC handlers update timestamp maps under `sync.Mutex`. `GetMetrics` handler reads data safely and returns JSON output.
- **SQLite Access:** Single writer per node. `WAL` mode enables concurrent reads. All critical writes (`term`, `votedFor`, `log`) flush to disk before sending RPC responses or client acknowledgments.
- **Election Liveness:** Election timeout calculation uses `base + rand.Intn(base)`. This calculation prevents split votes and guarantees leader election progress.
