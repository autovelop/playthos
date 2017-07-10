package scripting

import (
	"log"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("scripting")
	log.Println("added script to engine")
}
