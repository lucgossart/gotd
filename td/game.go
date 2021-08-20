package td

import "log"

// All the impurity of the game logic should happen in
// the update function, the rest of the code should be made
// of pure functions. (Actually State.Draw() has side effect, but
// should not modify the State object.)
type State struct {
	count      int
	animations []Animation
	position   Position
	displayer  Displayer
}

func (state State) Draw(screen PImage) {
	for _, animation := range state.animations {
		state.displayer.Display(
			screen,
			getImage(animation),
			state.position,
			animation.Scale,
		)
	}
}

func (state *State) Update() error {
	state.count++
	if state.count%3 == 0 {
		state.position.X++
		state.position.Y++
		log.Println(state.position)
	}
	for _, animation := range state.animations {
		animation.update()
	}
	return nil
}

func CreateState(count int, animations []Animation, displayer Displayer) State {
	return State{count, animations, Position{}, displayer}
}
