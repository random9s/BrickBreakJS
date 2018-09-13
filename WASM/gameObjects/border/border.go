package border

import (
	"syscall/js"
)

//CallbackCollector ...
type CallbackCollector func(func(args []js.Value)) js.Callback

//Border contains
type Border struct {
	elem   js.Value
	parent js.Value
	ctx    js.Value

	off, dim [2]float64
}

//New creates a new paddle game object
func New(elem, canvas js.Value, cc CallbackCollector) *Border {
	var iw = elem.Get("innerWidth").Float()
	var ih = elem.Get("innerHeight").Float()

	var b = &Border{
		elem,
		canvas,
		canvas.Call("getContext", "2d"),
		[2]float64{0.0, 0.0},
		[2]float64{iw, ih},
	}

	elem.Call("addEventListener", "resize", cc(func(args []js.Value) {
		var iw = elem.Get("innerWidth").Float()
		var ih = elem.Get("innerHeight").Float()

		b.dim[0] = iw
		b.dim[1] = ih
	}))

	return b
}

//Redraw draws the new paddle for each event
func (b *Border) Redraw() {
	//clear old paddle
	b.ctx.Call("beginPath")

	b.ctx.Call("rect", b.off[0], b.off[1], b.dim[0], b.dim[1])

	//color new paddle
	b.ctx.Set("fillStyle", "#F4E242")
	b.ctx.Call("fill")
	b.ctx.Call("closePath")
}
