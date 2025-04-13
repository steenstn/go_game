package main

import (
	"encoding/json"
	"fmt"
	"game/game"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

/*
floats as ints
1.2 = 120
0.01 = 1

*/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type JsonMessage struct {
	Type byte
	Msg  []byte
}

type MessageType byte

const (
	JoinMessage MessageType = 0
)

type Client struct {
	connection *websocket.Conn
	connected  bool
}

var clients = make([]*Client, 0) // TODO: Allocate max players at start

func main() {

	go gameLoop()

	http.HandleFunc("/join", join)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.html")
	})

	println("Listening")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
	}

}

func gameLoop() {
	for {
		game_update := game.Tick()

		position_array := make([]game.Position, 0, len(*game_update))

		for _, value := range *game_update {
			position_array = append(position_array, *value)
		}

		for i := range clients {
			if clients[i].connected == false {
				continue
			}
			aaa, _ := json.Marshal(position_array)
			gameUpdateMessage, _ := json.Marshal(JsonMessage{
				Type: byte(JoinMessage), Msg: aaa})
			err := clients[i].connection.
				WriteMessage(websocket.TextMessage, gameUpdateMessage)
			if err != nil {
				println("Failed to write message")
				println(err.Error())
				println("Marking player as disconnected")
				clients[i].connected = false
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func inputLoop(client *Client, entityId game.EntityId) {
	println("Starting inputLoop")

	for {
		if client.connected == false {
			break
		}

		_, msg, err := client.connection.ReadMessage()
		if err != nil {
			println("Input reading failed, player dropped")
			client.connected = false
			break
		}

		fmt.Printf("input: %b\n", msg)
		if len(msg) > 1 {
			println("Input message is too long, dropping")
			println(string(msg))
			println("--")
			continue
		}
		input := msg[0]
		game.HandleInput(input, entityId)
	}
}

func join(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s connected\n", request.Host)
	conn, upgradeError := upgrader.Upgrade(responseWriter, request, nil)

	if upgradeError != nil {
		println("Failed to upgrade")
		println(upgradeError.Error())
		return
	}

	entityId := game.AddPlayer()
	newClient := Client{connection: conn, connected: true}
	clients = append(clients, &newClient)
	go inputLoop(&newClient, entityId)

}
