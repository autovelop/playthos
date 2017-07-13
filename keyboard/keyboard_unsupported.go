// +build keyboard
// +build !darwin
// +build !linux
// +build !windows

package keyboard

import (
	"github.com/autovelop/playthos"
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

type Keyboard struct {
	engine.System
	keypress []func(...uint)
}

func (k *Keyboard) InitSystem() {
	k.keypress = make([]func(...uint), 118, 118)
}

func (k *Keyboard) DeleteEntity(entity *engine.Entity) {}

func (k *Keyboard) On(key uint, fn func(...uint)) {}

func (k *Keyboard) NewComponent(component engine.ComponentRoutine) {}

func (k *Keyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (k *Keyboard) NewIntegrant(integrant engine.IntegrantRoutine) {}

func (k *Keyboard) SendKeyPress(keycode int) {
}
