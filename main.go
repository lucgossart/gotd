package main

import (
	// "bytes"
	// "image"
	_ "image/png"
	// "fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/hajimehoshi/ebiten/v2/examples/resources/images"

	"game/ebitd"
	"game/load"
	"game/td"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

type Game struct {
	state *td.State
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

func (g *Game) Update() error {
	return g.state.Update()
}

func init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ouais grand")
}

func main() {
	var displayer ebitd.Displayer
	animations := load.Animations()
	count := 0
	state := td.CreateState(count, animations, displayer)
	game := Game{state}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
