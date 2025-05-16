// internal/config/config.go
package config

import (
    "bufio"
    "log"
    "os"
    "strings"
)

type Config struct {
    WebSocketURL string
}

func LoadConfig() *Config {
    file, err := os.Open("credentials")
    if err != nil {
        log.Fatalf("Error opening credentials file: %v", err)
    }
    defer file.Close()

    config := &Config{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue // Skip empty lines and comments
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            log.Fatalf("Invalid line in credentials file: %s", line)
        }

        key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

        switch key {
        case "WEBSOCKET_URL":
            config.WebSocketURL = value
        default:
            log.Printf("Unknown key in credentials file: %s", key)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading credentials file: %v", err)
    }

    if config.WebSocketURL == "" {
        log.Fatal("WEBSOCKET_URL is required in the credentials file")
    }

    return config
}
