package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/EdgeFiLabs/PTD/internal/config"
	"github.com/EdgeFiLabs/PTD/internal/websocket"
)

func main() {

	fmt.Println("-PROGRAM STARS-")
	// Load configuration from the credentials file
	cfg := config.LoadConfig()

	// Initialize WebSocket client with the loaded URL
	wsClient, err := websocket.NewClient(cfg.WebSocketURL)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer wsClient.Close()

	wsClient.SubscribeToNewPendingTransactions()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-wsClient.Done:
			return
		case <-interrupt:
			log.Println("Interrupt received, shutting down...")
			wsClient.Close()
			select {
			case <-wsClient.Done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
