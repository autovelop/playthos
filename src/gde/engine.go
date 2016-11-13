package gde

import (
	"fmt"
	"log"
)

type Engine struct {
	Entities     map[string]EntityRoutine
	Systems      map[int]SystemRoutine
	RenderSystem string
}

const (
	SystemRender = iota
	SystemInput
	SystemAnimation
	SystemNetwork
	SystemPhysics
	SystemAudio
)

func (e *Engine) Init() {
	e.Entities = make(map[string]EntityRoutine)
	e.Systems = make(map[int]SystemRoutine)
}

func (e *Engine) Update() {
	for _, v := range e.Systems {
		v.Update(&e.Entities)
	}
}

func (e *Engine) Shutdown() {
	for _, v := range e.Systems {
		v.Shutdown()
	}
}

func (e *Engine) GetEntity(id string) EntityRoutine {
	fmt.Println("Engine.GetEntity(string) returned EntityRoutine")
	return e.Entities[id]
}

func (e *Engine) GetSystem(sys int) SystemRoutine {
	log.Printf("System Get: %v", sys)
	return e.Systems[sys]
}

func (e *Engine) AddSystem(sys int, sysRoutine SystemRoutine) SystemRoutine {
	log.Printf("System Added: %v", sys)
	e.Systems[sys] = sysRoutine
	return e.Systems[sys]
}

func (e *Engine) LoadScene(scene *Scene) {
	render, err := e.GetSystem(SystemRender).(RenderRoutine)
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
	player.Add(e)

	transform := &Transform{}
	transform.Init()
	transform.SetProperty("Position", Vector3{0.2, -0.5, 0})
	transform.SetProperty("Rotation", Vector3{0, 0, 45})

	player.AddComponent(transform)
	player.AddComponent(renderer)

	// Create box entity
	box := &Entity{Id: "Box"}
	box.Init()
	box.Add(e)

	box_transform := &Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", Vector3{0.2, 0.5, 0})
	box_transform.SetProperty("Rotation", Vector3{0, 0, 0})
	box.AddComponent(box_transform)

	box.AddComponent(renderer)
}
