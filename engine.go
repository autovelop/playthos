package engine

import (
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
	return strings.Join(tags[:], ",")
}

// TODO do the below later so we can do validation for the tags in terms of platform/os etc.
// type TagList []string
// var platformTags []string
// var osTags []string
// var packageTags []string
// func init() {
// 	platformTags = make([]string, 0)
// 	osTags = make([]string, 0)
// 	packageTags = make([]string, 0)
// }
// func NewPackage(pckgPlatformTags *TagList, pckgOSTags *TagList, pckgTags *TagList) {
// 	// do validation so mobile platforms won't try build for windows/linux/macos and vice versa
// 	if len(*pckgPlatformTags) < 1 {
// 		log.Println("Package is missing platform tag")
// 		return
// 	} else if len(*pckgOSTags) < 1 {
// 		log.Println("Package is missing os tag")
// 		return
// 	} else if len(*pckgTags) < 1 {
// 		log.Println("Package is missing package tag")
// 		return
// 	}
// 	platformTags = append(platformTags, strings.Join(pckgPlatformTags[:], ","))
// 	osTags = append(osTags, pckgOSTags)
// 	packageTags = append(packageTags, pckgTags)
// }
// func GetTags() string {
// 	return fmt.Sprintf("%v,%v,%v", strings.Join(platformTags[:], ","), strings.Join(osTags[:], ","), strings.Join(packageTags[:], ","))
// }

type Engine struct {
	systems     []System
	entities    []*Entity
	newTime     time.Time
	currentTime time.Time
	accumulator int64
	deltaTime   int64
	frames      uint64
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

func (e *Engine) Prepare() {
	log.Println("Engine Prepare")
	e.systems = make([]System, 0)
	e.entities = make([]*Entity, 0)
	// for _, system := range e.systems {
	// 	system.Prepare()
	// }
}

func (e *Engine) NewSystem(system System) {
	system.Prepare()
	e.systems = append(e.systems, system)
}

func (e *Engine) NewEntity(entity *Entity) {
	e.entities = append(e.entities, entity)
}

func (e *Engine) GetBuildTags() string {
	return "opengl"
}

// import (
// 	"log"
// 	"time"
// )

// const (
// 	SystemRender = iota
// 	SystemUI
// 	SystemInputKeyboard
// 	SystemInputPointer
// 	SystemInputTouch
// 	SystemAnimation
// 	SystemNetwork
// 	SystemPhysics
// 	SystemAudio
// )

// type Engine struct {
// 	Entities    map[string]*Entity
// 	Systems     map[int]System
// 	debug       bool
// 	newTime     time.Time
// 	currentTime time.Time
// 	accumulator int64
// 	deltaTime   int64
// 	frames      uint64
// 	// framesCounter time.Duration
// 	// frameTime     time.Duration
// 	// startTime     time.Time
// 	// lastTime      time.Time
// 	// unproccTime   time.Duration
// }

// func (e *Engine) Init() {
// 	log.Printf("Engine > Init")

// 	e.Entities = make(map[string]*Entity)
// 	e.Systems = make(map[int]System)

// 	// e.frameTime = time.Duration(1000/60) * time.Millisecond
// 	e.currentTime = time.Now()
// 	// e.startTime = time.Now()
// }

// func (e *Engine) Update() {
// 	// e.newTime = time.Now()
// 	// frameTime := e.newTime.Sub(e.currentTime).Nanoseconds()
// 	// e.currentTime = e.newTime
// 	// e.accumulator += frameTime

// 	// for e.accumulator >= e.deltaTime {
// 	for _, v := range e.Systems {
// 		v.Update(&e.Entities)
// 	}
// 	// }
// 	// e.accumulator -= e.deltaTime

// 	// t += e.deltaTime

// 	// e.lastTime = e.startTime

// 	// e.unproccTime += passedTime
// 	// e.framesCounter += passedTime

// 	// for e.unproccTime > e.frameTime {
// 	// 	// log.Printf("Engine > Update")
// 	// 	e.unproccTime -= e.frameTime
// 	// 	e.frames++

// 	// 	if e.framesCounter >= time.Second {
// 	// 		log.Printf("%d FPS\n", e.frames)
// 	// 		e.frames = 0
// 	// 		e.framesCounter -= time.Second
// 	// 	}
// 	// }
// }

// func (e *Engine) Shutdown() {
// 	log.Printf("Engine > Shutdown")
// 	for _, v := range e.Systems {
// 		v.Stop()
// 	}
// }

// func (e *Engine) GetEntity(id string) *Entity {
// 	log.Printf("Engine > Entity > Get: %v", id)
// 	return e.Entities[id]
// }

// func (e *Engine) DeleteEntity(id string) {
// 	log.Printf("Engine > Entity > Delete: %v", id)
// 	delete(e.Entities, id)
// }

// func (e *Engine) GetSystem(sys int) System {
// 	log.Printf("Engine > System > Get: %v", sys)
// 	return e.Systems[sys]
// }

// func (e *Engine) AddSystem(sys int, sysRoutine System) System {
// 	log.Printf("Engine > System > Add: %v", sys)
// 	e.Systems[sys] = sysRoutine
// 	return e.Systems[sys]
// }
