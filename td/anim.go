package td


type Loader interface {
	Load(path string) (*Image, error)
}

type Image interface{}

type Screen interface{}

type Displayer interface {
	Display(screen Screen, image Image, position Position) error
}
