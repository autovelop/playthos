package engine

import (
	"fmt"
)

type Entity struct {
	*unit
	id         uint
	components []ComponentRoutine
	// id uint
	// engine *Engine
}

// func (e *Entity) Init(id uint) {
// 	e.id = id
// }

func (e *Entity) ID() uint {
	return e.id
}

// func (e *Entity) set(id uint, engine *Engine) {
// 	e.id = id
// 	e.engine = engine
// }

// func (e *Entity) ID() uint {
// 	return e.id
// }

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

// TODO
// func (e *Entity) DeleteComponent(component ComponentRoutine) {
// component.SetEntity(e)
// e.components = append(e.components, component)
// for _, system := range engine.systems {
// 	for _, component_type := range system.ComponentTypes() {
// 		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
// 			system.DeleteComponent(component)
// 		}
// 	}
// }
// }

// STIL NOT SURE IF THE IS NECEASSRY
// func (e *Entity) DeleteComponents() {
// 	e.components = []ComponentRoutine{}
// }

// func (e *Entity) SetActive(active bool) {
// 	e.active = active
// }

// func (e *Entity) Active() bool {
// 	return e.active
// }

// func (e *Entity) NewComponents(components ...ComponentRoutine) {
// 	for _, component := range components {
// 		component.SetEntity(e)
// 	}
// 	e.components = append(e.components, components...)
// }

// func (e *Entity) Components() []ComponentRoutine {
// 	return e.components
// }

func (e *Entity) Component(lookup interface{}) ComponentRoutine {
	for _, component := range e.components {
		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", lookup) {
			return component
		}
	}
	return nil
}

// func (e *Entity) UnRegisterAllFromSystems(engine *Engine) {
// 	// log.Printf("+ !! %v\n", len(e.components))
// 	for _, system := range engine.systems {
// 		system.UnRegisterEntity(e)
// 	}
// }

// func (e *Entity) RegisterAllToSystems(engine *Engine) {
// 	for _, system := range engine.systems {
// 		for _, component_type := range system.ComponentTypes() {
// 			for _, component := range e.components {
// 				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
// 					// component.RegisterToSystem(system)
// 					system.LoadComponent(component)
// 					// return
// 				}
// 			}
// 		}
// 	}
// }

// func (e *Entity) RegisterToSystems(engine *Engine, components ...ComponentRoutine) {
// 	for _, system := range engine.systems {
// 		for _, component_type := range system.ComponentTypes() {
// 			for _, component := range components {
// 				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
// 					// component.RegisterToSystem(system)
// 					system.LoadComponent(component)
// 					// return
// 				}
// 			}
// 		}
// 	}
// }
