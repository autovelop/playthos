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
	SystemInputKeyboard
	SystemInputPointer
	SystemInputTouch
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
		v.Stop()
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
