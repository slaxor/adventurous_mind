package persistence

type Image struct {
	Name   string
	Format string
	Path   string
}

type Sprite struct {
	ImageSeq []Image
}

type Actor struct {
	Name   string
	Sprite Sprite
	posx   float64
	posy   float64
}

type Page struct {
	Url        string
	Background Image
}

type Game struct {
	Actors []Actor
	Pages  []Page
}

func NewGame() Game {
	var g Game
	g.Actors = make([]Actor, 10)
	for _, a := range g.Actors {
		a.Name = "Foo"
	}

	g.Pages = make([]Page, 4)
	for _, p := range g.Pages {
		p.Url = "http://example.com"
	}

	return g
}
