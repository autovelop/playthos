package engine

import (
	"log"
	"time"
)

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

type Engine struct {
	Entities      map[string]*Entity
	Systems       map[int]System
	debug         bool
	frames        uint64
	framesCounter time.Duration
	frameTime     time.Duration
	startTime     time.Time
	lastTime      time.Time
	unproccTime   time.Duration
}

func (e *Engine) Printf(format string, data ...interface{}) {
	log.Printf("Engine > Init", data)
}

func (e *Engine) Init() {
	if e.debug {
		e.Printf("Engine > Init")
	}
	e.Entities = make(map[string]*Entity)
	e.Systems = make(map[int]System)

	e.frameTime = time.Duration(1000/60) * time.Millisecond
	e.lastTime = time.Now()
	e.startTime = time.Now()
	// check out wolfengo
}

func (e *Engine) Update() {
	e.startTime = time.Now()
	passedTime := e.startTime.Sub(e.lastTime)
	e.lastTime = e.startTime

	e.unproccTime += passedTime
	e.framesCounter += passedTime

	for e.unproccTime > e.frameTime {
		// log.Printf("Engine > Update")
		e.unproccTime -= e.frameTime
		e.frames++

		if e.framesCounter >= time.Second {
			log.Printf("%d FPS\n", e.frames)
			e.frames = 0
			e.framesCounter -= time.Second
		}
		for _, v := range e.Systems {
			v.Update(&e.Entities)
		}
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
