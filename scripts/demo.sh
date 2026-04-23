#!/bin/bash
# Usage: chmod +x demo.sh && ./demo.sh

NODES="node1:50051,node2:50051,node3:50051"

echo "DNP Demo"
echo "================"

echo -e "\nStep 1: Start 3 nodes"
docker compose up -d --build node1 node2 node3
sleep 4
read -p "Nodes started. Press Enter to continue..."

echo -e "\nStep 2: Read logs & verify match"
docker compose run --rm client validate --nodes $NODES
read -p "Logs match. Press Enter..."

echo -e "\nStep 3: Submit log 'SET user Alice'"
docker compose run --rm client submit --nodes $NODES --command "SET user Alice"
echo -n "Verify: "; docker compose run --rm client query --addr node1:50051 --key user
read -p "Alice saved. Press Enter..."

echo -e "\n📋 Step 4: Kill the leader"
echo "Current Status:"
docker compose run --rm client metrics --nodes $NODES 2>/dev/null | head -20
echo "Open another terminal and kill the leader:"
echo "   docker compose stop <leader_name>"
read -p "Leader killed. Press Enter to continue..."

echo -e "\nStep 5: Request status (Check leader change)"
docker compose run --rm client metrics --nodes $NODES 2>/dev/null | head -20
read -p "New leader elected. Press Enter..."

echo -e "\nStep 6: Submit more logs (Quorum test)"
docker compose run --rm client submit --nodes $NODES --command "SET user Bob"
read -p "Bob saved. Press Enter..."

echo -e "\nStep 7: Restart dead node & verify follower"
echo "Restart the dead node:"
echo "   docker compose start <dead_node_name>"
read -p "Node restarted. Waiting for catch-up..."
sleep 3
echo -n "Node Role: "; docker compose run --rm client query --addr node1:50051 2>/dev/null | grep -o 'role=[a-z]*'
read -p "Node rejoined as follower. Press Enter..."

echo -e "\nStep 8: Final verification (Logs match)"
docker compose run --rm client compare --nodes $NODES
echo -e "\nDemo Complete"