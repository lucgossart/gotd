package td

type mover interface {
	move(Position) Position
}

type Position struct {
	X float64
	Y float64
}

type vector struct {
	x float64
	y float64
}

type rectMover struct {
	dir vector
}

func (r rectMover) move(pos Position) Position {
	return translate(pos, r.dir)
}

func add(v1 vector, v2 vector) (v vector) {
	v.x = v1.x + v2.x
	v.y = v1.y + v2.y
	return
}

func translate(pos Position, v vector) Position {
	return Position{pos.X + v.x, pos.Y + v.y}
}
