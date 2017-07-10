package engine

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var packages []string
var systems []SystemRoutine
var updaters []Updater
var listeners []Listener
var integrants []IntegrantRoutine

func init() {
	log.Println("init engine")
}

func RegisterPackage(tags ...string) {
	packages = append(packages, tags...)
}

func GetTags() string {
	return strings.Join(packages[:], " ")
}

type Engine struct {
	// systems []System
	// observerables []Observerable
	entities []*Entity
	updaters []Updater
	settings *Settings
	running  bool

	newTime     time.Time
	currentTime time.Time
	accumulator int64
	deltaTime   int64
	frames      uint64
}

// New - creates and loads engine with empty settings
func New(s *Settings) *Engine {
	game := &Engine{}
	game.SetSettings(s)
	game.Init()
	return game
}

// Start - starts engine
func (e *Engine) Start() {
	e.running = true
	e.update()
}

// Stop - stops systems and then engine
func (e *Engine) Stop() {
	e.running = false
}

func (e *Engine) Init() {
	for _, integrant := range integrants {
		integrant.initUnit(e)
		integrant.InitIntegrant()
	}
	for _, system := range systems {
		for _, integrant := range integrants {
			system.NewIntegrant(integrant)
		}
		system.initUnit(e)
		system.InitSystem()
	}
}

// NewEntity - creates new entity
var eid uint = 0

func (e *Engine) NewEntity() *Entity {
	eid++
	entity := &Entity{
		&unit{
			e,
			true,
		},
		eid,
		[]ComponentRoutine{},
	}
	e.entities = append(e.entities, entity)
	return entity
}

// Entity - get entity by id
func (e *Engine) Entity(id uint) *Entity {
	// for _, entity := range e.entities {
	// 	if entity.ID() == id {
	// 		return entity
	// 	}
	// }
	return nil
}

// DeleteEntity - deletes entity by object/id
func (e *Engine) DeleteEntity(entity *Entity) {
	// for _, system := range systems {
	// 	system.DeleteEntity(entity)
	// }
	// REFACTOR NOTES
	// Systems need to have a copy of the entity components so if the
	// entity is delete, the system must also delete the delete entity's
	// components.

	// PRE REFACTOR
	// entity.UnRegisterAllFromSystems(e)

	// SEE IF THIS IS STILL NECESARRY
	// entity.DeleteComponents()

	// THIS BELOW MIGHT STILL BE NECESARY
	// copy(e.entities[idx:], e.entities[idx+1:])
	// e.entities[len(e.entities)-1] = nil
	// e.entities = e.entities[:len(e.entities)-1]
}

// NewSystem - creates new system
func NewSystem(s SystemRoutine) {
	systems = append(systems, s)
	if updater, ok := s.(Updater); ok {
		updaters = append(updaters, updater)
	}
	if listener, ok := s.(Listener); ok {
		listeners = append(listeners, listener)
	}
}

// DeleteSystem - deletes system by object
// Systems - returns all systems
// func (e *Engine) Systems() []SystemRoutine {
// 	return systems
// }

// NewIntegrant - creates new engine component
func NewIntegrant(integrant IntegrantRoutine) {
	// component.Init()
	integrants = append(integrants, integrant)
}

// Integrant - gets component by object
// func (e *Engine) DeleteIntegrant(lookup IntegrantRoutine) IntegrantRoutine {
// 	for _, component := range components {
// 		if fmt.Sprintf("%T", lookup) == fmt.Sprintf("%T", component) {
// 			return component
// 		}
// 	}
// 	log.Fatal("Integrant requested but doens't exist. Make sure all packages are imported")
// 	return nil
// }

// DeleteComponent - deletes component by object

// SetSettings - set settings
func (e *Engine) SetSettings(settings *Settings) {
	e.settings = settings
}

// Settings - get settings
func (e *Engine) Settings() *Settings {
	return e.settings
}

// update - updates each of the updaters systems
func (e *Engine) update() {
	log.Println("Engine Update")
	e.newTime = time.Now()
	frameTime := e.newTime.Sub(e.currentTime).Nanoseconds()
	e.currentTime = e.newTime
	e.accumulator += frameTime

	for e.accumulator >= e.deltaTime {
		for _, updater := range updaters {
			if e.running {
				updater.Update()
			}
		}
	}
	e.accumulator -= e.deltaTime
}

func (e *Engine) Listener(lookup Listener) Listener {
	for _, listener := range listeners {
		if fmt.Sprintf("%T", listener) == fmt.Sprintf("%T", lookup) {
			return listener
		}
	}
	log.Fatalf("%T - Listener requested but doens't exist. Make sure all packages are imported", lookup)
	return nil
}
