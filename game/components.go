package game

type EntityId int

var PositionRegistry = make(map[EntityId]*Position)
var VelocityRegistry = make(map[EntityId]*Velocity)
var PlayerInputRegistry = make(map[EntityId]*PlayerKeyPress)
var ForceRegistry = make(map[EntityId]*Force)

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
