package animation

import (
	"github.com/autovelop/playthos"
	"log"
)

func init() {
	engine.RegisterPackage("desktop", "animation")
	log.Println("added animation to engine")
}
