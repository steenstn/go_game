package main

import (
	"encoding/json"
	"fmt"
	"game/game"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

/*

https://github.com/0xFA11/MultiplayerNetworkingResources
https://gafferongames.com/

CLient reconciliation
https://www.gabrielgambetta.com/client-server-game-architecture.html


TODO:
- Gör connection loop som ber om saker en i taget, använd enumet
- CLient side prediction /reconciliation
	- Try and reset the counter at some point.
- Store entities locally?
- Send level and other information on start
  - Connecting state on client to get all the required data up front
- run length encoding vs bits
- Remove player if they disconnect
- chat?
- Nicer camera movement
- Refactor creating of entities, don't want to manually increase numEntities etc
- Spider spider.html

BUG
- Player can drop outside of level - What to do if user fetches outside array?
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
	Type    byte
	SubType byte
	Msg     []byte
}

type ServerFullError struct{}

func (m *ServerFullError) Error() string {
	return "Server is full"
}

type EntityType byte

// #export "enums.js"
const (
	Player EntityType = 0
	Fly    EntityType = 1
	Spider EntityType = 2
)

type MessageType byte

// #export "enums.js"
const (
	GameUpdateMessage     MessageType = 0
	SetupMessage          MessageType = 1
	PlayerPositionMessage MessageType = 2
)

type SetupMessageSubType byte

// #export "enums.js"
const (
	Level        SetupMessageSubType = 0
	PlayerSprite SetupMessageSubType = 1
	LevelTileset SetupMessageSubType = 2
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

type Input struct {
	Num   int
	Input byte
}

var clients = make([]*Client, 10)

var playerSprite, _ = os.ReadFile("playersprite.png")
var levelTileset, _ = os.ReadFile("tileset.png")

func main() {

	javascriptExport("main.go")
	javascriptParseInclude("client/client.html")

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

type lolDaMessage struct {
	Positions []game.Position
	Types     []int
}

func gameLoop() {
	for {
		game.Tick()
		game_update := game.PositionRegistry

		position_array := make([]game.Position, 0, len(game_update))
		types_array := make([]int, 0, len(game_update))
		for i := game.EntityId(0); i < game.NumEntities; i++ {
			position_array = append(position_array, *game.PositionRegistry[i])
			types_array = append(types_array, game.EntityTypeRegistry[i])
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
			daMessage := lolDaMessage{Positions: position_array, Types: types_array}
			positions, _ := json.Marshal(daMessage)
			gameUpdateMessage, _ := json.Marshal(JsonMessage{Type: byte(GameUpdateMessage), Msg: positions})
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
		// TODO: Validate
		/*if len(msg) > 1 {
			println("Input message is too long, dropping")
			println(string(msg))
			println("--")
			continue
		}*/
		parsedInput := Input{}
		json.Unmarshal(msg, &parsedInput)
		//input := msg[0]
		fmt.Printf("input: %b\n", parsedInput)
		println(parsedInput.Num)
		println(parsedInput.Input)
		//time.Sleep(300 * time.Millisecond)

		game.HandleInput(parsedInput.Input, entityId)
	}
}

func join(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s connected\n", request.Host)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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
	sendSetupMessage(clients[freeSlot], Level, currentLevelMessage)

	playerSpriteMessage, _ := json.Marshal(playerSprite)
	sendSetupMessage(clients[freeSlot], PlayerSprite, playerSpriteMessage)

	tilesetMessage, _ := json.Marshal(levelTileset)
	sendSetupMessage(clients[freeSlot], LevelTileset, tilesetMessage)

	go inputLoop(clients[freeSlot], entityId)
}

func sendSetupMessage(client *Client, subtype SetupMessageSubType, message []byte) error {
	jsonMessage, marshallError := json.Marshal(JsonMessage{Type: byte(SetupMessage), SubType: byte(subtype), Msg: message})
	if marshallError != nil {
		return marshallError
	}

	return client.connection.WriteMessage(websocket.TextMessage, jsonMessage)
}

/*
Loop to download all assets
*/
func connectionLoop(client *Client) {
	println("Starting connection loop")

	for {
		if client.status != Connecting {
			break
		}

		_, msg, err := client.connection.ReadMessage()
		if err != nil {
			println("Failed to read input")
			client.status = Disconnected
			break
		}
		println(msg)
	}
}

func findFreeSlot(clientArray []*Client) (int, error) {
	for i := range clientArray {
		if clientArray[i] == nil || clientArray[i].status == Disconnected {
			return i, nil
		}
	}

	return -1, &ServerFullError{}
}
