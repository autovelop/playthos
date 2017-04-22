package engine

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var tags []string

func init() {
	tags = make([]string, 0)
}
func RegisterPackage(newTags ...string) {
	tags = append(tags, newTags...)
}
func GetTags() string {
	return strings.Join(tags[:], " ")
}

var systems []System
var subjects []Subject
var components []Component

type Engine struct {
	systems     []System
	subjects    []Subject
	entities    []*Entity
	components  []Component
	newTime     time.Time
	currentTime time.Time
	accumulator int64
	deltaTime   int64
	frames      uint64
	prepared    bool
	running     bool
}

func NewComponent(component Component) {
	log.Printf("New Component Added: %T\n", component)
	components = append(components, component)
}

func (e *Engine) Update() {
	log.Println("Engine Update")
	e.newTime = time.Now()
	frameTime := e.newTime.Sub(e.currentTime).Nanoseconds()
	e.currentTime = e.newTime
	e.accumulator += frameTime

	for e.accumulator >= e.deltaTime {
		for _, system := range systems {
			system.Update()
		}
	}
	e.accumulator -= e.deltaTime
}

func init() {
	log.Println("init engine")
	systems = make([]System, 0)
	subjects = make([]Subject, 0)
	components = make([]Component, 0)
}

func (e *Engine) Run() {
	if e.prepared == false {
		log.Println("Engine must be prepared before it can start")
	} else {
		log.Println("Engine Running")
		e.running = true
		for e.running == true {
			e.Update()
		}
	}
}

func (e *Engine) Stop() {
	if e.running == false {
		log.Println("Engine cannot be stopped if it is not running")
	} else {
		log.Println("Engine Stopped")
		e.running = false
	}
}

func (e *Engine) Prepare() {
	log.Println("Engine Prepare")
	e.entities = make([]*Entity, 0)

	// swap the prelaunch slices into runtime slices
	e.components = components
	e.systems = systems
	e.subjects = subjects

	for _, component := range e.components {
		component.Prepare()
	}
	for _, system := range e.systems {
		system.Prepare()
	}

	for _, system := range e.systems {
		for _, component := range e.components {
			component.RegisterToSystem(system)
		}
	}

	for _, subject := range e.subjects {
		for _, component := range e.components {
			component.RegisterToSubject(subject)
		}
	}

	// for _, subject := range e.subjects {
	// 	for _, component := range e.components {
	// 		component.RegisterToSystem(subject)
	// 	}
	// }
	e.prepared = true
}

// during game launch
func NewSystem(system System) {
	systems = append(systems, system)
}
func NewSubject(subject Subject) {
	subjects = append(subjects, subject)
}

// during runtime
func (e *Engine) NewSystem(system System) {
	system.Prepare()
	e.systems = append(e.systems, system)
}
func (e *Engine) NewSubject(subject Subject) {
	subject.Prepare()
	e.subjects = append(e.subjects, subject)
}
func (e *Engine) NewEntity(entity *Entity) {
	e.entities = append(e.entities, entity)
}

func (e *Engine) GetSubject(lookup Subject) (Subject, error) {
	// e.entities = append(e.entities, entity)
	for _, subject := range e.subjects {
		if fmt.Sprintf("%T", subject) == fmt.Sprintf("%T", subject) {
			return subject, nil
		}
	}
	return nil, errors.New("Trying to get Subject that does not exist")
}
