package engine

import (
	"fmt"
)

// Entity is used to group components with an unique identifier in order to form a virtual/game object
type Entity struct {
	*unit
	id         uint
	components []ComponentRoutine
}

// ID returns unique entity identifier
func (e *Entity) ID() uint {
	return e.id
}

// ID returns unique entity identifier
func (e *Entity) Components() []ComponentRoutine {
	return e.components
}

// AddComponent adds a new component to an entity
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

// The below method is more tedious and resource intensive but generic way to get components of any type that implements ComponentRoutine
// Component returns a ComponentRoutine based on the given Component type
// func (e *Entity) Component(lookup interface{}) ComponentRoutine {
// 	for _, component := range e.components {
//		if reflect.TypeOf(component) == reflect.TypeOf(&Transform{}) {
// 		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", lookup) {
// 			return component
// 		}
// 	}
// 	return nil
// }
