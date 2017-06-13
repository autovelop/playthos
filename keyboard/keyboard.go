package keyboard

import (
	"github.com/autovelop/playthos"
	"log"
)

func init() {
	engine.RegisterPackage("desktop", "keyboard")
	log.Println("added keyboard to engine")
}
