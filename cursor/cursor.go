package cursor

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("desktop", "cursor")
	fmt.Println("> Cursor: Initializing")
}
