package game

import (
	"math"
	"math/rand/v2"
)

const TileWidth = 50

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

func distanceSquared(p1 Position, p2 Position) float64 {
	return (p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y)
}

func HandleCircleMovement(circleMovement map[EntityId]*CircleMovement, velocityRegistry map[EntityId]*Velocity) {
	for e := EntityId(0); e < NumEntities; e++ {
		thing, thingOk := CircleMovementRegistry[e]
		if thingOk {

			thing.Timer--

			if thing.Timer <= 0 {
				thing.Direction = thing.Direction * -1
				thing.Timer = rand.IntN(50) + 5
				velocity, _ := velocityRegistry[e]
				velocity.Vx = float64(2 * thing.Direction)
			}

		}
	}

}

func HandleAI(aiRegistry map[EntityId]*AIMovement, velocityRegistry map[EntityId]*Velocity, players []EntityId, positionRegistry map[EntityId]*Position) {
	for e := EntityId(0); e < NumEntities; e++ {
		ai, aiOk := aiRegistry[e]
		if aiOk {

			var dangerPosition = Position{}
			aiPosition, _ := positionRegistry[e]
			for player := range players {
				playerPosition, _ := positionRegistry[players[player]]
				if distanceSquared(*playerPosition, *aiPosition) < 8000 {
					dangerPosition = *playerPosition
					ai.State = 1
				} else {
					ai.State = 0
				}
			}
			if ai.State == 1 {

				velocity, _ := velocityRegistry[e]
				escapeAngle := math.Atan2(aiPosition.Y-dangerPosition.Y, aiPosition.X-dangerPosition.X)

				velocity.Vx = 7 * math.Cos(escapeAngle)
				velocity.Vy = 7 * math.Sin(escapeAngle)
			} else {

				ai.Timer--
				if ai.CurrentAngle > ai.TargetAngle {
					ai.CurrentAngle -= 0.5
				} else if ai.CurrentAngle < ai.TargetAngle {
					ai.CurrentAngle += 0.5
				}
				velocity, _ := velocityRegistry[e]
				velocity.Vx = 2 * math.Cos(ai.CurrentAngle)
				velocity.Vy = 2 * math.Sin(ai.CurrentAngle)

				if ai.Timer < 0 {
					ai.TargetAngle = rand.Float64() * 2 * math.Pi
					ai.Timer = rand.IntN(20)
				}
			}
		}
	}
}

func getArrayIndex(levelWidth int, tileWidth int, x float64, y float64) int {
	xPosition := math.Floor(x / float64(tileWidth))
	yPosition := float64(levelWidth) * math.Floor(y/float64(tileWidth))

	return int(xPosition + yPosition)
}

func HandleForce(gravityRegistry map[EntityId]*Force, velocityRegistry map[EntityId]*Velocity) {
	for e := EntityId(0); e < NumEntities; e++ {
			force, fOk := gravityRegistry[e]
			velocity, vOk := velocityRegistry[e]
			if vOk && fOk {
				velocity.Vy += force.Y
		}
	}
}

func MoveStuff(level *Level, tileWidth int, positionRegistry map[EntityId]*Position, velocityRegistry map[EntityId]*Velocity, gravityRegistry map[EntityId]*Force) {

	for e := EntityId(0); e < NumEntities; e++ {
		position, pOk := positionRegistry[e]
		velocity, vOk := velocityRegistry[e]

		if pOk && vOk {

			oldX := position.X
			oldY := position.Y
			position.X += velocity.Vx
			position.Y += velocity.Vy

			// Check collision right and left
			if velocity.Vx > 0 && level.Data[getArrayIndex(level.Width, tileWidth, position.X+5, position.Y)] == 1 {
				position.X = oldX
				velocity.Vx = 0
			} else if velocity.Vx < 0 && level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y)] == 1 {
				position.X = oldX
				velocity.Vx = 0
			}
			// Check collision down
			if velocity.Vy > 0 {

				if level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y+5)] == 1 {
					velocity.Vy = 0
					collided := true
					for range 100 {
						position.Y--
						if level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y+5)] == 0 {
							collided = false
							break
						}
					}
					// Still not out of the wall, reset to oldY
					if collided {
						position.Y = oldY
					}

				}
			} else if velocity.Vy < 0 {

				// Check collision up
				if level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y-2)] == 1 {
					velocity.Vy = 0
					collided := true
					for range 100 {
						position.Y++
						if level.Data[getArrayIndex(level.Width, tileWidth, position.X, position.Y-2)] == 0 {
							collided = false
							break
						}
					}
					// Still not out of the wall, reset to oldY
					if collided {
						position.Y = oldY
					}

				}
			}

		}
	}

}
