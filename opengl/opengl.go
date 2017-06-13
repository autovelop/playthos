package opengl

import (
	"github.com/autovelop/playthos"
	"log"
)

func init() {
	engine.RegisterPackage("desktop", "opengl")
	log.Println("added opengl to engine")
}
