
# PTD (Pending Transaction Daemon)

**PTD** is a WebSocket-based transaction monitoring daemon developed by EdgeFiLabs. It connects to a blockchain node or data provider via WebSocket and listens for new pending transactions in real time.

## 🛠️ Features

- Connects to a configurable WebSocket endpoint
- Subscribes to "newPendingTransactions" events
- Gracefully handles shutdown signals (e.g., Ctrl+C)
- Designed for modular extension and integration

## 📂 Directory Structure

```
.
├── cmd/                  # Application entry point (main.go)
├── internal/
│   ├── config/           # Configuration loading (credentials, endpoints)
│   └── websocket/        # WebSocket client logic and subscription handling
├── pkg/                  # Reusable utilities
├── go.mod                # Module definition
├── go.sum                # Dependency hashes
├── .gitignore
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or newer
- WebSocket endpoint for pending transaction stream

### Installation

```bash
git clone https://github.com/EdgeFiLabs/PTD.git
cd PTD
go build -o PTD ./cmd
./PTD
```

## ⚙️ Configuration

Edit the `credentials` file or modify `internal/config/config.go` to update your WebSocket endpoint or access credentials.

## 🧪 Example Run Output

```
-PROGRAM STARTS-
Listening to new pending transactions...
```

## 📦 Used In

This service is used in monitoring systems, transaction sniffers, and custom mempool analysis tools.

## 🧾 License

Private - © 2025 EdgeFiLabs. All rights reserved.
