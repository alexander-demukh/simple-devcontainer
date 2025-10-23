#!/bin/bash

# Wait for workspace to be mounted
while [ ! -f /workspace/main.go ]; do
    sleep 1
done

# Download dependencies
cd /workspace
go mod download

# Start the app
exec go run main.go
