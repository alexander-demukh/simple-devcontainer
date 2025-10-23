# Simple Dev Container

A Go application in a dev container with a `/ping` endpoint.

## Running the Application

### Single Instance

Run on port 8000 (default):
```bash
devcontainer exec --workspace-folder . go run main.go
```

Run on a custom port:
```bash
devcontainer exec --workspace-folder . go run main.go 8001
```

### Both Services

Run both services on ports 8000 and 8001 with a single command:
```bash
devcontainer exec --workspace-folder . ./start-both.sh
```

Or from inside the container:
```bash
./start-both.sh
```

The application will:
- Start an HTTP server on the specified port
- Respond to `/ping` with "pong"
- Print "App {port}: I'm working" every 5 seconds to stdout

## Testing

From your local machine:
```bash
curl http://localhost:8000/ping
curl http://localhost:8001/ping
```

From inside the container:
```bash
devcontainer exec --workspace-folder . curl http://localhost:8000/ping
devcontainer exec --workspace-folder . curl http://localhost:8001/ping
```

## Viewing Logs

When running with `devcontainer exec`, logs appear directly in your terminal. You can also view container logs with:

```bash
docker logs -f <container-id>
```
