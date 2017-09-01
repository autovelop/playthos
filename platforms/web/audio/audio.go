package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"webaudio", []string{"asset_loader"}, []string{"audio_player"}, []string{"web"}})
	fmt.Println("> Audio (Web): Initializing")
}
