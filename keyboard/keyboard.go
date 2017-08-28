package keyboard

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"keyboard", []string{"window_context"}, []string{"keyboard_input"}, []string{"windows", "linux", "darwin"}})
	fmt.Println("> Keyboard: Initializing")
}
