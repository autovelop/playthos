package scripting

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("scripting")
	fmt.Println("> Scripting: Initializing")
}
