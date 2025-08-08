package game

import "math/rand/v2"

var NumEntities EntityId = 0

const MAX_ENTITIES = 200
const TILE_SIZE = 50

var entities = make([]EntityId, MAX_ENTITIES)

var playerSpeed float64 = 8.0

type Level struct {
	Width  int
	Height int
	Data   []int
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

var PlayerEntities = make([]EntityId, 0)

var CurrentLevel Level

func AddPlayer() EntityId {
	PositionRegistry[NumEntities] = &Position{X: 40 * 50, Y: 60}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	PlayerInputRegistry[NumEntities] = &PlayerKeyPress{}
	GravityRegistry[NumEntities] = &Force{X: 0, Y: 1}
	EntityTypeRegistry[NumEntities] = 0
	PlayerStateRegistry[NumEntities] = &PlayerState{}

	NumEntities++

	playerEntityId := NumEntities - 1
	PlayerEntities = append(PlayerEntities, playerEntityId)
	return NumEntities - 1
}

func createFly(x float64, y float64) {
	PositionRegistry[NumEntities] = &Position{X: x, Y: y}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	AIRegistry[NumEntities] = &AIMovement{Timer: rand.IntN(100)}
	GravityRegistry[NumEntities] = &Force{X: 0, Y: 0}
	EntityTypeRegistry[NumEntities] = 1

	NumEntities++
}

func createSpider(x float64, y float64) {
	PositionRegistry[NumEntities] = &Position{X: x, Y: y}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 3, Vy: 0}
	RepeatingTimerRegistry[NumEntities] = &RepeatingTimer{GenerateStartValue: func() int { return rand.IntN(100)}} 
	CircleMovementRegistry[NumEntities] = &CircleMovement{Timer: 20, Direction: 1}
	EntityTypeRegistry[NumEntities] = 2

	NumEntities++
}

func InitGame() {
	println("Init game")
	for range 10 {
		createFly(30*50, 300)
	}

	createSpider(35*50, 300)

	// TODO: Get this from level, don't hardcode
	CurrentLevel.Width = 100
	CurrentLevel.Height = 40

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

func Tick() {

	HandleDaInput(PlayerInputRegistry, VelocityRegistry, PlayerStateRegistry)
	HandleTimers(RepeatingTimerRegistry)
	HandleAI(AIRegistry, VelocityRegistry, PlayerEntities, PositionRegistry)
	HandleCircleMovement(CircleMovementRegistry, VelocityRegistry, RepeatingTimerRegistry)
	HandleForce(GravityRegistry, VelocityRegistry)
	MoveStuff(&CurrentLevel, TILE_SIZE, PositionRegistry, VelocityRegistry, GravityRegistry, PlayerStateRegistry)

}
