/*
Package windows adds windows support to engine.
*/
package windows

import (
	"fmt"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("windows", &engine.Platform{"go", []string{"build"}, "-tags", []string{}, "", "386", "i686-w64-mingw32-gcc", ".exe", ""})
	engine.RegisterPackage(&engine.Package{"windows", []string{}, []string{"asset_loader"}, []string{"windows"}})
	fmt.Println("> Windows: Initializing")
}
