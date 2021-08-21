package td

import (
	"errors"
	"fmt"
)

type attacker interface {
	attack(unit)
}

type projectile interface {
	move()
	affect(unit, []unit)
}

// Interface for projectile shooters. shoot() must retrun a pointer towards a
// concrete projectile.
type shooter interface {
	shoot(projName string, direction vector, speed int) projectile
}

type spear struct {
	anim    *Animation
	pos     Position
	mover   rectMover
	effects []effect
}

func (s *spear) move() {
	s.pos = s.mover.move(s.pos)
}

func (s spear) affect(u unit, group []unit) {}

func newSpear(anim *Animation, pos Position, mover rectMover, effects []effect) *spear {
	return &spear{anim, pos, mover, effects}
}

// Factory returning projectile by keyword
func projByName(name string, pos Position, mover rectMover, effects []effect) (projectile, error) {
	anim, ok := s.Anims[name][""]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Pas d'animation pour le projectile %s", name))
	}
	switch name {
	case "spear":
		return newSpear(anim, pos, mover, effects), nil
	default:
		return nil, errors.New("Invalid projectile name")
	}
}
