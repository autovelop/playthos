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
			0.1, 0.1, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0,
			0.1, -0.1, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
			-0.1, -0.1, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
			-0.1, 0.1, 0.0, 0.0, 1.0, 1.0, 0.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	texture := &Texture{FilePath: "weapon.png"}
	texture.ReadTexture()
	renderer.LoadTexture(texture)
	render.LoadRenderer(renderer)

	// Create player entity
	player := &Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	transform := &Transform{}
	transform.Init()
	transform.SetProperty("Position", Vector3{0.5, 1.0, 0})
	transform.SetProperty("Rotation", Vector3{0, 0, 45})

	player.AddComponent(transform)
	player.AddComponent(renderer)

	// Create box entity
	box := &Entity{Id: "Box"}
	box.Init()
	box.Add(engine)

	box_transform := &Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", Vector3{0.2, 1.8, 0})
	box_transform.SetProperty("Rotation", Vector3{0, 0, 0})

	box.AddComponent(box_transform)
	box.AddComponent(renderer)

	// Create box entity 2
	box2 := &Entity{Id: "Box2"}
	box2.Init()
	box2.Add(engine)

	box2_transform := &Transform{}
	box2_transform.Init()
	box2_position := Vector3{0.1, 0.1, 0}
	box2_transform.SetProperty("Position", box2_position)
	box2_transform.SetProperty("Rotation", Vector3{0, 0, 0})

	box2.AddComponent(box2_transform)
	box2.AddComponent(renderer)

	// Lets test keyboard support
	keyInput, err := engine.GetSystem(SystemInputKeyboard).(Input)
	if !err {
		log.Println(err)
		return
	}

	// Right arrow
	keyInput.BindOn(262, func() {
		box2_position.X += 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Left arrow
	keyInput.BindOn(263, func() {
		box2_position.X -= 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Up arrow
	keyInput.BindOn(265, func() {
		box2_position.Y += 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Down arrow
	keyInput.BindOn(264, func() {
		box2_position.Y -= 0.1
		box2_transform.SetProperty("Position", box2_position)
	})

	// Lets test pointer support
	pointerInput, err := engine.GetSystem(SystemInputPointer).(Input)
	if !err {
		log.Printf("Pointer Input system not started/found\nERROR:\n%v\n\n", err)
		return
	}

	// pointer Move
	pointerInput.BindMove(func(x float64, y float64) {
		// log.Printf("%v, %v", x, y)
		box2_position.X = float32(x/360) * 1
		box2_position.Y = float32(y/640) * 2
		box2_transform.SetProperty("Position", box2_position)
	})

	// // Left click
	// pointerInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Right click
	// pointerInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Lets test touch support
	// touchInput, err := engine.GetSystem(SystemInputTouch).(InputRoutine)
	// if !err {
	// 	log.Printf("Touch Input system not started/found\nERROR:\n%v\n\n", err)
	// 	return
	// }

	// // Touch down
	// touchInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Touch up
	// touchInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })
}
