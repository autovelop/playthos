/*
Package physics adds the Physics updater system with the rigidbody component.
*/
package physics

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"physics", []string{}, []string{"physics"}, []string{"generic"}})
	fmt.Println("> Physics: Initializing")
}
