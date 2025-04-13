package game

type EntityId int

var PositionRegistry = make(map[EntityId]*Position)
var VelocityRegistry = make(map[EntityId]*Velocity)
var PlayerInputRegistry = make(map[EntityId]*PlayerKeyPress)
var ForceRegistry = make(map[EntityId]*Force)

type Position struct {
	X float32
	Y float32
}

type Velocity struct {
	Vx float32
	Vy float32
}

type Force struct {
	X float32
	Y float32
}
