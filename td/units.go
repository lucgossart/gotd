package td

type Position struct {
	X float64
	Y float64
}

type Unit interface {
	move()
}
