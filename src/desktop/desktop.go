package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"gde"
	components "gde/components"
	systems "gde/systems"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// Greate game engine
	engine := &gde.Engine{}
	engine.Init()

	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Create render system
	var render gde.SystemRoutine
	render = &systems.Render{}
	render.Init()
	render.Add(engine)

	// Create player entity
	var player gde.EntityRoutine
	player = &gde.Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	var transform gde.ComponentRoutine
	transform = &components.Transform{}
	transform.Init()
	transform.SetProperty("Position", mgl32.Vec3{0.2, -0.5, 0})
	transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 45})
	player.AddComponent(transform)

	var quad_renderer gde.ComponentRoutine
	quad_renderer = &components.Renderer{}
	quad_renderer.Init()
	player.AddComponent(quad_renderer)

	// Create box entity
	var box gde.EntityRoutine
	box = &gde.Entity{Id: "Box"}
	box.Init()
	box.Add(engine)

	var box_transform gde.ComponentRoutine
	box_transform = &components.Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", mgl32.Vec3{0.2, 0.5, 0})
	box_transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 0})
	box.AddComponent(box_transform)

	box.AddComponent(quad_renderer)

	for true {
		engine.Update()
	}
}
