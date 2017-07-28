package opengl

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage("desktop", "opengl")
	fmt.Println("> OpenGL: Initializing")
}
