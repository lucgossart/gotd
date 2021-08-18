package td

// Interface for the loop. Implements ebiten's Game interface. 
// All the impurity of the game logic should happen in 
// the update function, the rest of the code should be made
// of pure functions. (Actually Game.Draw() has side effect, but
// should not modify the Game object.
type Game struct {
	count     int
	animation []Image
	screen    Screen
	displayer Displayer
}

func (game *Game) Draw(screen *Image) {
	game.displayer.Display(game.screen, game.animation[0], Position{})
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 0, 1
}

func (game *Game) Update() error {
	game.count ++
	return nil
}


