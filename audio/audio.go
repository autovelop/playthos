package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("autovelop_playthos_audio")
	fmt.Println("> Audio: Initializing")
}
