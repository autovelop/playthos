package physics

import (
	"log"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("physics")
	log.Println("added physics to engine")
}
