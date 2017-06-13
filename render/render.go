package render

import (
	"log"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("render")
	log.Println("added render to engine")
}
