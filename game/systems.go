package game

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

func MoveStuff(positionRegistry map[EntityId]*Position, velocityRegistry map[EntityId]*Velocity, forceRegistry map[EntityId]*Force) {

	for e := EntityId(0); e < NumEntities; e++ {
		position, pOk := positionRegistry[e]
		velocity, vOk := velocityRegistry[e]
		force, fOk := forceRegistry[e]

		if pOk && vOk && fOk {
			velocity.Vx += force.X
			velocity.Vy += force.Y

			position.X += velocity.Vx
			position.Y += velocity.Vy
			if position.Y > 400 {
				position.Y = 400
				velocity.Vy = 0
			}
		}
	}

}
