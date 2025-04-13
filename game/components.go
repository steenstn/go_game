package game

type EntityId int

var PositionRegistry = make(map[EntityId]*Position)
var VelocityRegistry = make(map[EntityId]*Velocity)
var PlayerInputRegistry = make(map[EntityId]*PlayerKeyPress)

type Position struct {
	X int
	Y int
}

type Velocity struct {
	Vx int
	Vy int
}
