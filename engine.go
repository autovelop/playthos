/*
Package engine orchestrates all the platforms, systems, entities, assets, and deploy procedures of an application.

ECS

Playthos uses the ECS pattern to manage how objects are perceived in the virtual space. ECS stands for Entity-component-system.
What this means is that as a developer, you will be working with these in order to build an application and manipulate objects at runtime.
*/
package engine

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Registries
var packages []*Package            // Packages required by other imported packages. Used to deploy to multiple platforms, each using it's subset of packages, through a single execution.
var platforms map[string]*Platform // Platforms targeted by application. Used to deploy and to distinguish each platform's packages.
var systems []SystemRoutine        // Systems
var updaters []Updater             // Updater systems (ticks with engine loop)
var drawer Drawer                  // Drawer system (draws with engine loop or when able to)
var integrants []IntegrantRoutine  // Integrants
var listeners []Listener           // Listener integrant (on demand calls like input, audio, etc.)
var platformer Platformer          // Platformer integrant (loading and retrieving assets)
var assets []string                // Assets

// Build Flags
var play = true    // Run application immediately.
var deploy = false // Auto detect package dependencies and build an executable.

func init() {
	fmt.Println("> Engine: Initializing")
	assets = make([]string, 0)
	packages = make([]*Package, 0)
}

// RegisterPackage adds a Package to engine's package registry.
//
// TODO(F): Create a database of official package names.
func RegisterPackage(pckg *Package) {
	packages = append(packages, pckg)
}

// RegisterPlatform adds a Platform to engine's platform registry.
//
// TODO(F): Create a database of official platform names.
func RegisterPlatform(n string, p *Platform) {
	if platforms == nil {
		platforms = make(map[string]*Platform)
	}
	platforms[n] = p
}

// RegisterAsset adds asset string path engine registry.
func RegisterAsset(p string) {
	assets = append(assets, p)
}

// New initializes an Engine instance that could either deploy (platforms and packages are detected automatically) or run the application with an optional Settings parameter.
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

	game.init()
	return game
}

// NewSystem registers and organises the system into its appropriate registries (Drawer, Updater).
func NewSystem(s SystemRoutine) {
	systems = append(systems, s)

	if d, ok := s.(Drawer); ok {
		drawer = d
	}
	if u, ok := s.(Updater); ok {
		updaters = append(updaters, u)
	}
}

// NewIntegrant registers and organises the integrant into its appropriate registries (Platformer, Listener).
func NewIntegrant(i IntegrantRoutine) {
	integrants = append(integrants, i)

	if p, ok := i.(Platformer); ok {
		platformer = p
	}
	if l, ok := i.(Listener); ok {
		listeners = append(listeners, l)
	}
}

// LoadAsset instructs the current platform to load the asset correctly to be used for the application (binary, blob, etc.).
func LoadAsset(path string) {
	platformer.LoadAsset(path)
}

// Engine ties the ECS pattern together, manages application running state, and stores meta information.
type Engine struct {
	gameName string
	entities []*Entity
	settings *Settings
	running  bool
	once     bool
}

// Start updates engine state and executes the first update call.
func (e *Engine) Start() {
	fmt.Println("> Engine: Running")
	e.running = true
	e.update()
}

// Once executes a single engine update call.
func (e *Engine) Once() {
	fmt.Println("> Engine: Running")
	e.once = true
	e.update()
}

// Stop updates engine state in order to commence gracefully shutdown
func (e *Engine) Stop() {
	fmt.Println("> Engine: Stopping")
	e.running = false
}

// Stop gracefully stops all systems and integrants from running.
func (e *Engine) stop() {
	for _, system := range systems {
		system.SetActive(false)
		system.Destroy()
	}
	for _, integrant := range integrants {
		integrant.Destroy()
	}
}

// init detects which systems work with eachother and pairs them up. This always runs before engine is started (not when deploying).
func (e *Engine) init() {
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

var eid uint // Entity IDs
// NewEntity initializes and returns an empty Entity
//
// TODO(F): Generate unique entity ID in "NewEntity()"
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

// Entity returns Entity pointer by given ID
//
// TODO(F): Find better way of searching entities by ID
func (e *Engine) Entity(id uint) *Entity {
	for _, entity := range e.entities {
		if entity.ID() == id {
			return entity
		}
	}
	return nil
}

// Entities returns slice entity pointers
func (e *Engine) Entities() []*Entity {
	return e.entities
}

// Updaters returns slice updater system pointers
func (e *Engine) Updaters() []Updater {
	return updaters
}

// Listeners returns slice updater system pointers
func (e *Engine) Listeners() []Listener {
	return listeners
}

// Platformer returns platformer integrant pointer
func (e *Engine) Platformer() Platformer {
	return platformer
}

// Drawer returns drawer system pointer
func (e *Engine) Drawer() Drawer {
	return drawer
}

// DeleteEntity removes entity from all systems (also empties its components) and the engine's registry
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

// SetSettings overwrites the engines settings
func (e *Engine) SetSettings(settings *Settings) {
	e.settings = settings
}

// Settings returns the engines settings
func (e *Engine) Settings() *Settings {
	return e.settings
}

// Listener returns a Listener based on the given Listener type
func (e *Engine) Listener(lookup Listener) Listener {
	for _, listener := range listeners {
		if fmt.Sprintf("%T", listener) == fmt.Sprintf("%T", lookup) {
			return listener
		}
	}
	log.Fatalf("%T - Listener requested but doens't exist. Make sure all packages are imported", lookup)
	return nil
}

// Integrant returns an Integrant based on the given Integrant type
func (e *Engine) Integrant(lookup IntegrantRoutine) IntegrantRoutine {
	for _, i := range integrants {
		if fmt.Sprintf("%T", i) == fmt.Sprintf("%T", lookup) {
			return i
		}
	}
	log.Fatalf("%T - Integrant requested but doens't exist. Make sure all packages are imported", lookup)
	return nil
}

const frameCap float64 = 250

var (
	frames       uint64
	frameCounter time.Duration
	frameTime    = time.Duration(1000/frameCap) * time.Millisecond
	prevTime     = time.Now()
	unproccTime  time.Duration
	render       bool
)

// Game Loop
//
// BUG(F): Game loop currently performing very badly on desktop OSs
func (e *Engine) update() {
	for e.running || e.once {
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

		if e.once {
			e.once = false
		}
	}
	e.stop()
	time.Sleep(time.Second)
}
