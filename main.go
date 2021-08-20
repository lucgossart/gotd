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
	"game/paths"
	"game/td"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	state td.State
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

func (g *Game) Update() error {
	return (&g.state).Update()
}

func init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ouais grand")
}

func createShootingArcherAnim() td.Animation {
	var runningArcherPaths = paths.GetShootingArcher()
	scale := td.Scale{WFactor: 3, HFactor: 3}
	duration := 7
	animation, err := ebitd.CreateAnim(runningArcherPaths, scale, duration)
	if err != nil {
		log.Fatalf("Error loading archers: %v", err)
	}
	return animation
}

func createRunningArcherAnim() td.Animation {
	var runningArcherPaths = paths.GetRunningArcher()
	scale := td.Scale{WFactor: 3, HFactor: 3}
	duration := 7
	animation, err := ebitd.CreateAnim(runningArcherPaths, scale, duration)
	if err != nil {
		log.Fatalf("Error loading archers: %v", err)
	}
	return animation
}

func main() {
	var displayer ebitd.Displayer
	var runningArcherAnim td.Animation = createRunningArcherAnim()
	var shootingArcherAnim td.Animation = createRunningArcherAnim()
	animations := []td.Animation{runningArcherAnim, shootingArcherAnim}

	count := 0
	state := td.CreateState(count, animations, displayer)
	game := Game{state}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
