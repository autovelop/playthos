// +build desktop,keyboard

package keyboard

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/opengl-glfw"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func init() {
	engine.NewSystem(&Keyboard{})
}

type Action uint

const (
	// https://github.com/go-gl/glfw/blob/228fbf8cdbdda24bd57b5405bab240da3900b9a7/v3.2/glfw/glfw/include/GLFW/glfw3.h
	ActionRelease = 0
	ActionPress   = 1
	ActionRepeat  = 2

	KeyEscape = 9
	KeySpace  = 65
	KeyEnter  = 36
	KeyLeft   = 113
	KeyUp     = 111
	KeyDown   = 116
	KeyRight  = 114
)

var keyboard *Keyboard

type Keyboard struct {
	engine.System
	window     *glfw32.Window
	keypress   []func(...uint)
	keyrelease []func()
}

func (k *Keyboard) InitSystem() {
	log.Println("Keyboard Prepare")
	k.keypress = make([]func(...uint), 118, 118)
	k.keyrelease = make([]func(), 118, 118)
}

func (k *Keyboard) On(key uint, fn func(...uint)) {
	k.keypress[key] = fn
}

func (k *Keyboard) NewComponent(component engine.ComponentRoutine) {}

func (k *Keyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (k *Keyboard) NewIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		k.window = integrant.Window()
		break
	}

	// CANT GET THIS TO WORK
	// k.window.SetUserPointer(unsafe.Pointer(k))
	// DOING GLOBAL SADNESS INSTEAD
	keyboard = k

	k.window.SetKeyCallback(onKey)
}

func (k *Keyboard) SendKeyPress(keycode int) {
}

func (k *Keyboard) SendKeyRelease(keycode int) {
	if k.keyrelease[keycode] != nil {
		k.keyrelease[keycode]()
	}
}

func onKey(w *glfw32.Window, key glfw32.Key, scancode int, action glfw32.Action, mods glfw32.ModifierKey) {
	log.Printf("Key Action: %v\n", scancode)
	if keyboard.keypress[scancode] != nil {
		keyboard.keypress[scancode](uint(action))
	}
}
