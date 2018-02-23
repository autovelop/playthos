package windows

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("windows", &engine.Platform{"go", []string{"build", "-v"}, "-tags", []string{}, ".exe", ""})
	engine.RegisterPackage(&engine.Package{"win", []string{}, []string{"asset_loader"}, []string{"windows"}})
	fmt.Println("> Windows: Initializing")
}
