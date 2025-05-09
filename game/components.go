package game

type EntityId int

var PositionRegistry = make(map[EntityId]*Position)
var VelocityRegistry = make(map[EntityId]*Velocity)
var PlayerInputRegistry = make(map[EntityId]*PlayerKeyPress)
var GravityRegistry = make(map[EntityId]*Force)
var AIRegistry = make(map[EntityId]*AIMovement)

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
