param(
    [string]$Command = "SET validation powershell",
    [string]$CrashCommand = "SET after_crash powershell"
)

$ErrorActionPreference = "Stop"

docker compose up -d --build node1 node2 node3
Start-Sleep -Seconds 4

docker compose run --rm client validate --nodes "node1:50051,node2:50051,node3:50051" --command $Command

docker compose stop node2
docker compose run --rm client submit --nodes "node1:50051,node2:50051,node3:50051" --command $CrashCommand
docker compose start node2

Start-Sleep -Seconds 4
docker compose run --rm client validate --nodes "node1:50051,node2:50051,node3:50051" --command "SET restarted_node caught_up"
docker compose run --rm client compare --nodes "node1:50051,node2:50051,node3:50051"
