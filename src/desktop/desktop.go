package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde"
	"gde/keyboard"
	"gde/opengl"
	"gde/pointer"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Greate game engine
	engine := &gde.Engine{} // Set Device, OS, and OpenGL
	engine.Init()

	// Create render system
	render := &opengl.RenderOpenGL{}
	engine.AddSystem(gde.SystemRender, render)
	render.Init()
	window := render.GetWindow()

	// Create keyboard input system
	// window.GetUserPointer().
	keyInput := &keyboard.Keyboard{Window: window}
	engine.AddSystem(gde.SystemInputKeyboard, keyInput)
	keyInput.Init()

	// Escape
	keyInput.BindOn(256, func() {
		keyInput.Window.SetShouldClose(true)
	})

	// Create pointer input system
	pointerInput := &pointer.Pointer{Window: window}
	engine.AddSystem(gde.SystemInputPointer, pointerInput)
	pointerInput.Init()

	scene := &gde.Scene{}
	scene.LoadScene(engine)
	// engine.LoadScene(&gde.Scene{})

	for true {
		engine.Update()
	}
}
