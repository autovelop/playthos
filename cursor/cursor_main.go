// +build deploy cursor

package cursor

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/glfw"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	engine.NewSystem(&Cursor{})
}

// Action number representing a button associated with the cursor (click, tap, stare, etc.)
type Action uint

const (
	// https://github.com/go-gl/glfw/blob/228fbf8cdbdda24bd57b5405bab240da3900b9a7/v3.2/glfw/glfw/include/GLFW/glfw3.h
	ActionRelease = 0
	ActionPress   = 1
	ActionRepeat  = 2
	ActionMove    = 3

	Button1      = 0
	Button2      = 1
	Button3      = 2
	Button4      = 3
	Button5      = 4
	Button6      = 5
	Button7      = 6
	Button8      = 7
	ButtonLast   = Button8
	ButtonLeft   = Button1
	ButtonRight  = Button2
	ButtonMiddle = Button3
)

var cursor *Cursor

// Cursor defines the position of a cursor on the current display
type Cursor struct {
	engine.System
	window      *glfw32.Window
	X           float64
	Y           float64
	buttonpress []func(...uint)
	move        []func(...uint)
}

// InitSystem called when the system plugs into the engine
func (c *Cursor) InitSystem() {
	c.buttonpress = make([]func(...uint), 32, 32) // this is probably too much but safe for now
}

// Destroy called when engine is gracefully shutting down
func (c *Cursor) Destroy() {}

// DeleteEntity removes all entity's compoents from this system
func (c *Cursor) DeleteEntity(entity *engine.Entity) {}

// On binds an given event based on a key
func (c *Cursor) On(key uint, fn func(...uint)) {
	if c.Active() {
		if key == ActionMove {
			c.move = append(c.move, fn)
		} else {
			c.buttonpress[key] = fn
		}
	}
}

// AddComponent unorphans a component by adding it to this system
func (c *Cursor) AddComponent(component engine.ComponentRoutine) {}

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (c *Cursor) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (c *Cursor) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		c.window = integrant.Window()
		break
	}

	// CANT GET THIS TO WORK
	// c.window.SetUserPointer(unsafe.Pointer(c))
	// DOING GLOBAL SADNESS INSTEAD
	cursor = c

	c.window.SetCursorPosCallback(onMove)
	c.window.SetMouseButtonCallback(onButton)
}

// onMove called every time the cursor moves
func onMove(w *glfw32.Window, x float64, y float64) {
	cursor.X = x
	cursor.Y = y
	for _, fn := range cursor.move {
		fn(uint(x), uint(y))
	}
	// cursor.[keycode](uint(action))
}

// onButton is called every time a cursor associated action is emitted
func onButton(w *glfw32.Window, button glfw32.MouseButton, action glfw32.Action, mods glfw32.ModifierKey) {
	if cursor.buttonpress[button] != nil {
		cursor.buttonpress[button](uint(action), uint(cursor.X), uint(cursor.Y))
	}
}
