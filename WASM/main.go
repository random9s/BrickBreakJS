package main

import (
	"syscall/js"

	"github.com/random9s/BrickBreak/controller"
	"github.com/random9s/BrickBreak/gameObjects/border"
	"github.com/random9s/BrickBreak/gameObjects/paddle"
	"github.com/random9s/BrickBreak/gameObjects/screen"
)

var (
	d, w      js.Value
	canvas    js.Value
	callbacks []js.Callback
)

func initialize() {
	//Main vas are document and window
	d = js.Global().Get("document")
	w = js.Global().Get("window")
	callbacks = make([]js.Callback, 0)

	initCanvas()
}

func initCanvas() {
	//Get game canvas for WebGL
	canvas = d.Call("getElementById", "game-view")

	//Get viewport width and height
	var iw = w.Get("innerWidth").Float()
	var ih = w.Get("innerHeight").Float()

	canvas.Call("setAttribute", "width", iw)
	canvas.Call("setAttribute", "height", ih)
}

func JSCallback(fn func(args []js.Value)) js.Callback {
	var cb = js.NewCallback(fn)
	callbacks = append(callbacks, cb)
	return cb
}

func ReleaseCallbacks() {
	for _, cb := range callbacks {
		cb.Release()
	}
}

//In this, main initializes all global variables for later use
func main() {
	initialize()
	done := make(chan struct{}, 0)
	defer ReleaseCallbacks()

	var border = border.New(w, canvas, JSCallback)
	var renderer = controller.NewRenderer(
		screen.New(w, canvas, JSCallback),
		paddle.New(d, canvas, JSCallback),
	)

	var mainLoop js.Callback
	mainLoop = JSCallback(func(args []js.Value) {
		renderer.Redraw()

		js.Global().Call("requestAnimationFrame", mainLoop)
	})

	border.Redraw()
	js.Global().Call("requestAnimationFrame", mainLoop)

	<-done
}
