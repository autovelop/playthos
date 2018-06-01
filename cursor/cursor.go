/*
Package cursor adds the Cursor updater system.
*/
package cursor

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"cursor", []string{"window_context"}, []string{"cursor_input"}, []string{"windows", "linux", "darwin"}})
	fmt.Println("> Cursor: Initializing")
}
