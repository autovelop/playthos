/*
Package collision adds the Collision updater system with the collider component as well as helpers for collision detection
*/
package collision

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"collision", []string{}, []string{"collision"}, []string{"generic"}})
	fmt.Println("> Collision: Initializing")
}
