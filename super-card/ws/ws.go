package ws

import (
	"log"
	"net/http"
	"super_card/deps/clients"
	"super_card/deps/rooms"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request, room *rooms.Room, one map[string]any) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("webhook upgrade error", err)
		return
	}

	log.Println("client connected")

	client := clients.New(conn)
	room.Register(client)

	room.Broadcast(one)

	go client.WriteMessage()

	// blocks until the client disconnects; unregisters via its own defer
	client.ReadMessage(room)

	log.Print("client disconnected")
}
