package engine

import (
	"fmt"
	"log"
	// "strings"
	"time"
)

var package_tags []string
var unloaded_systems []System
var unloaded_observerables []Observerable
var unloaded_components []ComponentRoutine

// var Game *Engine

func init() {
	log.Println("init engine")
	// Game = &Engine{}
	// Game.entities = make([]*Entity, 0)
	// Game.ready = true

	package_tags = make([]string, 0)
}

func RegisterPackage(tags ...string) {
	package_tags = append(package_tags, tags...)
}

// func GetTags() string {
// 	return strings.Join(tags[:], " ")
// }

type Engine struct {
	systems       []System
	observerables []Observerable
	entities      []*Entity
	components    []ComponentRoutine
	newTime       time.Time
	currentTime   time.Time
	settings      *Settings
	accumulator   int64
	deltaTime     int64
	frames        uint64
	ready         bool
	running       bool
}

func New(settings *Settings) *Engine {
	game := &Engine{}
	game.settings = settings
	game.LoadComponents()
	game.LoadObserverables()
	game.LoadSystems()
	return game
}

func (e *Engine) Update() {
	log.Println("Engine Update")
	e.newTime = time.Now()
	frameTime := e.newTime.Sub(e.currentTime).Nanoseconds()
	e.currentTime = e.newTime
	e.accumulator += frameTime

	for e.accumulator >= e.deltaTime {
		for _, system := range e.systems {
			system.Update()
		}
	}
	e.accumulator -= e.deltaTime
}

func (e *Engine) Run() {
	e.Update()
}

// func Run() {
// 	if Game.ready == false {
// 		log.Println("Engine must be ready before it can start")
// 	} else {
// 		log.Println("Engine Running")
// 		Game.running = true
// 		for Game.running == true {
// 			Game.Update()
// 		}
// 	}
// }

// func Stop() {
// 	if Game.running == false {
// 		log.Println("Engine cannot be stopped if it is not running")
// 	} else {
// 		log.Println("Engine Stopped")
// 		Game.running = false
// 	}
// }

func (e *Engine) NewComponent(component ComponentRoutine) {
	component.Prepare(e.settings)
	e.components = append(e.components, component)
}

func (e *Engine) LoadObserverables() {
	// there should be no systems loaded here
	for _, observerable := range unloaded_observerables {
		e.observerables = append(e.observerables, observerable)
	}
	unloaded_observerables = []Observerable{}

	for _, observerable := range e.observerables {
		observerable.Prepare(e.settings)
		for _, component := range e.components {
			observerable.LoadComponent(component)
		}
	}
}
func (e *Engine) LoadSystems() {
	// there should be no systems loaded here
	for _, system := range unloaded_systems {
		e.systems = append(e.systems, system)
	}
	unloaded_systems = []System{}

	for _, system := range e.systems {
		system.Prepare(e.settings)
		for _, component := range e.components {
			system.LoadComponent(component)
		}
	}
}

func (e *Engine) LoadComponents() {
	// there should be no components loaded here
	for _, component := range unloaded_components {
		e.components = append(e.components, component)
	}
	unloaded_components = []ComponentRoutine{}

	for _, component := range e.components {
		component.Prepare(e.settings)
	}
}

func NewUnloadedObserverable(observerable Observerable) {
	unloaded_observerables = append(unloaded_observerables, observerable)
}
func NewUnloadedSystem(system System) {
	unloaded_systems = append(unloaded_systems, system)
}

func NewUnloadedComponent(component ComponentRoutine) {
	unloaded_components = append(unloaded_components, component)
}

var idnum uint = 0

func (e *Engine) NewEntity() *Entity {
	idnum++
	e.entities = append(e.entities, &Entity{ID: idnum})
	return e.entities[len(e.entities)-1]
}

func (e *Engine) GetEntity(id uint) *Entity {
	for _, entity := range e.entities {
		if entity.ID == id {
			return entity
		}
	}
	return nil
}

func (e *Engine) DeleteEntity(ent *Entity) {
	for idx, entity := range e.entities {
		log.Println("delete")
		if entity == ent {
			entity.UnRegisterAllFromSystems(e)
			entity.DeleteComponents()

			copy(e.entities[idx:], e.entities[idx+1:])
			e.entities[len(e.entities)-1] = nil
			e.entities = e.entities[:len(e.entities)-1]
			break
		}
	}
}

func (e *Engine) GetObserverable(lookup Observerable) Observerable {
	for _, observerable := range e.observerables {
		if fmt.Sprintf("%T", observerable) == fmt.Sprintf("%T", observerable) {
			return observerable
		}
	}
	log.Fatalf("%T - System requested but doens't exist. Make sure all packages are imported", lookup)
	return nil
}

func (e *Engine) GetSystem(lookup System) System {
	for _, system := range e.systems {
		if fmt.Sprintf("%T", system) == fmt.Sprintf("%T", system) {
			return system
		}
	}
	log.Fatal("System requested but doens't exist. Make sure all packages are imported")
	return nil
}
