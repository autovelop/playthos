package collision

import (
	"log"

	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("collision")
	log.Println("added collision to engine")
}
