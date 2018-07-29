/*
Package desktop adds desktop support to engine.
*/
package desktop

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPlatform("desktop", &engine.Platform{"go", []string{"build", "-v"}, "-tags", []string{}, "", ""})
	engine.RegisterPackage(&engine.Package{"desktop", []string{}, []string{"asset_loader"}, []string{"desktop"}})
	fmt.Println("> Desktop: Initializing")
}
