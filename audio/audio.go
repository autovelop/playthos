package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("desktop", "audio")
	fmt.Println("> Audio: Initializing")
}
