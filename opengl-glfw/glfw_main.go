// +build glfw
// +build linux windows darwin

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
	window  *glfw.Window
	monitor *glfw.Monitor
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
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)

	var err error
	if settings.Fullscreen {
		g.monitor = glfw.GetPrimaryMonitor()
		if int(settings.ResolutionX) <= 0 {
			vidMode := g.monitor.GetVideoMode()
			settings.ResolutionX = float32(vidMode.Width)
			settings.ResolutionY = float32(vidMode.Height)
		}
	} else {
		if int(settings.ResolutionX) <= 0 {
			settings.ResolutionX = float32(800)
			settings.ResolutionY = float32(600)
		}
	}
	g.window, err = glfw.CreateWindow(int(settings.ResolutionX), int(settings.ResolutionY), "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	g.window.MakeContextCurrent()
}

func (g *GLFW) DeleteIntegrant() {
	settings := g.Engine().Settings()
	g.window.SetMonitor(nil, 0, 0, int(settings.ResolutionX), int(settings.ResolutionY), 0)
	g.window.SetShouldClose(true)
}

func (g *GLFW) Window() *glfw.Window {
	return g.window
}
