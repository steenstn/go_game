package game

import "math/rand/v2"

var NumEntities EntityId = 0

const MAX_ENTITIES = 100

var entities = make([]EntityId, MAX_ENTITIES)

var playerSpeed float64 = 7.0

type Level struct {
	Width  int
	Height int
	Data   []int
}

var PlayerEntities = make([]EntityId, 0)

var CurrentLevel Level

func AddPlayer() EntityId {
	PositionRegistry[NumEntities] = &Position{X: 40 * 50, Y: 60}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	PlayerInputRegistry[NumEntities] = &PlayerKeyPress{}
	GravityRegistry[NumEntities] = &Force{X: 0, Y: 1}

	NumEntities++

	playerEntityId := NumEntities - 1
	PlayerEntities = append(PlayerEntities, playerEntityId)
	return NumEntities - 1
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func createThing(x float64, y float64) {
	PositionRegistry[NumEntities] = &Position{X: x, Y: y}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	AIRegistry[NumEntities] = &AIMovement{Timer: rand.IntN(100)}
	GravityRegistry[NumEntities] = &Force{X: 0, Y: 0}

	NumEntities++
}

func InitGame() {
	println("Init game")
	for range 10 {
		createThing(30*50, 100)
	}

	// TODO: Get this from level, don't hardcode
	CurrentLevel.Width = 50
	CurrentLevel.Height = 20

	println("Loading level")
	CurrentLevel.Data = LoadLevel("game/level.txt")
}

func HandleInput(input byte, entityId EntityId) {
	player := PlayerInputRegistry[entityId]
	player.Up = input&1 > 0
	player.Down = input&2 > 0
	player.Left = input&4 > 0
	player.Right = input&8 > 0
}

func Tick() map[EntityId]*Position {

	HandleDaInput(PlayerInputRegistry, VelocityRegistry)
	HandleAI(AIRegistry, VelocityRegistry, PlayerEntities, PositionRegistry)
	MoveStuff(&CurrentLevel, PositionRegistry, VelocityRegistry, GravityRegistry)

	return PositionRegistry
}
