package gde

import (
	"log"
)

// Loads data from file/db
type Scene struct {
	name string
}

func (s *Scene) LoadScene(engine *Engine) {
	render, err := engine.GetSystem(SystemRender).(RenderRoutine)
	if !err {
		log.Println(err)
		return
	}

	// Simple Quad mesh renderer
	renderer := &Renderer{}
	renderer.Init()
	renderer.LoadMesh(&Mesh{
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
	player := &Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	transform := &Transform{}
	transform.Init()
	transform.SetProperty("Position", Vector3{0.2, -0.5, 0})
	transform.SetProperty("Rotation", Vector3{0, 0, 45})

	player.AddComponent(transform)
	player.AddComponent(renderer)

	// Create box entity
	box := &Entity{Id: "Box"}
	box.Init()
	box.Add(engine)

	box_transform := &Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", Vector3{0.2, 0.5, 0})
	box_transform.SetProperty("Rotation", Vector3{0, 0, 0})
	box.AddComponent(box_transform)

	box.AddComponent(renderer)

	// Create box entity 2
	box2 := &Entity{Id: "Box2"}
	box2.Init()
	box2.Add(engine)

	box2_transform := &Transform{}
	box2_transform.Init()
	box2_position := Vector3{-0.2, 0.5, 0}
	box2_transform.SetProperty("Position", box2_position)
	box2_transform.SetProperty("Rotation", Vector3{0, 0, 0})
	box2.AddComponent(box2_transform)

	box2.AddComponent(renderer)

	keyInput, err := engine.GetSystem(SystemInputKeyboard).(InputRoutine)
	if !err {
		log.Println(err)
		return
	}

	// Right arrow
	keyInput.Bind(262, func() {
		box2_position.X += 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Left arrow
	keyInput.Bind(263, func() {
		box2_position.X -= 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Up arrow
	keyInput.Bind(265, func() {
		box2_position.Y += 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Down arrow
	keyInput.Bind(264, func() {
		box2_position.Y -= 0.1
		box2_transform.SetProperty("Position", box2_position)
	})
}
