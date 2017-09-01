package keyboard

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"webkeyboard", []string{"window_context"}, []string{"keyboard_input"}, []string{"web"}})
	fmt.Println("> Keyboard (Web): Initializing")
}
