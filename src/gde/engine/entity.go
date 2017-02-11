package engine

import (
	"fmt"
	// "log"
)

type Entity struct {
	Id         string
	Components map[string]ComponentRoutine
}

func (e *Entity) Init() {
	// log.Printf("Entity > Init")
	e.Components = make(map[string]ComponentRoutine)
}

func (e *Entity) Add(engine *Engine) {
	// log.Printf("Entity > Add: %v", e.Id)
	engine.Entities[e.Id] = e
}

func (e *Entity) Get() *Entity {
	// log.Printf("Entity > Get: %v", e.Id)
	return e
}

func (e *Entity) AddComponent(comp ComponentRoutine) {
	// log.Printf("Entity > Component > Add: %T", comp)
	e.Components[fmt.Sprintf("%T", comp)] = comp
}

func (e *Entity) GetComponent(comp interface{}) ComponentRoutine {
	// if e.Components[fmt.Sprintf("%T", comp)] == nil {
	// 	log.Printf("Component %T does not exist for Entity (ID: %v)", comp, e.Id)
	// 	// DO ERROR HANDLING HERE
	// }
	return e.Components[fmt.Sprintf("%T", comp)]
}

func (e *Entity) GetComponents() map[string]ComponentRoutine {
	// log.Printf("Entity > Components > Get: %v (count)", len(e.Components))
	return e.Components
}

func (e *Entity) GetId() string {
	// log.Printf("Entity > Id > Get: %v", e.Id)
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
