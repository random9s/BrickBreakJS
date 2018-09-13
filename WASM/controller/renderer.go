package controller

//Redrawer is used to redraw a game object
type Redrawer interface {
	Redraw()
}

//Renderer renders all game objects
type Renderer struct {
	Objects []Redrawer
}

//NewRenderer ...
func NewRenderer(objs ...Redrawer) *Renderer {
	var o = make([]Redrawer, 0)
	if len(objs) > 0 {
		o = objs
	}

	return &Renderer{
		o,
	}
}

//Redraw ...
func (rc *Renderer) Redraw() {
	for _, obj := range rc.Objects {
		obj.Redraw()
	}
}

//Add appends a new redrawer object
func (rc *Renderer) Add(obj Redrawer) {
	rc.Objects = append(rc.Objects, obj)
}
