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
- CLient side prediction /reconciliation
- Send level
- run length encoding vs bits
- Remove player if they disconnect
- chat?
- Nicer camera movement
- Spider spider.html

BUG
- Player can drop outside of level
*/

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

type ServerFullError struct{}

func (m *ServerFullError) Error() string {
	return "Server is full"
}

const (
	JoinMessage           MessageType = 0
	LevelMessage          MessageType = 1
	PlayerPositionMessage MessageType = 2
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
	entityId   game.EntityId
}

var clients = make([]*Client, 10)

func main() {

	includeStuff("client/client.html")
	game.InitGame()
	go gameLoop()

	http.HandleFunc("/join", join)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "out/client/client.html")
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

		position_array := make([]game.Position, 0, len(game_update))

		for _, value := range game_update {
			position_array = append(position_array, *value)
		}

		for i := range clients {
			if clients[i] == nil || clients[i].status != Connected {
				continue
			}

			// Send player position
			playerPosition, _ := json.Marshal(game_update[clients[i].entityId])

			playerPositionMessage, _ := json.Marshal(JsonMessage{Type: byte(PlayerPositionMessage), Msg: playerPosition})
			clients[i].connection.WriteMessage(websocket.TextMessage, playerPositionMessage)

			// Send all positions
			positions, _ := json.Marshal(position_array)
			gameUpdateMessage, _ := json.Marshal(JsonMessage{Type: byte(JoinMessage), Msg: positions})
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

		//fmt.Printf("input: %b\n", msg)
		if len(msg) > 1 {
			println("Input message is too long, dropping")
			println(string(msg))
			println("--")
			continue
		}
		input := msg[0]
		time.Sleep(300 * time.Millisecond)

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

	freeSlot, err := findFreeSlot(clients)
	if err != nil {
		println(err.Error())
		return
	}

	entityId := game.AddPlayer()
	clients[freeSlot] = &Client{connection: conn, status: Connected, entityId: entityId}

	currentLevelMessage, _ := json.Marshal(game.CurrentLevel.Data)
	gameUpdateMessage, _ := json.Marshal(JsonMessage{Type: byte(LevelMessage), Msg: currentLevelMessage})
	err = clients[freeSlot].connection.WriteMessage(websocket.TextMessage, gameUpdateMessage)
	if err != nil {
		println(err.Error())
		return
	}

	go inputLoop(clients[freeSlot], entityId)
}

func findFreeSlot(clientArray []*Client) (int, error) {
	for i := range clientArray {
		if clientArray[i] == nil || clientArray[i].status == Disconnected {
			return i, nil
		}
	}

	return -1, &ServerFullError{}
}
