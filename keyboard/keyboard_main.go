// +build autovelop_playthos_keyboard,autovelop_playthos_glfw !play

package keyboard

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/opengl-glfw"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
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

	KeyUnknown      = -1
	KeySpace        = 32
	KeyApostrophe   = 39 /* ' */
	KeyComma        = 44 /* , */
	KeyMinus        = 45 /* - */
	KeyPeriod       = 46 /* . */
	KeySlash        = 47 /* / */
	Key0            = 48
	Key1            = 49
	Key2            = 50
	Key3            = 51
	Key4            = 52
	Key5            = 53
	Key6            = 54
	Key7            = 55
	Key8            = 56
	Key9            = 57
	KeySemicolon    = 59 /* ; */
	KeyEqual        = 61 /* = */
	KeyA            = 65
	KeyB            = 66
	KeyC            = 67
	KeyD            = 68
	KeyE            = 69
	KeyF            = 70
	KeyG            = 71
	KeyH            = 72
	KeyI            = 73
	KeyJ            = 74
	KeyK            = 75
	KeyL            = 76
	KeyM            = 77
	KeyN            = 78
	KeyO            = 79
	KeyP            = 80
	KeyQ            = 81
	KeyR            = 82
	KeyS            = 83
	KeyT            = 84
	KeyU            = 85
	KeyV            = 86
	KeyW            = 87
	KeyX            = 88
	KeyY            = 89
	KeyZ            = 90
	KeyLeftBracket  = 91  /* [ */
	KeyBackslash    = 92  /* \ */
	KeyRightBracket = 93  /* ] */
	KeyGraveAccent  = 96  /* ` */
	KeyWorld_1      = 161 /* non-US #1 */
	KeyWorld_2      = 162 /* non-US #2 */
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPage_down    = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF13          = 302
	KeyF14          = 303
	KeyF15          = 304
	KeyF16          = 305
	KeyF17          = 306
	KeyF18          = 307
	KeyF19          = 308
	KeyF20          = 309
	KeyF21          = 310
	KeyF22          = 311
	KeyF23          = 312
	KeyF24          = 313
	KeyF25          = 314
	KeyKP0          = 320
	KeyKP1          = 321
	KeyKP2          = 322
	KeyKP3          = 323
	KeyKP4          = 324
	KeyKP5          = 325
	KeyKP6          = 326
	KeyKP7          = 327
	KeyKP8          = 328
	KeyKP9          = 329
	KeyKPDecimal    = 330
	KeyKPMultiply   = 332
	KeyKPSubtract   = 333
	KeyKPAdd        = 334
	KeyKPEnter      = 335
	KeyKPEqual      = 336
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyMenu         = 348
	KeyLast         = KeyMenu
)

var keyboard *Keyboard

type Keyboard struct {
	engine.System
	window   *glfw32.Window
	keypress []func(...uint)
}

func (k *Keyboard) InitSystem() {
	k.keypress = make([]func(...uint), 512, 512) // this is probably too much but safe for now
}

func (k *Keyboard) Destroy() {
}

func (k *Keyboard) DeleteEntity(entity *engine.Entity) {}

func (k *Keyboard) On(key uint, fn func(...uint)) {
	// if k.Active() && k.keypress[key] != nil {
	if k.Active() {
		// log.Fatal("here")
		k.keypress[key] = fn
	}
}

func (k *Keyboard) AddComponent(component engine.ComponentRoutine) {}

func (k *Keyboard) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (k *Keyboard) AddIntegrant(integrant engine.IntegrantRoutine) {
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

func onKey(w *glfw32.Window, keycode glfw32.Key, scancode int, action glfw32.Action, mods glfw32.ModifierKey) {
	// log.Printf("Key Code: %v\n", key)
	// log.Printf("Scan Code: %v\n", scancode)
	if keyboard.keypress[keycode] != nil {
		keyboard.keypress[keycode](uint(action))
	}
}
