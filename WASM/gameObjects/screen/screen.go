package screen

import (
	"syscall/js"
)

//CallbackCollector ...
type CallbackCollector func(func(args []js.Value)) js.Callback

//Screen game object
type Screen struct {
	elem   js.Value
	parent js.Value
	ctx    js.Value

	off, dim [2]float64
}

//New creates a new screen game object
func New(elem, canvas js.Value, cc CallbackCollector) *Screen {
	var iw = elem.Get("innerWidth").Float()
	var ih = elem.Get("innerHeight").Float()

	var s = &Screen{
		elem,
		canvas,
		canvas.Call("getContext", "2d"),
		[2]float64{80.0, 45.0},
		[2]float64{iw - 160.0, ih - 75.0},
	}

	elem.Call("addEventListener", "resize", cc(func(args []js.Value) {
		var iw = elem.Get("innerWidth").Float()
		var ih = elem.Get("innerHeight").Float()

		s.dim[0] = iw - 160.0
		s.dim[1] = ih - 75.0
	}))

	return s
}

//Redraw draws the new screen for each event
func (s *Screen) Redraw() {
	s.ctx.Call("beginPath")
	s.ctx.Call("rect", s.off[0], s.off[1], s.dim[0], s.dim[1])
	s.ctx.Set("fillStyle", "#000000")
	s.ctx.Call("fill")
	s.ctx.Call("closePath")
}
