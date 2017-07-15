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
	window       *glfw.Window
	monitor      *glfw.Monitor
	majorVersion int
	minorVersion int
}

func (g *GLFW) OpenGLVersion() (int, int) {
	return g.majorVersion, g.minorVersion
}

func (g *GLFW) InitIntegrant() {
	// start at the top and work our way down
	// g.majorVersion = 3
	// g.minorVersion = 3
	if g.majorVersion == 0 && g.minorVersion == 0 {
		g.majorVersion = 4
		g.minorVersion = 5
	}
	settings := g.Engine().Settings()

	log.Printf("GLFW Prepare (%v.%v)\n", g.majorVersion, g.minorVersion)
	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw")
		glfw.Terminate()
	}
	glfw.WindowHint(glfw.ContextVersionMajor, g.majorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, g.minorVersion)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.Focused, glfw.True)

	// Don't know why this isn't working. TODO: report to go-glfw
	// glfw.WindowHint(glfw.Iconified, glfw.True)

	glfw.WindowHint(glfw.Decorated, glfw.False)

	glfw.WindowHint(glfw.Resizable, glfw.True)

	var err error
	if settings.Fullscreen {
		glfw.WindowHint(glfw.Maximized, glfw.True)
		glfw.WindowHint(glfw.Floating, glfw.True)
		glfw.WindowHint(glfw.AutoIconify, glfw.True)
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
		switch g.majorVersion {
		case 4:
			switch g.minorVersion {
			case 5:
				g.minorVersion = 1
				break
			case 1:
				g.majorVersion = 3
				g.minorVersion = 3
				break
			}
			break
		case 3:
			log.Fatalf("Playthos doesn't support OpenGL version older than v3.3\nerr:%v", err)
			// panic(err)
			break
		}
		// g.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
		g.InitIntegrant()
	}
	g.window.MakeContextCurrent()
}

func (g *GLFW) DeleteIntegrant() {

	settings := g.Engine().Settings()
	if settings.Fullscreen {
		g.window.SetMonitor(nil, 0, 0, int(settings.ResolutionX), int(settings.ResolutionY), 0)
	}
	g.window.SetShouldClose(true)
	// glfw.Terminate()
}

func (g *GLFW) Window() *glfw.Window {
	return g.window
}
