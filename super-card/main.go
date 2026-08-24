package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
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
	http.HandleFunc("/", ping)

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
	go room.Run()

	room.ClientsLength()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		rn := rand.IntN(20)
		one := data.SuperStars[rn]
		// fmt.Println(one)
		ws.WsHandler(w, r, room, one)
	})

	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
