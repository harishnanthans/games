package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"super_card/deps/rooms"
	"super_card/ws"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

type data struct {
	SuperStars []map[string]any `json:"superstars"`
}

func main() {
	http.HandleFunc("/ping", ping)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	fileByte, err := os.ReadFile("./assets/data.json")
	if err != nil {
		log.Fatal(err)
	}

	var data data
	err = json.Unmarshal(fileByte, &data)
	if err != nil {
		log.Fatal("error while parse the file", err)
	}

	room := rooms.NewRoom()
	room.SetCards(data.SuperStars)

	go room.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.WsHandler(w, r, room)
	})

	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
