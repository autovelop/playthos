package gde

import (
	"fmt"
	"log"
)

type Engine struct {
	Entities map[string]EntityRoutine
	Systems  map[string]SystemRoutine
}

// type EngineRoutine interface {
// 	Init()
// 	Update()
// 	GetEntity() EntityRoutine
// 	GetSystem(interface{}) SystemRoutine
// }

func (e *Engine) Init() {
	// fmt.Println("Engine.Init() executed")
	e.Entities = make(map[string]EntityRoutine)
	e.Systems = make(map[string]SystemRoutine)
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

func (e *Engine) GetSystem(id interface{}) SystemRoutine {
	fmt.Println("Engine.GetSystem(interface{}) returned SystemRoutine")
	log.Printf("%T", id)
	return e.Systems[fmt.Sprintf("%T", id)]
}
