// +build glfw

package glfw

import (
	"github.com/autovelop/playthos"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	engine.NewIntegrant(&GLFW{})
	log.Println("added glfw integrant to engine")
}

type GLFW struct {
	engine.Integrant
	window *glfw.Window
}

func (g *GLFW) InitIntegrant() {
	settings := g.Engine().Settings()

	log.Println("GLFW Prepare")
	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw")
		glfw.Terminate()
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)

	var err error
	if settings.Fullscreen {
		g.window, err = glfw.CreateWindow(int(settings.ResolutionX), int(settings.ResolutionY), "Cube", glfw.GetPrimaryMonitor(), nil)
	} else {
		g.window, err = glfw.CreateWindow(int(settings.ResolutionX), int(settings.ResolutionY), "Cube", nil, nil)
	}
	if err != nil {
		panic(err)
	}
	g.window.MakeContextCurrent()
}

func (g *GLFW) Window() *glfw.Window {
	return g.window
}
