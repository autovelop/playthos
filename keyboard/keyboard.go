/*
Package keyboard adds a empty Keyboard listener integrant to be overwriten by the platform specific integrant.
*/
package keyboard

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"keyboard", []string{"window_context"}, []string{"keyboard_input"}, []string{"generic"}})
	fmt.Println("> Keyboard: Initializing")
}
