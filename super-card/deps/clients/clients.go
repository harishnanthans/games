package clients

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Hub is what a Client needs from the room it belongs to.
// Declared here so clients does not import rooms.
type Hub interface {
	Register(*Client)
	Unregister(*Client)
	Broadcast(any)
}

type Client struct {
	Conn *websocket.Conn
	Send chan any
}

func New(ws *websocket.Conn) *Client {
	return &Client{
		Conn: ws,
		Send: make(chan any, 256),
	}
}

func (c *Client) ReadMessage(h Hub) {
	defer func() {
		h.Unregister(c)
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read message fail")
			return
		}

		h.Broadcast(msg)
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
		c.Conn.Close()
	}()

	for msg := range c.Send {
		byte, _ := json.Marshal(msg)
		err := c.Conn.WriteMessage(websocket.TextMessage, byte)
		if err != nil {
			log.Println("write message fail")
			return
		}
	}
}
