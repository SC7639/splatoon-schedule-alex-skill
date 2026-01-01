#!/bin/bash
# Script to open dev container in external terminal
# Usage: ./open-external-terminal.sh

# Find the VS Code dev container
CONTAINER=$(docker ps --format '{{.Names}}\t{{.Image}}' | grep 'vsc-' | awk '{print $1}' | head -1)

if [ -z "$CONTAINER" ]; then
    echo "No VS Code dev container found. Make sure the container is running."
    exit 1
fi

echo "Connecting to container: $CONTAINER"

# Try zsh first, then bash, then sh
if docker exec -it "$CONTAINER" which zsh > /dev/null 2>&1; then
    docker exec -it "$CONTAINER" /bin/zsh
elif docker exec -it "$CONTAINER" which bash > /dev/null 2>&1; then
    docker exec -it "$CONTAINER" /bin/bash
else
    docker exec -it "$CONTAINER" /bin/sh
fi

