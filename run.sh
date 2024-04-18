#!/bin/bash

# Start docker containers in detached mode, running in the background
docker-compose up -d

# Define a function to handle cleanup
cleanup() {
    echo "Stopping containers and performing cleanup..."
    docker-compose down
    exit 1
}

# Trap SIGINT signal (Ctrl+C) and execute the cleanup function
trap cleanup SIGINT

# Wait indefinitely, keeping the script running
while true; do
    sleep 1
done