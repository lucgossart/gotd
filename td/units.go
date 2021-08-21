package td

import "log"

type unit interface {
	move()
	setAnim(string) Animation
}

// Archer unit. Has a current animation anim
// which can be set by setAnim with the keyword
// move or shoot.
type archer struct {
	anim         *Animation
	animsByKword map[string]*Animation
	pos          Position
}

func newArcher(
	anim *Animation,
	animsByKword map[string]*Animation,
	pos Position,
) *archer {
	newAnim := *anim
	log.Printf("anim: %p, newAnim: %p", anim, &newAnim)
	return &archer{&newAnim, animsByKword, pos}
}

// Non pure. Moves position by given offsets.
func (a *archer) move(xOffset float64, yOffset float64) {
	a.setAnim("move")
	a.pos.X += xOffset
	a.pos.Y += yOffset
}

func (a *archer) shoot() {
	a.setAnim("shoot")
}

// No need for deep copy as the pointer to Images remains
// unchanged for the type Animation.
func (a *archer) setAnim(s string) {
	commonAnimAddress, ok := a.animsByKword[s]
	if !ok {
		panic("Not an appropriate keyword for animation")
	}
	anim := *commonAnimAddress
	a.anim = &anim
}
