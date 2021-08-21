package td

import (
	"log"
	"math/rand"
	"time"
)

// All the impurity of the game logic should happen in
// the update function, the rest of the code should be made
// of pure functions. (Actually State.Draw() has side effect, but
// should not modify the State object.)
// All display happens through abstract Displayer.
type State struct {
	count     int
	Anims     map[string]map[string]*Animation
	archers   []*archer
	spears    []*spear
	displayer Displayer
}

func (state State) Draw(screen PImage) {
	for _, archer := range state.archers {
		state.displayer.Display(screen, getImage(*archer.anim), archer.pos, archer.anim.Scale)
	}
}

var archerShoots map[*archer]bool = make(map[*archer]bool, 5)

func (state *State) Update() error {
	state.count++
	if state.count%1 == 0 {
		for _, archer := range state.archers {
			log.Println(archerShoots[archer])
			if archer.pos.X < 30 {
				archer.move(2, 1)
			} else if !archerShoots[archer] {
				log.Println("a")
				archer.shoot()
				archerShoots[archer] = true
			}
			archer.anim.change()
		}
	}
	return nil
}

// Expects a map of map of pointers to animations, for instance
// amap["archer"] = map[string]*Animation{"move": anim1, "shoot": anim2}
// etc.
// The entities created copy the underlying reference animations without
// modifying them.
func CreateState(count int, amap map[string]map[string]*Animation, displayer Displayer) *State {
	var archers [5]*archer
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		anim := *amap["archer"]["move"]
		archers[i] = newArcher(&anim, amap["archer"], Position{X: 1000 * rand.Float64(), Y: 600 * rand.Float64()})
	}
	spears := []*spear{}
	return &State{count, amap, archers[:], spears, displayer}
}
