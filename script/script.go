package script

import (
	"log"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("script")
	log.Println("added script to engine")
}
