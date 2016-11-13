package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"gde"
	"gde/components"
	"gde/geometry"
	"gde/systems/opengl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Greate game engine
	engine := &gde.Engine{} // Set Device, OS, and OpenGL
	engine.Init()

	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Create render system
	render := &opengl.RenderOpenGL{}
	render.Add(engine)
	render.Init()

	// Simple Quad mesh renderer
	renderer := &components.Renderer{}
	renderer.Init()
	renderer.LoadMesh(&geometry.Mesh{
		Vertices: []float32{
			0.1, 0.1, 0.0,
			0.1, -0.1, 0.0,
			-0.1, -0.1, 0.0,
			-0.1, 0.1, 0.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	render.LoadRenderer(renderer)

	// Create player entity
	player := &gde.Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	transform := &components.Transform{}
	transform.Init()
	transform.SetProperty("Position", mgl32.Vec3{0.2, -0.5, 0})
	transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 45})

	player.AddComponent(transform)
	player.AddComponent(renderer)

	// Create box entity
	box := &gde.Entity{Id: "Box"}
	box.Init()
	box.Add(engine)

	box_transform := &components.Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", mgl32.Vec3{0.2, 0.5, 0})
	box_transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 0})
	box.AddComponent(box_transform)

	box.AddComponent(renderer)

	for true {
		engine.Update()
	}
}
