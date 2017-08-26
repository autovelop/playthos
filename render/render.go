package render

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"render", []string{"window_context", "drawing"}, []string{"render"}, []string{"generic"}})
	fmt.Println("> Render: Initializing")
}
