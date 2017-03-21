package mouse

import (
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde/engine"
	"gde/input"
)

var pointerInput *MoveListener

type MoveListener struct {
	input.InputListener

	move func(float64, float64)

	Window *glfw.Window
}

func (i *MoveListener) Init() {
	i.Window.SetCursorPosCallback(mouse_move)

	// TODO: You can do better than this Fanus
	pointerInput = i
}

func (i *MoveListener) Update(entities *map[string]*engine.Entity) {
}

func (i *MoveListener) Stop() {
}

func (i *MoveListener) BindMove(callback func(float64, float64)) {
	i.move = callback
}

// TODO: You can do better than this Fanus
func mouse_move(w *glfw.Window, x float64, y float64) {
	if pointerInput.move != nil {
		pointerInput.move(x, y)
	}
}

// belongs in click_listener.go (TODO)
// i.Window.SetMouseButtonCallback(click_callback)
func pointer_button_callback(w *glfw.Window, key glfw.Key, action glfw.Action) {
}

func (i *MoveListener) ButtonDown(key int) {
	log.Printf("%v\n", key)
}

func (i *MoveListener) ButtonUp(key int) {
	log.Printf("%v\n", key)
}
