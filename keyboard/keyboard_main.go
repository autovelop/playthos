// +build desktop,keyboard

package keyboard

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/opengl-glfw"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func init() {
	keyboard := &Keyboard{}
	engine.NewUnloadedObserverable(keyboard)
}

const (
	KeyCode_ESCAPE = 9
	KeyCode_SPACE  = 65
	KeyCode_ENTER  = 36
	KeyCode_LEFT   = 113
	KeyCode_UP     = 111
	KeyCode_DOWN   = 116
	KeyCode_RIGHT  = 114
)

var keyboard *Keyboard

type Keyboard struct {
	window     *glfw32.Window
	keypress   []func()
	keyrelease []func()
}

func (k *Keyboard) Prepare(settings *engine.Settings) {
	log.Println("Keyboard Prepare")
	k.keypress = make([]func(), 118, 118)
	k.keyrelease = make([]func(), 118, 118)
}

func (k *Keyboard) UnRegisterEntity(entity *engine.Entity) {
}
func (k *Keyboard) LoadComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *glfw.GLFW:
		k.window = component.GetWindow()
		break
	}

	// CANT GET THIS TO WORK
	// k.window.SetUserPointer(unsafe.Pointer(k))
	// DOING GLOBAL SADNESS INSTEAD
	keyboard = k

	k.window.SetKeyCallback(onKey)
}

func (k *Keyboard) SendKeyPress(keycode int) {
	if k.keypress[keycode] != nil {
		k.keypress[keycode]()
	}
}

func (k *Keyboard) SendKeyRelease(keycode int) {
	if k.keyrelease[keycode] != nil {
		k.keyrelease[keycode]()
	}
}

func (k *Keyboard) OnKey(keycode int, fnpress func(), fnrelease func()) {
	k.keypress[keycode] = fnpress
	k.keyrelease[keycode] = fnrelease
}

func (k *Keyboard) OnKeyPress(keycode int, fn func()) {
	k.keypress[keycode] = fn
}

func (k *Keyboard) OnKeyRelease(keycode int, fn func()) {
	k.keyrelease[keycode] = fn
}

func onKey(w *glfw32.Window, key glfw32.Key, scancode int, action glfw32.Action, mods glfw32.ModifierKey) {
	log.Printf("Key Action: %v\n", scancode)
	if action == glfw32.Press {
		keyboard.SendKeyPress(scancode)
	} else if action == glfw32.Release {
		keyboard.SendKeyRelease(scancode)
	}
}
