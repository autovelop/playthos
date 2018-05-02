package engine

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Platform consts to use for packages that are targeted at specific platforms
const (
	PlatformLinux   = "lin"
	PlatformWindows = "win"
	PlatformMacOS   = "mac"
	PlatformAndroid = "and"
	PlatformIOS     = "ios"
)

var packages []*Package

var platforms map[string]*Platform
var systems []SystemRoutine
var updaters []Updater
var drawer Drawer
var platformer Platformer
var listeners []Listener
var integrants []IntegrantRoutine
var assetRegistry []string

var play = true
var deploy = false

func init() {
	fmt.Println("> Engine: Initializing")
	assetRegistry = make([]string, 0)
}

func RegisterPlatform(n string, p *Platform) {
	if platforms == nil {
		platforms = make(map[string]*Platform)
	}
	platforms[n] = p
}

func RegisterPackage(pckg *Package) {
	packages = append(packages, pckg)
}

type Engine struct {
	gameName    string
	gamePackage string

	entities []*Entity
	updaters []Updater
	settings *Settings
	running  bool
}

func RegisterAsset(p string) {
	assetRegistry = append(assetRegistry, p)
}

func New(n string, s ...*Settings) *Engine {
	if deploy {
		initDeploy(n, os.Args[1])
		return nil
	}
	game := &Engine{}
	game.gameName = n

	if len(s) > 0 {
		settings := s[0]
		game.SetSettings(settings)
	} else {
		game.SetSettings(&Settings{false, 800, 600, true})
	}

	game.Init()
	return game
}

func (e *Engine) Start() {
	fmt.Println("> Engine: Enjoy!")
	e.running = true
	e.update()
}

func (e *Engine) Stop() {
	for _, system := range systems {
		system.SetActive(false)
		system.Destroy()
	}
	for _, integrant := range integrants {
		integrant.Destroy()
	}
	e.running = false
}

func (e *Engine) Init() {
	if play {
		for _, integrant := range integrants {
			integrant.initUnit(e)
			integrant.InitIntegrant()
			integrant.SetActive(true)
			for _, i := range integrants {
				i.AddIntegrant(integrant)
			}
		}
		for _, system := range systems {
			for _, integrant := range integrants {
				system.AddIntegrant(integrant)
			}
			system.initUnit(e)
			system.InitSystem()
			system.SetActive(true)
		}
	}
}

var eid uint = 0

func (e *Engine) NewEntity() *Entity {
	eid++
	entity := &Entity{
		&unit{
			e,
			true,
			0,
		},
		eid,
		[]ComponentRoutine{},
	}
	e.entities = append(e.entities, entity)
	return entity
}

func (e *Engine) Entity(id uint) *Entity {
	// for _, entity := range e.entities {
	// 	if entity.ID() == id {
	// 		return entity
	// 	}
	// }
	return nil
}

func (e *Engine) DeleteEntity(entity *Entity) {
	for _, system := range systems {
		system.DeleteEntity(entity)
	}
	for i := 0; i < len(e.entities); i++ {
		ent := e.entities[i]
		if ent.ID() == entity.ID() {
			copy(e.entities[i:], e.entities[i+1:])
			e.entities[len(e.entities)-1] = nil
			e.entities = e.entities[:len(e.entities)-1]
		}
	}
}

func NewSystem(s SystemRoutine) {
	systems = append(systems, s)

	if d, ok := s.(Drawer); ok {
		drawer = d
	}
	if u, ok := s.(Updater); ok {
		updaters = append(updaters, u)
	}
}

func NewIntegrant(i IntegrantRoutine) {
	integrants = append(integrants, i)

	if p, ok := i.(Platformer); ok {
		platformer = p
	}
	if l, ok := i.(Listener); ok {
		listeners = append(listeners, l)
	}
}

func (e *Engine) SetSettings(settings *Settings) {
	e.settings = settings
}

func (e *Engine) Settings() *Settings {
	return e.settings
}

func LoadAsset(path string) {
	platformer.LoadAsset(path)
}

func (e *Engine) update() {
	const frameCap float64 = 250
	var (
		frames       uint64
		frameCounter time.Duration
		frameTime    time.Duration = time.Duration(1000/frameCap) * time.Millisecond
		prevTime     time.Time     = time.Now()
		unproccTime  time.Duration
		render       bool
	)

	for e.running {
		startTime := time.Now()
		elapsed := startTime.Sub(prevTime)
		prevTime = startTime

		unproccTime += elapsed
		frameCounter += elapsed

		for unproccTime > frameTime {
			unproccTime -= frameTime

			for _, updater := range updaters {
				updater.Update()
			}

			if frameCounter >= time.Second {
				frames = 0
				frameCounter -= time.Second
			}
			render = true
		}

		if render && drawer != nil {
			drawer.Draw()
			frames++
		} else {
			time.Sleep(time.Millisecond)
		}
	}
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

func (e *Engine) Integrant(lookup IntegrantRoutine) IntegrantRoutine {
	for _, i := range integrants {
		if fmt.Sprintf("%T", i) == fmt.Sprintf("%T", lookup) {
			return i
		}
	}
	log.Fatalf("%T - Integrant requested but doens't exist. Make sure all packages are imported", lookup)
	return nil
}
