package game

import (
	"game/queue"
)

type Player struct {
	X          int
	Y          int
	Vx         int
	Vy         int
	keyPresses PlayerKeyPress
}

var players = make([]*Player, 0)
var playerKeyPresses = make([]*PlayerKeyPress, 0)

var InputQueue = queue.NewQueue(10)

func AddPlayer() *Player {
	newPlayer := Player{
		X:          100,
		Y:          100,
		Vx:         0,
		Vy:         0,
		keyPresses: PlayerKeyPress{},
	}
	players = append(players, &newPlayer)
	playerKeyPresses = append(playerKeyPresses, &PlayerKeyPress{})
	return &newPlayer
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func HandleInput(player *Player, input string) {

	if input == "1" {
		player.keyPresses.Up = true
	} else {
		player.keyPresses.Up = false
	}
	if input == "2" {
		player.keyPresses.Down = true
	} else {
		player.keyPresses.Down = false
	}

	if input == "4" {
		player.keyPresses.Left = true
	} else {
		player.keyPresses.Left = false
	}

	if input == "8" {
		player.keyPresses.Right = true
	} else {
		player.keyPresses.Right = false
	}

}

func Tick() []*Player {

	for i := range players {
		if players[i].keyPresses.Right {
			players[i].Vx = 1
		} else if players[i].keyPresses.Left {
			players[i].Vx = -1
		} else {
			players[i].Vx = 0
		}
		if players[i].keyPresses.Up {
			players[i].Vy = -1
		} else if players[i].keyPresses.Down {
			players[i].Vy = 1
		} else {
			players[i].Vy = 0
		}

		players[i].X += players[i].Vx
		players[i].Y += players[i].Vy
	}

	return players
}
