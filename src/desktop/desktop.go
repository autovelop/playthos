package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde"
	"gde/input"
	"gde/opengl"
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

	// Create keyboard input system
	keyInput := &input.Keyboard{Window: render.Window}
	engine.AddSystem(gde.SystemInputKeyboard, keyInput)
	keyInput.Init()

	// Escape
	keyInput.Bind(256, func() {
		keyInput.Window.SetShouldClose(true)
	})

	// Create pointer input system
	mouseInput := &input.Pointer{}
	engine.AddSystem(gde.SystemInputPointer, mouseInput)
	mouseInput.Init()

	scene := &gde.Scene{}
	scene.LoadScene(engine)
	// engine.LoadScene(&gde.Scene{})

	for true {
		engine.Update()
	}
}
