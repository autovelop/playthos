package input

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde"
)

var keyboardInput *Keyboard

// Lots can be improved here
// Dont make a dynamic map of the keys bound in keymap. Rather
// predetermine a fixed set in memory
// Research best practices
type Keyboard struct {
	gde.Input

	keydown func()
	keyup   func()
	keymap  map[int]func()

	// TEMP - not a big fan of this because glfw should be its own input system
	Window *glfw.Window
}

func (i *Keyboard) Init() {
	i.keymap = make(map[int]func())
	i.Window.SetKeyCallback(key_callback)
	// TEMP - again, not a fan setting this globally :(
	keyboardInput = i
}

func (i *Keyboard) Update(entities *map[string]gde.EntityRoutine) {
	// fmt.Printf("%v\n", len(*entities))
}

func (i *Keyboard) Stop() {
}

func (i *Keyboard) Bind(key int, callback func()) {
	fmt.Printf("Bind - %v\n", key)
	i.keymap[key] = callback
}

func (i *Keyboard) KeyDown(key int) {
	fmt.Printf("KeyDown - %v\n", key)
}

func (i *Keyboard) KeyUp(key int) {
	fmt.Printf("KeyUp - %v\n", key)
}

func key_callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		keyboardInput.KeyDown(int(key))
		binding := keyboardInput.keymap[int(key)]
		if binding != nil {
			binding()
		}
	} else if action == glfw.Release {
		keyboardInput.KeyUp(int(key))
	}
	// if key == glfw.KeyEscape && action == glfw.Press {
	// 	w.SetShouldClose(true)
	// }
}
