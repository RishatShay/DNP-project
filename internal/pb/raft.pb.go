package pb

import proto "github.com/golang/protobuf/proto"

const _ = proto.ProtoPackageIsVersion4

type LogEntry struct {
	Index   int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Term    int64  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{0}
}

type RequestVoteRequest struct {
	Term         int64  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	LastLogIndex int64  `protobuf:"varint,3,opt,name=last_log_index,json=lastLogIndex,proto3" json:"last_log_index,omitempty"`
	LastLogTerm  int64  `protobuf:"varint,4,opt,name=last_log_term,json=lastLogTerm,proto3" json:"last_log_term,omitempty"`
}

func (m *RequestVoteRequest) Reset()         { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()    {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{1}
}

type RequestVoteResponse struct {
	Term        int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
}

func (m *RequestVoteResponse) Reset()         { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()    {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{2}
}

type AppendEntriesRequest struct {
	Term         int64       `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     string      `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	PrevLogIndex int64       `protobuf:"varint,3,opt,name=prev_log_index,json=prevLogIndex,proto3" json:"prev_log_index,omitempty"`
	PrevLogTerm  int64       `protobuf:"varint,4,opt,name=prev_log_term,json=prevLogTerm,proto3" json:"prev_log_term,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit int64       `protobuf:"varint,6,opt,name=leader_commit,json=leaderCommit,proto3" json:"leader_commit,omitempty"`
}

func (m *AppendEntriesRequest) Reset()         { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()    {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{3}
}

type AppendEntriesResponse struct {
	Term       int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success    bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	MatchIndex int64 `protobuf:"varint,3,opt,name=match_index,json=matchIndex,proto3" json:"match_index,omitempty"`
}

func (m *AppendEntriesResponse) Reset()         { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()    {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{4}
}

type SubmitEntryRequest struct {
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (m *SubmitEntryRequest) Reset()         { *m = SubmitEntryRequest{} }
func (m *SubmitEntryRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitEntryRequest) ProtoMessage()    {}
func (*SubmitEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{5}
}

type SubmitEntryResponse struct {
	Ok       bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	LeaderId string `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Index    int64  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Term     int64  `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	Message  string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *SubmitEntryResponse) Reset()         { *m = SubmitEntryResponse{} }
func (m *SubmitEntryResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitEntryResponse) ProtoMessage()    {}
func (*SubmitEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{6}
}

type QueryLogRequest struct{}

func (m *QueryLogRequest) Reset()         { *m = QueryLogRequest{} }
func (m *QueryLogRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLogRequest) ProtoMessage()    {}
func (*QueryLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{7}
}

type QueryLogResponse struct {
	NodeId       string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Role         string            `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	LeaderId     string            `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	CurrentTerm  int64             `protobuf:"varint,4,opt,name=current_term,json=currentTerm,proto3" json:"current_term,omitempty"`
	CommitIndex  int64             `protobuf:"varint,5,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
	LastApplied  int64             `protobuf:"varint,6,opt,name=last_applied,json=lastApplied,proto3" json:"last_applied,omitempty"`
	Entries      []*LogEntry       `protobuf:"bytes,7,rep,name=entries,proto3" json:"entries,omitempty"`
	State        map[string]string `protobuf:"bytes,8,rep,name=state,proto3" json:"state,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *QueryLogResponse) Reset()         { *m = QueryLogResponse{} }
func (m *QueryLogResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLogResponse) ProtoMessage()    {}
func (*QueryLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{8}
}

type GetMetricsRequest struct{}

func (m *GetMetricsRequest) Reset()         { *m = GetMetricsRequest{} }
func (m *GetMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetricsRequest) ProtoMessage()    {}
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{9}
}

type EntryMetric struct {
	Index               int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	SubmittedAtUnixNano int64 `protobuf:"varint,2,opt,name=submitted_at_unix_nano,json=submittedAtUnixNano,proto3" json:"submitted_at_unix_nano,omitempty"`
	CommittedAtUnixNano int64 `protobuf:"varint,3,opt,name=committed_at_unix_nano,json=committedAtUnixNano,proto3" json:"committed_at_unix_nano,omitempty"`
	AppliedAtUnixNano   int64 `protobuf:"varint,4,opt,name=applied_at_unix_nano,json=appliedAtUnixNano,proto3" json:"applied_at_unix_nano,omitempty"`
	ApplyDelayNanos     int64 `protobuf:"varint,5,opt,name=apply_delay_nanos,json=applyDelayNanos,proto3" json:"apply_delay_nanos,omitempty"`
}

func (m *EntryMetric) Reset()         { *m = EntryMetric{} }
func (m *EntryMetric) String() string { return proto.CompactTextString(m) }
func (*EntryMetric) ProtoMessage()    {}
func (*EntryMetric) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{10}
}

type GetMetricsResponse struct {
	NodeId         string         `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Role           string         `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	CurrentTerm    int64          `protobuf:"varint,3,opt,name=current_term,json=currentTerm,proto3" json:"current_term,omitempty"`
	CommitIndex    int64          `protobuf:"varint,4,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
	LastApplied    int64          `protobuf:"varint,5,opt,name=last_applied,json=lastApplied,proto3" json:"last_applied,omitempty"`
	ReplicationLag int64          `protobuf:"varint,6,opt,name=replication_lag,json=replicationLag,proto3" json:"replication_lag,omitempty"`
	Metrics        []*EntryMetric `protobuf:"bytes,7,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (m *GetMetricsResponse) Reset()         { *m = GetMetricsResponse{} }
func (m *GetMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricsResponse) ProtoMessage()    {}
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRaft, []int{11}
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "raft.LogEntry")
	proto.RegisterType((*RequestVoteRequest)(nil), "raft.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "raft.RequestVoteResponse")
	proto.RegisterType((*AppendEntriesRequest)(nil), "raft.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "raft.AppendEntriesResponse")
	proto.RegisterType((*SubmitEntryRequest)(nil), "raft.SubmitEntryRequest")
	proto.RegisterType((*SubmitEntryResponse)(nil), "raft.SubmitEntryResponse")
	proto.RegisterType((*QueryLogRequest)(nil), "raft.QueryLogRequest")
	proto.RegisterType((*QueryLogResponse)(nil), "raft.QueryLogResponse")
	proto.RegisterMapType((map[string]string)(nil), "raft.QueryLogResponse.StateEntry")
	proto.RegisterType((*GetMetricsRequest)(nil), "raft.GetMetricsRequest")
	proto.RegisterType((*EntryMetric)(nil), "raft.EntryMetric")
	proto.RegisterType((*GetMetricsResponse)(nil), "raft.GetMetricsResponse")
}

var fileDescriptorRaft = []byte{}
