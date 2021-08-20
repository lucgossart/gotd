package ebitd

import (
	"log"
	"fmt"
	"errors"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"

	"game/td"
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

func (Displayer) Display(
	screen   td.PImage,
	image    td.PImage,
	position td.Position,
	scale    td.Scale,
) error {
	ebiscreen, ok := screen.(*ebiten.Image)
	if !ok {fmt.Println(1); panic("a")}
	ebimage, ok := image.(*ebiten.Image)
	if !ok  {log.Println(2); return errors.New("a")}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.WFactor, scale.HFactor)
	op.GeoM.Translate(position.X, position.Y)
	ebiscreen.DrawImage(ebimage, op)
	return nil
}

func CreateAnim(paths []string, scale td.Scale, duration int) (anim td.Animation, err error) {
	var loader Loader
	var list []td.PImage = make([]td.PImage, 0, len(paths))
	for _, path := range(paths) {
		image, err := loader.Load(path)
		if err != nil { return  td.Animation{}, err }
		tdImage := td.PImage(image)
		list = append(list, tdImage)
	}
	return td.Animation{
		Images: list,
		Scale: scale,
		Duration: duration,
	}, nil
}
