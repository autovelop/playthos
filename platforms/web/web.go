/*
Package web adds web/javascript support to engine.
*/
package web

import (
	"fmt"
	"github.com/autovelop/playthos"
	"os"
)

func init() {
	engine.RegisterPlatform("web", &engine.Platform{fmt.Sprintf("%v/bin/gopherjs", os.Getenv("GOPATH")), []string{"build"}, "--tags", []string{}, ".js", "github.com/gopherjs/gopherjs"})
	engine.RegisterPackage(&engine.Package{"web", []string{}, []string{"asset_loader"}, []string{"web"}})
	fmt.Printf("> Web: Initializing (GOPATH: %v)\n", os.Getenv("GOPATH"))
}
