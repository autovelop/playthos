/*
Package web adds web/javascript support to engine.
*/
package web

import (
	"fmt"
	"go/build"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("web", &engine.Platform{fmt.Sprintf("%v/bin/gopherjs", build.Default.GOPATH), []string{"build"}, "--tags", []string{}, "linux", "", "", ".js", "github.com/gopherjs/gopherjs"})
	engine.RegisterPackage(&engine.Package{"web", []string{}, []string{"asset_loader"}, []string{"web"}})
	fmt.Printf("> Web: Initializing (GOPATH: %v)\n", build.Default.GOPATH)
}
