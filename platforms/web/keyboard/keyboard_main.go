// +build deploy webkeyboard

package keyboard

import (
	"github.com/autovelop/playthos"
	// glfw "github.com/autovelop/playthos/glfw"
	"github.com/autovelop/playthos/keyboard"
	// glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	engine.NewIntegrant(&WebKeyboard{})
}

type WebKeyboard struct {
	engine.Integrant
	keyboard *keyboard.Keyboard
}

func (k *WebKeyboard) InitIntegrant()                                 {}
func (k *WebKeyboard) Destroy()                                       {}
func (k *WebKeyboard) DeleteEntity(entity *engine.Entity)             {}
func (k *WebKeyboard) AddComponent(component engine.ComponentRoutine) {}

func (k *WebKeyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (k *WebKeyboard) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *keyboard.Keyboard:
		k.keyboard = integrant
		document := js.Global.Get("document")
		document.Call("addEventListener", "keydown", k.onKey, false)
		document.Call("addEventListener", "keyup", k.onKey, false)
		// if k.window != nil {
		// 	k.window.SetKeyCallback(k.onKey)
		// }
		break
	}
}

func (k *WebKeyboard) onKey(ev *js.Object) {
	keycode := int(ev.Get("keyCode").Int())
	eventType := ev.Get("type").String()
	if k.keyboard.IsSet(keycode) {
		switch eventType {
		case "keydown":
			k.keyboard.Emit(keycode, int(1))
			break
		case "keyup":
			k.keyboard.Emit(keycode, int(0))
			break
		}
	}
}
