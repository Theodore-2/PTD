// internal/websocket/subscription.go
package websocket

import (
    "log"
)

type SubscriptionRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
    ID      int           `json:"id"`
}

func (c *Client) SubscribeToNewPendingTransactions() {
    log.Println("[+] Connected to node via WebSocket for mempool")
    log.Println("[+] Waiting for sync ...")

    request := SubscriptionRequest{
        JSONRPC: "2.0",
        Method:  "eth_subscribe",
        Params:  []interface{}{"newPendingTransactions"},
        ID:      1,
    }

    err := c.Conn.WriteJSON(request)
    if err != nil {
        log.Println("write:", err)
        return
    }

    go func() {
        for {
            _, message, err := c.Conn.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("Received: %s", message)
        }
    }()
}
