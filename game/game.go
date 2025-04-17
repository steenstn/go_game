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

var currentLevel = Level{
	Width:  10,
	Height: 10,
	Data:   levelData,
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

func AddPlayer() EntityId {

	PositionRegistry[NumEntities] = &Position{X: 50, Y: 50}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	PlayerInputRegistry[NumEntities] = &PlayerKeyPress{}
	ForceRegistry[NumEntities] = &Force{X: 0, Y: 1}

	NumEntities++

	return NumEntities - 1
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
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
	MoveStuff(&currentLevel, PositionRegistry, VelocityRegistry, ForceRegistry)

	return &PositionRegistry
}
