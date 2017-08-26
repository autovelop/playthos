package glfw

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"glfw", []string{}, []string{"window_context"}, []string{"linux", "windows", "darwin"}})
	fmt.Println("> GLFW: Initializing")
}
