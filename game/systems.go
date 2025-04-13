package game

import "fmt"

func HandleDaInput(playerinputRegistry map[EntityId]*PlayerKeyPress, velocityRegistry map[EntityId]*Velocity) {

	for e := EntityId(0); e < NumEntities; e++ {

		player, pOk := playerinputRegistry[e]
		velocity, vOk := velocityRegistry[e]

		if pOk && vOk {
			if player.Right {
				velocity.Vx = playerSpeed
			} else if player.Left {
				velocity.Vx = -playerSpeed
			} else {
				velocity.Vx = 0
			}
			if player.Up {
				velocity.Vy = -playerSpeed
			} else if player.Down {
				velocity.Vy = playerSpeed
			} else {
				velocity.Vy = 0
			}
		}
	}

}

func MoveStuff(positionRegistry map[EntityId]*Position, velocityRegistry map[EntityId]*Velocity) {

	for e := EntityId(0); e < NumEntities; e++ {
		position, pOk := positionRegistry[e]
		velocity, vOk := velocityRegistry[e]

		if pOk && vOk {
			position.X += velocity.Vx
			position.Y += velocity.Vy
			fmt.Printf("pos: %d", position.X)
		}
	}
	/*
		for i := range players {
			if players[i].keyPresses.Right {
				players[i].Vx = playerSpeed
			} else if players[i].keyPresses.Left {
				players[i].Vx = -playerSpeed
			} else {
				players[i].Vx = 0
			}
			if players[i].keyPresses.Up {
				players[i].Vy = -playerSpeed
			} else if players[i].keyPresses.Down {
				players[i].Vy = playerSpeed
			} else {
				players[i].Vy = 0
			}

			players[i].X += players[i].Vx
			players[i].Y += players[i].Vy
		}
	*/
}
