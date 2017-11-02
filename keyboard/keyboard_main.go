// +build deploy keyboard

package keyboard

import (
	"github.com/autovelop/playthos"
)

func init() {
	engine.NewIntegrant(&Keyboard{})
}

type Keyboard struct {
	engine.Integrant
	keypress []func(...int)
}

func (k *Keyboard) InitIntegrant() {
	k.keypress = make([]func(...int), 350, 350) // this is probably too much but safe for now
}

func (k *Keyboard) Destroy() {
}

func (k *Keyboard) AddComponent(component engine.ComponentRoutine) {}

func (k *Keyboard) AddIntegrant(integrant engine.IntegrantRoutine) {
}

func (k *Keyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}
func (k *Keyboard) DeleteEntity(entity *engine.Entity) {}

func (k *Keyboard) IsSet(key int) bool {
	return k.keypress[key] != nil
}

func (k *Keyboard) Emit(kc int, a int) {
	k.keypress[kc](a)
}

func (k *Keyboard) On(key int, fn func(...int)) {
	k.keypress[key] = fn
}

const (
	ActionRelease = 0
	ActionPress   = 1
	ActionRepeat  = 2
)

var (
	KeyUnknown      int = 0
	KeySpace        int = 0
	KeyApostrophe   int = 0
	KeyComma        int = 0
	KeyMinus        int = 0
	KeyPeriod       int = 0
	KeySlash        int = 0
	Key0            int = 0
	Key1            int = 0
	Key2            int = 0
	Key3            int = 0
	Key4            int = 0
	Key5            int = 0
	Key6            int = 0
	Key7            int = 0
	Key8            int = 0
	Key9            int = 0
	KeySemicolon    int = 0
	KeyEqual        int = 0
	KeyA            int = 0
	KeyB            int = 0
	KeyC            int = 0
	KeyD            int = 0
	KeyE            int = 0
	KeyF            int = 0
	KeyG            int = 0
	KeyH            int = 0
	KeyI            int = 0
	KeyJ            int = 0
	KeyK            int = 0
	KeyL            int = 0
	KeyM            int = 0
	KeyN            int = 0
	KeyO            int = 0
	KeyP            int = 0
	KeyQ            int = 0
	KeyR            int = 0
	KeyS            int = 0
	KeyT            int = 0
	KeyU            int = 0
	KeyV            int = 0
	KeyW            int = 0
	KeyX            int = 0
	KeyY            int = 0
	KeyZ            int = 0
	KeyLeftBracket  int = 0
	KeyBackslash    int = 0
	KeyRightBracket int = 0
	KeyGraveAccent  int = 0
	KeyWorld_1      int = 0
	KeyWorld_2      int = 0
	KeyEscape       int = 0
	KeyEnter        int = 0
	KeyTab          int = 0
	KeyBackspace    int = 0
	KeyInsert       int = 0
	KeyDelete       int = 0
	KeyRight        int = 0
	KeyLeft         int = 0
	KeyDown         int = 0
	KeyUp           int = 0
	KeyPageUp       int = 0
	KeyPageDown     int = 0
	KeyHome         int = 0
	KeyEnd          int = 0
	KeyCapsLock     int = 0
	KeyScrollLock   int = 0
	KeyNumLock      int = 0
	KeyPrintScreen  int = 0
	KeyPause        int = 0
	KeyF1           int = 0
	KeyF2           int = 0
	KeyF3           int = 0
	KeyF4           int = 0
	KeyF5           int = 0
	KeyF6           int = 0
	KeyF7           int = 0
	KeyF8           int = 0
	KeyF9           int = 0
	KeyF10          int = 0
	KeyF11          int = 0
	KeyF12          int = 0
	KeyF13          int = 0
	KeyF14          int = 0
	KeyF15          int = 0
	KeyF16          int = 0
	KeyF17          int = 0
	KeyF18          int = 0
	KeyF19          int = 0
	KeyF20          int = 0
	KeyF21          int = 0
	KeyF22          int = 0
	KeyF23          int = 0
	KeyF24          int = 0
	KeyF25          int = 0
	KeyKP0          int = 0
	KeyKP1          int = 0
	KeyKP2          int = 0
	KeyKP3          int = 0
	KeyKP4          int = 0
	KeyKP5          int = 0
	KeyKP6          int = 0
	KeyKP7          int = 0
	KeyKP8          int = 0
	KeyKP9          int = 0
	KeyKPDecimal    int = 0
	KeyKPMultiply   int = 0
	KeyKPSubtract   int = 0
	KeyKPAdd        int = 0
	KeyKPEnter      int = 0
	KeyKPEqual      int = 0
	KeyLeftShift    int = 0
	KeyLeftControl  int = 0
	KeyLeftAlt      int = 0
	KeyLeftSuper    int = 0
	KeyRightShift   int = 0
	KeyRightControl int = 0
	KeyRightAlt     int = 0
	KeyRightSuper   int = 0
	KeyMenu         int = 0
	KeyLast         int = 0
)
