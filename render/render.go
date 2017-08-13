package render

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("autovelop_playthos_render")
	fmt.Println("> Render: Ready")
}
