/*
Package windows adds windows support to engine.
*/
package windows

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("windows", &engine.Platform{"go", []string{"build"}, "-tags", []string{}, "", ".exe", ""})
	engine.RegisterPackage(&engine.Package{"windows", []string{}, []string{"asset_loader"}, []string{"windows"}})
	fmt.Println("> Windows: Initializing")
}
