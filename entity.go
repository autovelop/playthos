package engine

import (
	"fmt"
)

type Entity struct {
	*unit
	id         uint
	components []ComponentRoutine
}

// ID returns unique entity identifier
func (e *Entity) ID() uint {
	return e.id
}

func (e *Entity) AddComponent(component ComponentRoutine) {
	if play {
		component.initUnit(e.engine)
		component.initComponent(e)

		e.components = append(e.components, component)
		for _, system := range systems {
			for _, componentType := range system.ComponentTypes() {
				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", componentType) {
					system.AddComponent(component)
				}
			}
		}
		component.SetActive(true)
	}
}

// Component returns ComponentRoutine of empty component given as parameter
func (e *Entity) Component(lookup interface{}) ComponentRoutine {
	for _, component := range e.components {
		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", lookup) {
			return component
		}
	}
	return nil
}
