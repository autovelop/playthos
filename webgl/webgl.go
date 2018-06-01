/*
Package webgl adds the WebGL drawer system.
*/
package webgl

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"webgl", []string{"asset_loader"}, []string{"window_context", "drawing"}, []string{"web"}})
	fmt.Println("> WebGL: Initializing")
}
