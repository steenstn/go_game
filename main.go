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
	player     *game.Player
}

var clients = make([]*Client, 0) // TODO: Allocate max players at start

func main() {
	go gameLoop()

	http.HandleFunc("/join", join)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.html")
	})

	http.ListenAndServe(":8080", nil)

}

func gameLoop() {
	for {
		game_update := game.Tick()
		for i := range clients {
			if clients[i].connected == false {
				continue
			}
			aaa, _ := json.Marshal(game_update)
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

func inputLoop(client *Client) {
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
		input := string(msg)
		game.HandleInput(client.player, input)
		println(input)
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

	newPlayer := game.AddPlayer()
	newClient := Client{connection: conn, connected: true, player: newPlayer}
	clients = append(clients, &newClient)
	go inputLoop(&newClient)

	/*
		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()

			if err != nil {
				println("Failed to read message")
				println(err.Error())
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	*/
}
