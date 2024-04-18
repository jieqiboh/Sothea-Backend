#!/bin/bash

docker compose down

## Define a function to handle cleanup
#cleanup() {
#    echo "Stopping containers and performing cleanup..."
#    docker-compose down
#    exit 1
#}
#
## Trap SIGINT signal (Ctrl+C) and execute the cleanup function
#trap cleanup SIGINT
#
## Wait indefinitely, keeping the script running
#while true; do
#    sleep 1
#done