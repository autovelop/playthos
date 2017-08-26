package web

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("web", &engine.Platform{"gopherjs", []string{"build"}, "--tags", []string{}, ".js"})
	engine.RegisterPackage(&engine.Package{"web", []string{}, []string{"asset_loader"}, []string{"web"}})
	fmt.Println("> Web: Initializing")
}
