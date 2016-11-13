package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"

	"gde"
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

	engine.LoadScene(&gde.Scene{})

	for true {
		engine.Update()
	}
}
