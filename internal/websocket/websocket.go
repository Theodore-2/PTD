// internal/websocket/websocket.go
package websocket

import (
    "github.com/gorilla/websocket"
)

type Client struct {
    Conn *websocket.Conn
    Done chan struct{}
}

func NewClient(url string) (*Client, error) {
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        return nil, err
    }
    return &Client{
        Conn: conn,
        Done: make(chan struct{}),
    }, nil
}

func (c *Client) Close() {
    c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
    close(c.Done)
}
