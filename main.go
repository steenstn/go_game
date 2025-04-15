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

https://github.com/0xFA11/MultiplayerNetworkingResources
https://gafferongames.com/


TODO:
- Send level
- run length encoding vs bits
*/

/*
floats as ints
1.2 = 120
0.01 = 1

*/

var level = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 1, 1, 0, 0, 1, 0, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type JsonMessage struct {
	Type byte
	Msg  []byte
}

type MessageType byte

type ServerFullError struct{}

func (m *ServerFullError) Error() string {
	return "Server is full"
}

const (
	JoinMessage  MessageType = 0
	LevelMessage MessageType = 1
)

type ClientStatus byte

const (
	Disconnected ClientStatus = 0
	Connecting   ClientStatus = 1
	Connected    ClientStatus = 2
)

type Client struct {
	connection *websocket.Conn
	status     ClientStatus
}

var clients = make([]*Client, 10)

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
			if clients[i] == nil || clients[i].status != Connected {
				continue
			}

			aaa, _ := json.Marshal(position_array)
			gameUpdateMessage, _ := json.Marshal(JsonMessage{
				Type: byte(JoinMessage), Msg: aaa})
			err := clients[i].connection.WriteMessage(websocket.TextMessage, gameUpdateMessage)

			if err != nil {
				println("Failed to write message")
				println(err.Error())
				println("Marking player as disconnected")
				clients[i].status = Disconnected
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func inputLoop(client *Client, entityId game.EntityId) {
	println("Starting inputLoop")

	for {
		if client.status != Connected {
			break
		}

		_, msg, err := client.connection.ReadMessage()
		if err != nil {
			println("Input reading failed, player dropped")
			client.status = Disconnected
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
	freeSlot, err := findFreeSlot()
	if err != nil {
		println(err.Error())
		return
	}

	// TODO: Send level and stuff

	clients[freeSlot] = &Client{connection: conn, status: Connected}

	aaa, _ := json.Marshal(level)
	gameUpdateMessage, _ := json.Marshal(JsonMessage{Type: byte(LevelMessage), Msg: aaa})
	err = clients[freeSlot].connection.WriteMessage(websocket.TextMessage, gameUpdateMessage)
	if err != nil {
		println(err.Error())
		return
	}

	go inputLoop(clients[freeSlot], entityId)
}

func findFreeSlot() (int, error) {
	for i := range clients {
		if clients[i] == nil || clients[i].status == Disconnected {
			return i, nil
		}
	}

	return -1, &ServerFullError{}
}
