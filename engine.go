package engine

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var tags []string

var Game *Engine

func init() {
	log.Println("init engine")
	Game = &Engine{}
	Game.entities = make([]*Entity, 0)
	Game.ready = true

	tags = make([]string, 0)
}

func RegisterPackage(newTags ...string) {
	tags = append(tags, newTags...)
}

func GetTags() string {
	return strings.Join(tags[:], " ")
}

type Engine struct {
	systems       []System
	observerables []Observerable
	entities      []*Entity
	components    []ComponentRoutine
	newTime       time.Time
	currentTime   time.Time
	accumulator   int64
	deltaTime     int64
	frames        uint64
	ready         bool
	running       bool
}

func (e *Engine) Update() {
	log.Println("Engine Update")
	e.newTime = time.Now()
	frameTime := e.newTime.Sub(e.currentTime).Nanoseconds()
	e.currentTime = e.newTime
	e.accumulator += frameTime

	for e.accumulator >= e.deltaTime {
		for _, system := range Game.systems {
			system.Update()
		}
	}
	e.accumulator -= e.deltaTime
}

func Run() {
	if Game.ready == false {
		log.Println("Engine must be ready before it can start")
	} else {
		log.Println("Engine Running")
		Game.running = true
		for Game.running == true {
			Game.Update()
		}
	}
}

func Stop() {
	if Game.running == false {
		log.Println("Engine cannot be stopped if it is not running")
	} else {
		log.Println("Engine Stopped")
		Game.running = false
	}
}

func NewComponent(component ComponentRoutine) {
	component.Prepare()
	Game.components = append(Game.components, component)
}

func NewSystem(system System) {
	system.Prepare()
	for _, component := range Game.components {
		system.LoadComponent(component)
		// component.RegisterToSystem(system)
	}
	Game.systems = append(Game.systems, system)
}

func NewObserverable(observerable Observerable) {
	for _, component := range Game.components {
		component.RegisterToObserverable(observerable)
	}
	Game.observerables = append(Game.observerables, observerable)
}

var idnum uint = 0

func NewEntity() *Entity {
	idnum++
	Game.entities = append(Game.entities, &Entity{ID: idnum})
	// log.Println(Game.entities)
	return Game.entities[len(Game.entities)-1]
}

func GetEntity(id uint) *Entity {
	for _, entity := range Game.entities {
		if entity.ID == id {
			return entity
		}
	}
	return nil
}

func DeleteEntity(ent *Entity) {
	for e, entity := range Game.entities {
		if entity == ent {
			// for _, component := range entity.Components() {
			// 	log.Printf("-- %T\n", component)
			// }

			entity.UnRegisterAllFromSystems(Game)
			entity.DeleteComponents()

			// log.Printf("-- %v\n", "test")
			// for _, component := range entity.Components() {
			// 	log.Printf("-- %T\n", component)
			// }

			copy(Game.entities[e:], Game.entities[e+1:])
			Game.entities[len(Game.entities)-1] = nil
			Game.entities = Game.entities[:len(Game.entities)-1]
			break
		}
	}
}

func GetObserverable(lookup Observerable) Observerable {
	for _, observerable := range Game.observerables {
		if fmt.Sprintf("%T", observerable) == fmt.Sprintf("%T", observerable) {
			return observerable
		}
	}
	log.Fatal("System requested but doens't exist. Make sure all packages are imported")
	return nil
	// return nil, errors.New("Trying to get Observerable that does not exist")
}
