package paddle

import (
	"syscall/js"
)

//CallbackCollector ...
type CallbackCollector func(func(args []js.Value)) js.Callback

//Paddle contains
type Paddle struct {
	elem   js.Value
	parent js.Value
	ctx    js.Value
	pos    [2]float64
}

//New creates a new paddle game object
func New(elem, canvas js.Value, cc CallbackCollector) *Paddle {
	var p = &Paddle{
		elem,
		canvas,
		canvas.Call("getContext", "2d"),
		[2]float64{0.0, 300.0},
	}

	elem.Call("addEventListener", "mousemove", cc(func(args []js.Value) {
		e := args[0]
		p.pos[0] = e.Get("clientX").Float()
	}))

	return p
}

//Redraw draws the new paddle for each event
func (p *Paddle) Redraw() {
	//clear old paddle
	p.ctx.Call("beginPath")

	//draw new paddle
	p.ctx.Call("rect", p.pos[0], p.pos[1], 100, 25)

	//color new paddle
	p.ctx.Set("fillStyle", "#FFFFFF")
	p.ctx.Call("fill")
	p.ctx.Call("closePath")
}
