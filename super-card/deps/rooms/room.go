package rooms

import (
	"fmt"
	"math/rand/v2"
	"super_card/deps/clients"
)

type Room struct {
	Cards      []map[string]any
	Clients    map[*clients.Client]any
	broadcast  chan any
	register   chan *clients.Client
	unregister chan *clients.Client
}

func NewRoom() *Room {
	return &Room{
		Cards:      []map[string]any{},
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

func (r *Room) SetCards(cards []map[string]any) {
	r.Cards = cards
}

func (r *Room) GetCards() []map[string]any {
	return r.Cards
}

// dealCards shuffles a copy of the deck and splits it evenly across every
// connected client, replacing their previous hand, then hands each client its
// own cards. Any remainder from an uneven split stays undealt. Only Run's
// goroutine may call this — it owns Clients and the clients' Cards.
func (r *Room) dealCards() {
	n := len(r.Clients)
	if n == 0 {
		return
	}
	per := len(r.Cards) / n

	deck := make([]map[string]any, len(r.Cards))
	copy(deck, r.Cards)
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	i := 0
	for c := range r.Clients {
		c.Cards = deck[i*per : (i+1)*per]
		i++
		select {
		case c.Send <- map[string]any{"type": "hand", "cards": c.Cards}:
			// delivered to this client's outbox
		default:
			// client's send buffer is full/stuck — skip the deal for them
		}
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.register:
			r.Clients[client] = true
			r.dealCards()
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
