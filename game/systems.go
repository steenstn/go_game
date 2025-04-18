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

func getArrayIndex(levelWidth int, tileWidth int, x float64, y float64) int {
	xPosition := math.Floor(x / float64(tileWidth))
	yPosition := float64(levelWidth) * math.Floor(y/float64(tileWidth))

	return int(xPosition + yPosition)
}

func MoveStuff(level *Level, positionRegistry map[EntityId]*Position, velocityRegistry map[EntityId]*Velocity, gravityRegistry map[EntityId]*Force) {

	tileWidth := 50
	for e := EntityId(0); e < NumEntities; e++ {
		position, pOk := positionRegistry[e]
		velocity, vOk := velocityRegistry[e]
		force, fOk := gravityRegistry[e]

		if pOk && vOk && fOk {

			velocity.Vy += force.Y

			oldX := position.X
			oldY := position.Y
			position.X += velocity.Vx
			position.Y += velocity.Vy

			// Check collision down
			if level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y+5)] == 1 {
				position.Y = oldY
				velocity.Vy = 0
			}

			// Check collision right and left
			if velocity.Vx > 0 && level.Data[getArrayIndex(level.Width, tileWidth, position.X+5, position.Y)] == 1 {
				position.X = oldX
				velocity.Vx = 0
			} else if velocity.Vx < 0 && level.Data[getArrayIndex(level.Width, tileWidth, position.X-5, position.Y)] == 1 {
				position.X = oldX
				velocity.Vx = 0
			}

		}
	}

}
