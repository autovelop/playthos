// +build deploy glfw

package glfw

import (
	"fmt"
	"log"
	"runtime"

	"github.com/autovelop/playthos"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
	engine.NewIntegrant(&GLFW{})
	fmt.Println("> GLFW: Ready")
}

// Don't like this at all
var thisglfw *GLFW

// GLFW defines the window and monitor of a rendering application
type GLFW struct {
	engine.Integrant
	window       *glfw.Window
	monitor      *glfw.Monitor
	settings     *engine.Settings
	majorVersion int
	minorVersion int
}

// OpenGLVersion returns the major and minor versions of opengl on the supported operating system
func (g *GLFW) OpenGLVersion() (int, int) {
	return g.majorVersion, g.minorVersion
}

// InitIntegrant called when the integrant plugs into the engine
func (g *GLFW) InitIntegrant() {
	// start at the top and work our way down
	// g.majorVersion = 3
	// g.minorVersion = 3
	if g.majorVersion == 0 && g.minorVersion == 0 {
		g.majorVersion = 4
		g.minorVersion = 5
	}
	g.settings = g.Engine().Settings()

	// log.Printf("GLFW Prepare (%v.%v)\n")
	fmt.Printf("> GLFW: Version %v.%v\n", g.majorVersion, g.minorVersion)
	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		glfw.Terminate()
		log.Fatalf("Failed to initialize GLFW: %v", err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, g.majorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, g.minorVersion)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.Focused, glfw.True)

	// Don't know why this isn't working. TODO: report to go-glfw
	// glfw.WindowHint(glfw.Iconified, glfw.True)

	// Might need this line for MacOS true fullscreen
	// glfw.WindowHint(glfw.Decorated, glfw.False)

	glfw.WindowHint(glfw.Resizable, glfw.True)
	// glfw.WindowHint(glfw.Samples, 4)

	var err error
	if g.settings.Fullscreen {
		fmt.Printf("> GLFW: Fullscreen = Yes\n")
		glfw.WindowHint(glfw.Maximized, glfw.True)
		glfw.WindowHint(glfw.Floating, glfw.True)
		glfw.WindowHint(glfw.AutoIconify, glfw.True)
		g.monitor = glfw.GetPrimaryMonitor()
		if int(g.settings.ResolutionX) <= 0 {
			vidMode := g.monitor.GetVideoMode()
			g.settings.ResolutionX = float32(vidMode.Width)
			g.settings.ResolutionY = float32(vidMode.Height)
		}
	} else {
		fmt.Printf("> GLFW: Fullscreen = No\n")
		// Sometimes when the same process run new GLFW instances sequentially, these properties are not reset. Therefore we have to explicitly set them.
		glfw.WindowHint(glfw.Maximized, glfw.False)
		glfw.WindowHint(glfw.Floating, glfw.False)
		glfw.WindowHint(glfw.AutoIconify, glfw.False)
		g.monitor = nil
		if int(g.settings.ResolutionX) <= 0 {
			g.settings.ResolutionX = float32(800)
			g.settings.ResolutionY = float32(600)
		}
	}
	fmt.Printf("> GLFW: Resolution = %vx%v | Monitor = %v\n", g.settings.ResolutionX, g.settings.ResolutionY, g.monitor)
	g.window, err = glfw.CreateWindow(int(g.settings.ResolutionX), int(g.settings.ResolutionY), "Game", g.monitor, nil)
	if err != nil {
		log.Printf("An error occured with the active version of GLFW.\nerr:%v", err)
		switch g.majorVersion {
		case 4:
			switch g.minorVersion {
			case 5:
				g.minorVersion = 3
				break
			case 3:
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
		g.InitIntegrant()
	}

	if g.settings.Cursor {
		// g.window.SetCursorPosCallback(onCursorMove)
	} else {
		g.window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	}

	g.window.MakeContextCurrent()

	thisglfw = g
}

// Destroy called when engine is gracefully shutting down
func (g *GLFW) Destroy() {
	// defer glfw.Terminate()
	// log.Fatal("here")
	g.SetActive(false)
	if g.settings.Fullscreen {
		g.window.SetMonitor(nil, 0, 0, int(g.settings.ResolutionX), int(g.settings.ResolutionY), 0)
	}
	g.window.SetShouldClose(true)
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (g *GLFW) AddIntegrant(engine.IntegrantRoutine) {
}

// Window returns pointer to the current glfw window
func (g *GLFW) Window() *glfw.Window {
	return g.window
}

// func onCursorMove(w *glfw.Window, x float64, y float64) {
// need to revisit this at some point. need to prevent OSX from showing the menu drawer
// if y <= 4 {
// 	w.SetCursorPos(x, 5)
// }
// }
