package game

var NumEntities EntityId = 0

const MAX_ENTITIES = 100

var entities = make([]EntityId, MAX_ENTITIES)

var playerSpeed float64 = 5.0

type Level struct {
	Width  int
	Height int
	Data   []int
}

var levelData = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 1, 1, 0, 0, 1, 0, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var CurrentLevel Level

func AddPlayer() EntityId {
	PositionRegistry[NumEntities] = &Position{X: 60, Y: 60}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	PlayerInputRegistry[NumEntities] = &PlayerKeyPress{}
	GravityRegistry[NumEntities] = &Force{X: 0, Y: 1}

	NumEntities++

	return NumEntities - 1
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func InitGame() {
	println("Init game")

	CurrentLevel.Width = 50
	CurrentLevel.Height = 20
	//	CurrentLevel.Data = levelData

	CurrentLevel.Data = LoadLevel("game/level.txt")

}

func HandleInput(input byte, entityId EntityId) {
	player := PlayerInputRegistry[entityId]
	player.Up = input&1 > 0
	player.Down = input&2 > 0
	player.Left = input&4 > 0
	player.Right = input&8 > 0
}

func Tick() *map[EntityId]*Position {

	HandleDaInput(PlayerInputRegistry, VelocityRegistry)
	MoveStuff(&CurrentLevel, PositionRegistry, VelocityRegistry, GravityRegistry)

	return &PositionRegistry
}
