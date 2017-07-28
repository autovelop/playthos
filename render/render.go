package render

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("render")
	fmt.Println("> Render: Ready")
}
