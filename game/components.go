package game

type EntityId int

var PositionRegistry = make(map[EntityId]*Position)
var VelocityRegistry = make(map[EntityId]*Velocity)
var PlayerInputRegistry = make(map[EntityId]*PlayerKeyPress)
var GravityRegistry = make(map[EntityId]*Force)
var AIRegistry = make(map[EntityId]*AIMovement)
var CircleMovementRegistry = make(map[EntityId]*CircleMovement)
var EntityTypeRegistry = make(map[EntityId]int)
var PlayerStateRegistry = make(map[EntityId]*PlayerState)

type EntityType byte

const (
	Player EntityType = 0
	Fly    EntityType = 1
	Spider EntityType = 2
)

type Position struct {
	X float64
	Y float64
}

type Velocity struct {
	Vx float64
	Vy float64
}

type Force struct {
	X float64
	Y float64
}

type AIMovement struct {
	Timer        int
	CurrentAngle float64
	TargetAngle  float64
	State        int
}

type CircleMovement struct {
	Timer     int
	Direction int
}

var MAX_JUMP int = 4
type PlayerState struct {
	JumpCounter int
	Jumping bool
}
