package game

type Player struct {
	X          int
	Y          int
	Vx         int
	Vy         int
	keyPresses PlayerKeyPress
}

var NumEntities EntityId = 0

const MAX_ENTITIES = 100

var entities = make([]EntityId, MAX_ENTITIES)

var playerSpeed = 3
var players = make([]*Player, 0)
var playerKeyPresses = make([]*PlayerKeyPress, 0)

func AddPlayer() (*Player, EntityId) {

	PositionRegistry[NumEntities] = &Position{X: 50, Y: 50}
	VelocityRegistry[NumEntities] = &Velocity{Vx: 0, Vy: 0}
	PlayerInputRegistry[NumEntities] = &PlayerKeyPress{}
	NumEntities++

	newPlayer := Player{
		X:          100,
		Y:          100,
		Vx:         0,
		Vy:         0,
		keyPresses: PlayerKeyPress{},
	}
	players = append(players, &newPlayer)
	playerKeyPresses = append(playerKeyPresses, &PlayerKeyPress{})

	return &newPlayer, NumEntities - 1
}

type PlayerKeyPress struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func HandleInput(player *Player, input byte, entityId EntityId) {

	playa := PlayerInputRegistry[entityId]
	playa.Up = input&1 > 0
	playa.Down = input&2 > 0
	playa.Left = input&4 > 0
	playa.Right = input&8 > 0
	/*
	   player.keyPresses.Up = input&1 > 0
	   player.keyPresses.Down = input&2 > 0
	   player.keyPresses.Left = input&4 > 0
	   player.keyPresses.Right = input&8 > 0
	*/
}

type PlayerMessage struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func Tick() *map[EntityId]*Position {

	HandleDaInput(PlayerInputRegistry, VelocityRegistry)
	MoveStuff(PositionRegistry, VelocityRegistry)

	// TODO return the position registry to the client
	return &PositionRegistry
}
