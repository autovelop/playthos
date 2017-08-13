package scripting

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("autovelop_playthos_scripting")
	fmt.Println("> Scripting: Initializing")
}
