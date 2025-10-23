# Simple Dev Container

A Go application in a dev container with a `/ping` endpoint.

## Running the Application

After starting the dev container, run the application with:

```bash
devcontainer exec --workspace-folder . go run main.go
```

Or from inside the container:

```bash
go run main.go
```

The application will:
- Start an HTTP server on port 8000
- Respond to `/ping` with "pong"
- Print "I'm working" every 5 seconds to stdout

## Testing

From your local machine:
```bash
curl http://localhost:8000/ping
```

From inside the container:
```bash
devcontainer exec --workspace-folder . curl http://localhost:8000/ping
```

## Viewing Logs

When running with `devcontainer exec`, logs appear directly in your terminal. You can also view container logs with:

```bash
docker logs -f <container-id>
```
