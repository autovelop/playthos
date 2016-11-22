package pointer

import (
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde"
)

var pointerInput *Pointer

type Pointer struct {
	gde.Input

	move func(float64, float64)

	Window *glfw.Window
}

func (i *Pointer) Init() {
	// i.Window.SetMouseButtonCallback(click_callback)
	i.Window.SetCursorPosCallback(pointer_callback)

	pointerInput = i
}

func (i *Pointer) Update(entities *map[string]gde.EntityRoutine) {
	// log.Printf("%v\n", len(*entities))
}

func (i *Pointer) Stop() {
}

func (i *Pointer) BindMove(callback func(float64, float64)) {
	i.move = callback
}

func (i *Pointer) ButtonDown(key int) {
	log.Printf("%v\n", key)
}

func (i *Pointer) ButtonUp(key int) {
	log.Printf("%v\n", key)
}

func pointer_callback(w *glfw.Window, x float64, y float64) {
	pointerInput.move(x, y)
}

func pointer_button_callback(w *glfw.Window, key glfw.Key, action glfw.Action) {
	if action == glfw.Press {
	}
}
