package td

// import "log"

// Interface representing a loader loading image from filename.
type Loader interface {
	Load(path string) (*Image, error)
}

type Image interface{}
// Interface to pointer of image.
type PImage interface{}

// Specifies by how much an image should be resized.
type Scale struct {
	WFactor float64
	HFactor float64
}

// Contains a slice of PImages, and internal index,
// duration and nbTicks.
// The coeffs specify the factor by which to multiply the
// base image.
type Animation struct{
	nbTicks     int
	Duration    int
	index       int
	Images      []PImage
	Scale       Scale
}

// Pure function,
// Returns image corresponding to a.index.
func getImage(a Animation) PImage {
	return a.Images[a.index]
}

// All the impurity should reside here.
// Increases a.nbTicks and rotates a.index
// if a.nbTicks is greater than a.Duration
func (a *Animation)update() {
	a.nbTicks++
	if a.nbTicks > a.Duration {
		a.index = (a.index + 1) % len(a.Images)
		a.nbTicks = 0
	}
}

type Displayer interface {
	Display(screen PImage, image PImage, position Position, scale Scale) error
}
