#!/bin/bash

# Start both applications
go run main.go 8000 &
go run main.go 8001 &

# Wait for both processes
wait
