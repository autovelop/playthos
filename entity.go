package engine

import (
	"fmt"
)

type Entity struct {
	components []Component
}

func (e *Entity) NewComponent(component Component) {
	e.components = append(e.components, component)
}

func (e *Entity) GetComponent(component Component) {
	here
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
