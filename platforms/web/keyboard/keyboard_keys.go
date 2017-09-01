// +build !deploy,webkeyboard

package keyboard

import kb "github.com/autovelop/playthos/keyboard"

func init() {
	// kb.KeyUnknown = -1
	kb.KeySpace = 32
	kb.KeyApostrophe = 222 /* ' */
	kb.KeyComma = 188      /* , */
	kb.KeyMinus = 189      /* - */
	kb.KeyPeriod = 190     /* . */
	kb.KeySlash = 191      /* / */
	kb.Key0 = 48
	kb.Key1 = 49
	kb.Key2 = 50
	kb.Key3 = 51
	kb.Key4 = 52
	kb.Key5 = 53
	kb.Key6 = 54
	kb.Key7 = 55
	kb.Key8 = 56
	kb.Key9 = 57
	kb.KeySemicolon = 186 /* ; */
	kb.KeyEqual = 187     /* = */
	kb.KeyA = 65
	kb.KeyB = 66
	kb.KeyC = 67
	kb.KeyD = 68
	kb.KeyE = 69
	kb.KeyF = 70
	kb.KeyG = 71
	kb.KeyH = 72
	kb.KeyI = 73
	kb.KeyJ = 74
	kb.KeyK = 75
	kb.KeyL = 76
	kb.KeyM = 77
	kb.KeyN = 78
	kb.KeyO = 79
	kb.KeyP = 80
	kb.KeyQ = 81
	kb.KeyR = 82
	kb.KeyS = 83
	kb.KeyT = 84
	kb.KeyU = 85
	kb.KeyV = 86
	kb.KeyW = 87
	kb.KeyX = 88
	kb.KeyY = 89
	kb.KeyZ = 90
	kb.KeyLeftBracket = 219  /* [ */
	kb.KeyBackslash = 220    /* \ */
	kb.KeyRightBracket = 221 /* ] */
	kb.KeyGraveAccent = 192  /* ` */
	// kb.KeyWorld_1 = 161     /* non-US #1 */
	// kb.KeyWorld_2 = 162     /* non-US #2 */
	kb.KeyEscape = 27
	kb.KeyEnter = 13
	kb.KeyTab = 9
	kb.KeyBackspace = 8
	kb.KeyInsert = 45
	kb.KeyDelete = 46
	kb.KeyRight = 39
	kb.KeyLeft = 37
	kb.KeyDown = 40
	kb.KeyUp = 38
	kb.KeyPageUp = 34
	kb.KeyPageDown = 34
	kb.KeyHome = 36
	kb.KeyEnd = 35
	kb.KeyCapsLock = 20
	// kb.KeyScrollLock = 281
	kb.KeyNumLock = 144
	// kb.KeyPrintScreen = 283
	// kb.KeyPause = 284
	kb.KeyF1 = 112
	kb.KeyF2 = 113
	kb.KeyF3 = 114
	kb.KeyF4 = 115
	kb.KeyF5 = 116
	kb.KeyF6 = 117
	kb.KeyF7 = 118
	kb.KeyF8 = 119
	kb.KeyF9 = 120
	kb.KeyF10 = 121
	kb.KeyF11 = 122
	kb.KeyF12 = 123
	// kb.KeyF13 = 302
	// kb.KeyF14 = 303
	// kb.KeyF15 = 304
	// kb.KeyF16 = 305
	// kb.KeyF17 = 306
	// kb.KeyF18 = 307
	// kb.KeyF19 = 308
	// kb.KeyF20 = 309
	// kb.KeyF21 = 310
	// kb.KeyF22 = 311
	// kb.KeyF23 = 312
	// kb.KeyF24 = 313
	// kb.KeyF25 = 314
	// kb.KeyKP0 = 320
	// kb.KeyKP1 = 321
	// kb.KeyKP2 = 322
	// kb.KeyKP3 = 323
	// kb.KeyKP4 = 324
	// kb.KeyKP5 = 325
	// kb.KeyKP6 = 326
	// kb.KeyKP7 = 327
	// kb.KeyKP8 = 328
	// kb.KeyKP9 = 329
	// kb.KeyKPDecimal = 330
	// kb.KeyKPMultiply = 332
	// kb.KeyKPSubtract = 333
	// kb.KeyKPAdd = 334
	// kb.KeyKPEnter = 335
	// kb.KeyKPEqual = 336
	kb.KeyLeftShift = 16
	kb.KeyLeftControl = 16
	kb.KeyLeftAlt = 18
	// kb.KeyLeftSuper = 343
	kb.KeyRightShift = 16
	kb.KeyRightControl = 17
	kb.KeyRightAlt = 18
	// kb.KeyRightSuper = 347
	// kb.KeyMenu = 348
	// kb.KeyLast = KeyMenu
}
