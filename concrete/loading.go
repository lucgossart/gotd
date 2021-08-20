package ebitd

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)


// Implements the loader interface.
// Loads images via the Load method.
type Loader struct{}

// Load images from @path.
func (l Loader) Load(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img, nil
}

type Displayer struct{}

type Position struct {
	x int
	y int
}

func (Displayer) Display(screen *ebiten.Image, image ebiten.Image, position Position) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(&image, op)
}
