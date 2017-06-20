package audio

import (
	"github.com/autovelop/playthos"
	"log"
)

func init() {
	engine.RegisterPackage("desktop", "audio")
	log.Println("added audio to engine")
}
