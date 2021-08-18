package main

import (
	// "bytes"
	// "image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/examples/resources/images"

	"./td"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game td.Game

// func (g *Game) Update() error {
// 	g.count ++
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ouais grand")

	game := Game{}

	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
