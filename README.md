
# PTD

**PTD (Placeholder Trading Daemon)** is a modular, extensible Golang project designed for managing websocket connections, configuration layers, and logging functionality in a clean, scalable architecture.

## 📁 Project Structure

```
.
├── cmd/                  # Entry point (main.go)
├── internal/             # Internal application logic
│   ├── config/           # Configuration loading (config.go)
│   └── websocket/        # WebSocket subscriptions and handlers
├── pkg/                  # Shared utilities (e.g., logger)
├── go.mod                # Go module file
├── go.sum                # Go dependencies checksum
```

## 🚀 Getting Started

### Requirements
- Go 1.20+
- WebSocket API endpoint (configurable)

### Build and Run

```bash
go build -o PTD ./cmd
./PTD
```

## 🧠 Components

### `/cmd/main.go`
Program entry point. Initializes config, logger, and starts websocket service.

### `/internal/config/`
Handles configuration loading from files, environment, or flags.

### `/internal/websocket/`
Contains websocket connection logic, including subscription lifecycle.

### `/pkg/utils/logger.go`
Centralized logging using standard Go log or external logger.

## 🛡️ Security

Make sure the following files are excluded from your repository:

- `PTD` (binary)
- `credentials` (if containing secrets)

These are already included in `.gitignore`.

## 📜 License

MIT License © 2025 Arda Çimen
