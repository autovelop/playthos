package keyboard

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde/engine"
	"gde/input"
)

var keyboardInput *KeyListener

// Lots can be improved here
// Dont make a dynamic map of the keys bound in keymap. Rather
// predetermine a fixed set in memory
// Research best practices
type KeyListener struct {
	input.InputListener

	keydown func()
	keyup   func()
	keymap  map[int]func()

	// TEMP - not a big fan of this because glfw should be its own input system
	Window *glfw.Window
}

func (i *KeyListener) Init() {
	i.keymap = make(map[int]func())
	i.Window.SetKeyCallback(key_callback)
	// TEMP - again, not a fan setting this globally :(
	keyboardInput = i
}

// pointer to map of pointers?!?!?
func (i *KeyListener) Update(entities *map[string]*engine.Entity) {
	// TODO
}

func (i *KeyListener) Stop() {
	// TODO
}

func (i *KeyListener) BindOn(key int, callback func()) {
	fmt.Printf("Key BindOn - %v\n", key)
	i.keymap[key] = callback
}

func (i *KeyListener) KeyDown(key int) {
	fmt.Printf("KeyDown - %v\n", key)
}

func (i *KeyListener) KeyUp(key int) {
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
}
