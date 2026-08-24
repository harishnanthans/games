package rooms

import (
	"fmt"
	"super_card/deps/clients"
)

type Room struct {
	Cards      []int
	Clients    map[*clients.Client]any
	broadcast  chan any
	register   chan *clients.Client
	unregister chan *clients.Client
}

func NewRoom() *Room {
	return &Room{
		Cards:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Clients:    make(map[*clients.Client]any),
		broadcast:  make(chan any),
		register:   make(chan *clients.Client),
		unregister: make(chan *clients.Client),
	}
}

// Register, Unregister and Broadcast satisfy clients.Hub, so callers hand off
// work to Run's goroutine without knowing channels are involved.
func (r *Room) Register(c *clients.Client) { r.register <- c }

func (r *Room) Unregister(c *clients.Client) { r.unregister <- c }

func (r *Room) Broadcast(msg any) { r.broadcast <- msg }

func (r *Room) ClientsLength() int {
	return len(r.Clients)
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.register:
			r.Clients[client] = true
		case client := <-r.unregister:
			// map membership is the guard against closing Send twice
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.Send)
			}
		case msg := <-r.broadcast:
			fmt.Println("broadcast msg", msg)
			for client := range r.Clients {
				select {
				case client.Send <- msg:
					// delivered to this client's outbox
				default:
					// client's send buffer is full/stuck — drop them
					delete(r.Clients, client)
					close(client.Send)
				}
			}
		}
	}
}
