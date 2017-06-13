// +build desktop,linux,opengl,render,glfw

package glfw

import (
	"github.com/autovelop/playthos"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func init() {
	engine.NewUnloadedComponent(&GLFW{})
	log.Println("added glfw comp to engine")
}

type GLFW struct {
	engine.ComponentRoutine
	window *glfw.Window
}

func (g *GLFW) Prepare(settings *engine.Settings) {
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

func (g *GLFW) GetWindow() *glfw.Window {
	return g.window
}

func (g *GLFW) RegisterToSystem(system engine.System) {
	// log.Println("Registering GLFW to system")
	system.LoadComponent(g)
}

func (g *GLFW) RegisterToObserverable(observer engine.Observerable) {
	// log.Println("Registering GLFW to observer")
	observer.LoadComponent(g)
}
