package engine

import (
	"fmt"
	"log"
)

type Entity struct {
	Id         string
	components map[string]ComponentRoutine
}

func (e *Entity) Init() {
	log.Printf("Entity > Init")
	e.components = make(map[string]ComponentRoutine)
}

func (e *Entity) Add(engine *Engine) {
	log.Printf("Entity > Add: %v", e.Id)
	engine.Entities[e.Id] = e
}

func (e *Entity) Get() *Entity {
	log.Printf("Entity > Get: %v", e.Id)
	return e
}

func (e *Entity) AddComponent(comp ComponentRoutine) {
	log.Printf("Entity > Component > Add: %T", comp)
	e.components[fmt.Sprintf("%T", comp)] = comp
}

func (e *Entity) GetComponent(comp interface{}) ComponentRoutine {
	log.Printf("Entity > Component > Get: %T", comp)
	return e.components[fmt.Sprintf("%T", comp)]
}

func (e *Entity) GetComponents() map[string]ComponentRoutine {
	log.Printf("Entity > Components > Get: %v (count)", len(e.components))
	return e.components
}

func (e *Entity) GetId() string {
	log.Printf("Entity > Id > Get: %v", e.Id)
	return e.Id
}

// Still noob when I made this
// type EntityRoutine interface {
// 	GetId() string
// 	Get() *Entity
// 	Add(*Engine)
// 	AddComponent(ComponentRoutine)
// 	Component(string) ComponentRoutine
// 	Components() map[string]ComponentRoutine
// 	Init()
// }
