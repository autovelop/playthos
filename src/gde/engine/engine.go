package engine

import (
	"log"
)

type Engine struct {
	Entities map[string]*Entity
	Systems  map[int]System
	platform *Platform
}

const (
	SystemRender = iota
	SystemUI
	SystemInputKeyboard
	SystemInputPointer
	SystemInputTouch
	SystemAnimation
	SystemNetwork
	SystemPhysics
	SystemAudio
)

func (e *Engine) Init(platform *Platform) {
	log.Printf("Engine > Init")
	e.Entities = make(map[string]*Entity)
	e.Systems = make(map[int]System)
	e.platform = platform
}

func (e *Engine) GetPlatform() *Platform {
	return e.platform
}

func (e *Engine) Update() {
	// log.Printf("Engine > Update")
	for _, v := range e.Systems {
		v.Update(&e.Entities)
	}
}

func (e *Engine) Shutdown() {
	log.Printf("Engine > Shutdown")
	for _, v := range e.Systems {
		v.Stop()
	}
}

func (e *Engine) GetEntity(id string) *Entity {
	log.Printf("Engine > Entity > Get: %v", id)
	return e.Entities[id]
}

func (e *Engine) GetSystem(sys int) System {
	log.Printf("Engine > System > Get: %v", sys)
	return e.Systems[sys]
}

func (e *Engine) AddSystem(sys int, sysRoutine System) System {
	log.Printf("Engine > System > Add: %v", sys)
	e.Systems[sys] = sysRoutine
	return e.Systems[sys]
}
