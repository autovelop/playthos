package keyboard

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("desktop", "keyboard")
	fmt.Println("> Keyboard: Initializing")
}
