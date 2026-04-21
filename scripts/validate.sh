#!/usr/bin/env sh
set -eu

COMMAND="${1:-SET validation shell}"
CRASH_COMMAND="${2:-SET after_crash shell}"

docker compose up -d --build node1 node2 node3
sleep 4

docker compose run --rm client validate --nodes "node1:50051,node2:50051,node3:50051" --command "$COMMAND"

docker compose stop node2
docker compose run --rm client submit --nodes "node1:50051,node2:50051,node3:50051" --command "$CRASH_COMMAND"
docker compose start node2

sleep 4
docker compose run --rm client validate --nodes "node1:50051,node2:50051,node3:50051" --command "SET restarted_node caught_up"
docker compose run --rm client compare --nodes "node1:50051,node2:50051,node3:50051"
