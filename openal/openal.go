/*
Package openal adds the OpenAL listener integrant.
*/
package openal

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"openal", []string{"asset_loader"}, []string{"audio_player"}, []string{"linux", "windows", "darwin"}})
	fmt.Println("> OpenAL: Initializing")
}
