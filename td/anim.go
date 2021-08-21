package td

// import "log"

// Interface representing a loader loading image from filename.
type Loader interface {
	Load(path string) (*Image, error)
}

type Image interface{}

// Interface to pointer of image.
type PImage interface{}

type Displayer interface {
	Display(screen PImage, image PImage, position Position, scale Scale) error
}

// Specifies by how much an image should be resized.
type Scale struct {
	WFactor float64
	HFactor float64
}

// Contains a slice of PImages, and internal index,
// duration and nbTicks.
// The coeffs specify the factor by which to multiply the
// base image.
// It is important that no method modifies the underlying
// array of Images.
type Animation struct {
	nbTicks  int
	Duration int
	index    int
	images   []PImage
	Scale    Scale
}

func NewAnim(
	nbTicks int,
	Duration int,
	index int,
	images []PImage,
	Scale Scale,
) *Animation {
	return &Animation{nbTicks, Duration, index, images, Scale}
}

// Pure function,
// Returns image corresponding to a.index.
func getImage(a Animation) PImage {
	return a.images[a.index]
}

// All the impurity should reside here.
// Increases a.nbTicks and rotates a.index
// if a.nbTicks is greater than a.Duration
func (a *Animation) change() {
	a.nbTicks++
	if a.nbTicks > a.Duration {
		a.index = (a.index + 1) % len(a.images)
		a.nbTicks = 0
	}
}
