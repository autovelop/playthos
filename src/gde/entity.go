package main

import (
	"fmt"
)

type Entity struct {
	components []Component
}

func (e *Entity) NewComponent(component Component) {
	e.components = append(e.components, component)
}

func (e *Entity) RegisterToSystems(engine *Engine) {
	for _, system := range engine.systems {
		for _, component_type := range system.ComponentTypes() {
			for _, component := range e.components {
				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
					component.RegisterToSystem(system)
				}
			}
		}
	}
}

// type Entity struct {
// 	Id         string
// 	Components []ComponentRoutine
// }

// func (e *Entity) Init() {
// 	// log.Printf("Entity > Init")
// 	e.Components = make([]ComponentRoutine, 0)
// }

// func (e *Entity) Add(engine *Engine) {
// 	// log.Printf("Entity > Add: %v", e.Id)
// 	engine.Entities[e.Id] = e
// }

// func (e *Entity) Get() *Entity {
// 	// log.Printf("Entity > Get: %v", e.Id)
// 	return e
// }

// func (e *Entity) AddComponent(comp ComponentRoutine) {
// 	// log.Printf("Entity > Component > Add: %T", comp)
// 	// e.Components[fmt.Sprintf("%T", comp)] = comp
// 	// e.Components[fmt.Sprintf("%T", comp)] = comp
// 	e.Components = append(e.Components, comp)
// }

// func (e *Entity) GetComponent(id string) ComponentRoutine {
// 	// log.Printf("Entity > Component > Add: %T", comp)
// 	// if e.Components[fmt.Sprintf("%T", comp)] == nil {
// 	// 	log.Printf("Component %T does not exist for Entity (ID: %v)", comp, e.Id)
// 	// 	// DO ERROR HANDLING HERE
// 	// }
// 	for _, comp := range e.Components {
// 		if comp.Id() == id {
// 			return comp
// 		}
// 	}
// 	return nil
// }
// func (e *Entity) GetComponentByStr(id string) ComponentRoutine {
// 	for _, comp := range e.Components {
// 		if comp.Id() == id {
// 			return comp
// 		}
// 	}
// 	return nil
// 	// log.Printf("Entity > Component > Add: %v", comp)
// 	// return e.Components[comp]
// }

// func (e *Entity) GetComponents() []ComponentRoutine {
// 	// log.Printf("Entity > Components > Get: %v (count)", len(e.Components))
// 	return e.Components
// }

// func (e *Entity) GetId() string {
// 	// log.Printf("Entity > Id > Get: %v", e.Id)
// 	return e.Id
// }
