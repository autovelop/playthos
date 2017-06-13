package glfw

import (
	"github.com/autovelop/playthos"
	"log"
)

func init() {
	engine.RegisterPackage("glfw")
	log.Println("added glfw to engine")
}
