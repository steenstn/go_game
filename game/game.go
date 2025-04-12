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

// TODO Handle multipl keys being pressed, handle as numbers
func HandleInput(player *Player, input byte) {

	player.keyPresses.Up = input&1 > 0
	player.keyPresses.Down = input&2 > 0
	player.keyPresses.Left = input&4 > 0
	player.keyPresses.Right = input&8 > 0

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
