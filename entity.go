package engine

import (
	"fmt"
	"log"
)

type Entity struct {
	components []ComponentRoutine
	active     bool
	ID         uint
}

func (e *Entity) NewComponent(component ComponentRoutine) {
	log.Println("NewComponent")
	component.SetEntity(e)
	e.components = append(e.components, component)
}

func (e *Entity) SetActive(active bool) {
	e.active = active
}

func (e *Entity) IsActive() bool {
	return e.active
}

func (e *Entity) NewComponents(components ...ComponentRoutine) {
	for _, component := range components {
		component.SetEntity(e)
	}
	e.components = append(e.components, components...)
}

func (e *Entity) Components() []ComponentRoutine {
	return e.components
}

func (e *Entity) DeleteComponents() {
	e.components = []ComponentRoutine{}
}

func (e *Entity) GetComponent(component_type interface{}) ComponentRoutine {
	for _, component := range e.components {
		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
			return component
		}
	}
	return nil
}

func (e *Entity) UnRegisterAllFromSystems(engine *Engine) {
	// log.Printf("+ !! %v\n", len(e.components))
	for _, system := range engine.systems {
		for _, component := range e.components {
			for _, component_type := range system.ComponentTypes() {
				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
					// log.Printf("+ DEL %T\n", component_type)
					// component.UnRegisterFromSystem(system)
					return
				}
			}
			// }
		}
	}
}

func (e *Entity) RegisterAllToSystems(engine *Engine) {
	for _, system := range engine.systems {
		for _, component_type := range system.ComponentTypes() {
			for _, component := range e.components {
				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
					// component.RegisterToSystem(system)
					system.LoadComponent(component)
					// return
				}
			}
		}
	}
}

func (e *Entity) RegisterToSystems(engine *Engine, components ...ComponentRoutine) {
	for _, system := range engine.systems {
		for _, component_type := range system.ComponentTypes() {
			for _, component := range components {
				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
					// component.RegisterToSystem(system)
					system.LoadComponent(component)
					// return
				}
			}
		}
	}
}
