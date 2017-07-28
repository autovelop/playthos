package animation

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("desktop", "animation")
	fmt.Println("> Animation: Initializing")
}
