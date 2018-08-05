// +build deploy glfw,keyboard

package keyboard

import (
	"fmt"
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/glfw"
	"github.com/autovelop/playthos/keyboard"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	engine.NewIntegrant(&GLFWKeyboard{})
	fmt.Println("> GLFW (Keyboard): Ready")
}

type GLFWKeyboard struct {
	engine.Integrant
	keyboard *keyboard.Keyboard
	window   *glfw32.Window
}

func (k *GLFWKeyboard) InitIntegrant()                                 {}
func (k *GLFWKeyboard) Destroy()                                       {}
func (k *GLFWKeyboard) DeleteEntity(entity *engine.Entity)             {}
func (k *GLFWKeyboard) AddComponent(component engine.ComponentRoutine) {}

func (k *GLFWKeyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (k *GLFWKeyboard) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		k.window = integrant.Window()
		if k.keyboard != nil {
			k.window.SetKeyCallback(k.onKey)
		}
		break
	case *keyboard.Keyboard:
		k.keyboard = integrant
		if k.window != nil {
			k.window.SetKeyCallback(k.onKey)
		}
		break
	}
}

func (k *GLFWKeyboard) onKey(w *glfw32.Window, keycode glfw32.Key, scancode int, action glfw32.Action, mods glfw32.ModifierKey) {
	if k.keyboard.IsSet(int(keycode)) {
		k.keyboard.Emit(int(keycode), int(action))
	}
}
