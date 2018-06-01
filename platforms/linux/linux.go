/*
Package linux adds linux support to engine.
*/
package linux

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("linux", &engine.Platform{"go", []string{"build", "-v"}, "-tags", []string{}, "", ""})
	engine.RegisterPackage(&engine.Package{"lin", []string{}, []string{"asset_loader"}, []string{"linux"}})
	fmt.Println("> Linux: Initializing")
}
