/*
Package scripting adds the Scripting updater system.
*/
package scripting

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"scripting", []string{}, []string{"scripting"}, []string{"generic"}})
	fmt.Println("> Scripting: Initializing")
}
