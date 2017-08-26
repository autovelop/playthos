package animation

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"animation", []string{}, []string{"animation"}, []string{"generic"}})
	fmt.Println("> Animation: Initializing")
}
