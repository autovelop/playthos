/*
Package animation adds the Animation updater system with the clip and keyframe components
*/
package animation

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"animation", []string{}, []string{"animation"}, []string{"generic"}})
	fmt.Println("> Animation: Initializing")
}
