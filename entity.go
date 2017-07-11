package engine

import (
	"fmt"
)

type Entity struct {
	*unit
	id         uint
	components []ComponentRoutine
}

func (e *Entity) ID() uint {
	return e.id
}

func (e *Entity) NewComponent(component ComponentRoutine) {
	component.initUnit(e.engine)
	component.initComponent(e)

	e.components = append(e.components, component)
	for _, system := range systems {
		for _, component_type := range system.ComponentTypes() {
			if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
				system.NewComponent(component)
			}
		}
	}
	component.SetActive(true)
}
func (e *Entity) Component(lookup interface{}) ComponentRoutine {
	for _, component := range e.components {
		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", lookup) {
			return component
		}
	}
	return nil
}
