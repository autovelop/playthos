/*
Package opengl adds the OpenGL drawer system.
*/
package opengl

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"opengl", []string{"window_context", "asset_loader"}, []string{"drawing"}, []string{"linux", "windows", "darwin"}})
	fmt.Println("> OpenGL: Initializing")
}
