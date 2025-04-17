package game

import "math"

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
				velocity.Vy -= 2
			}
		}
	}

}

// TODO: Ska det inte vara en pekare här? Eller är det redan det?
func getArrayIndex(levelWidth int, tileWidth int, x float64, y float64) int {
	xPosition := math.Floor(x / float64(tileWidth))
	yPosition := float64(levelWidth) * math.Floor(y/float64(tileWidth))

	return int(xPosition + yPosition)
}

func MoveStuff(level *Level, positionRegistry map[EntityId]*Position, velocityRegistry map[EntityId]*Velocity, forceRegistry map[EntityId]*Force) {

	for e := EntityId(0); e < NumEntities; e++ {
		position, pOk := positionRegistry[e]
		velocity, vOk := velocityRegistry[e]
		force, fOk := forceRegistry[e]

		if pOk && vOk && fOk {
			velocity.Vx += force.X
			velocity.Vy += force.Y

			oldX := position.X
			oldY := position.Y
			position.X += velocity.Vx
			position.Y += velocity.Vy

			if level.Data[getArrayIndex(level.Width, 50, position.X, position.Y)] == 1 {
				position.X = oldX
				position.Y = oldY
				velocity.Vx = 0
				velocity.Vy = 0

			}
		}
	}

}
