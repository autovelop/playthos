// +build cursor
// +build glfw
// +build linux windows darwin

package cursor

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/opengl-glfw"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	engine.NewSystem(&Cursor{})
}

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

type Cursor struct {
	engine.System
	window      *glfw32.Window
	X           float64
	Y           float64
	buttonpress []func(...uint)
	move        []func(...uint)
}

func (c *Cursor) InitSystem() {
	c.buttonpress = make([]func(...uint), 32, 32) // this is probably too much but safe for now
}

func (c *Cursor) Destroy() {}

func (c *Cursor) DeleteEntity(entity *engine.Entity) {}

func (c *Cursor) On(key uint, fn func(...uint)) {
	if c.Active() {
		if key == ActionMove {
			c.move = append(c.move, fn)
		} else {
			c.buttonpress[key] = fn
		}
	}
}

func (c *Cursor) AddComponent(component engine.ComponentRoutine) {}

func (c *Cursor) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

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

func onMove(w *glfw32.Window, x float64, y float64) {
	cursor.X = x
	cursor.Y = y
	for _, fn := range cursor.move {
		fn(uint(x), uint(y))
	}
	// cursor.[keycode](uint(action))
}

func onButton(w *glfw32.Window, button glfw32.MouseButton, action glfw32.Action, mods glfw32.ModifierKey) {
	if cursor.buttonpress[button] != nil {
		cursor.buttonpress[button](uint(action), uint(cursor.X), uint(cursor.Y))
	}
}
