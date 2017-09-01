package engine

import (
	"fmt"
	// "github.com/jteeuwen/go-bindata"
	// "io/ioutil"
	"log"
	// "os"
	// "os/exec"
	// "path/filepath"
	// "runtime"
	// "strings"
	"time"
)

const (
	PlatformLinux   = "lin"
	PlatformWindows = "win"
	PlatformMacOS   = "mac"
	PlatformAndroid = "and"
	PlatformIOS     = "ios"
)

var packages []*Package

// var dependencies []string
// var resolvers []string
var platforms map[string]*Platform
var systems []SystemRoutine
var updaters []Updater
var drawer Drawer
var platformer Platformer
var listeners []Listener
var integrants []IntegrantRoutine

var play bool = true
var deploy bool = false

func init() {
	fmt.Println("> Engine: Initializing")
}

func RegisterPlatform(n string, p *Platform) {
	if platforms == nil {
		platforms = make(map[string]*Platform)
	}
	platforms[n] = p
}

// func RegisterDependency(deps ...string) {
// 	dependencies = append(dependencies, deps...)
// }

// func RegisterResolver(res ...string) {
// 	resolvers = append(resolvers, res...)
// }

func RegisterPackage(pckg *Package) {
	packages = append(packages, pckg)
}

// func GetTags() string {
// 	packages = removeDuplicates(packages)
// 	return strings.Join(packages[:], " ")
// }

type Engine struct {
	gameName    string
	gamePackage string

	entities []*Entity
	updaters []Updater
	settings *Settings
	running  bool
}

func New(n string, p string, s ...*Settings) *Engine {
	// fmt.Printf("%v\n%v\n", platforms, packages)
	if deploy {
		initDeploy(n, p)
		return nil
	}
	// Objectives:
	// Every time you run, it will always include all the systems. The inits for registering the packages, and the mains to make sure the code compiles
	// IF deploying
	// - check registered OSs
	// ELSE (and also if deployed)
	// - get local OS
	// - validate systems required with OS
	// - run
	// log.Println("New()")
	game := &Engine{}
	game.gamePackage = p
	game.gameName = n
	// no tags, just to rerun go run with the system tags in
	// if !play && !deploy && !deployed {
	// 	play = true
	// 	game.Run()
	// 	return game
	// }
	// var osDetect string
	// switch runtime.GOOS {
	// case "windows":
	// 	osDetect = PlatformWindows
	// 	break
	// case "linux":
	// 	osDetect = PlatformLinux
	// 	break
	// case "darwin":
	// 	osDetect = PlatformMacOS
	// 	break
	// }
	if len(s) > 0 {
		settings := s[0]
		// if len(settings.Platforms) <= 0 {
		// 	settings.Platforms = []string{osDetect}
		// }
		game.SetSettings(settings)
	} else {
		game.SetSettings(&Settings{false, 800, 600, true})
	}

	// deploy tag, just to rerun go run again with system and os tags
	// if deploy {
	// 	game.Deploy(game.settings.Platforms...)
	// 	os.Exit(0)
	// }

	// deployed or play tag, just runs game
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

// func (e *Engine) Run() {
// 	// cmd := exec.Command("pwd")
// 	log.Println("Run()")
// 	cmdArgs := []string{
// 		"-c",
// 		fmt.Sprintf("go run -tags=\"play %v\" *.go", GetTags()),
// 	}
// 	log.Println(fmt.Sprintf("go run -tags=\"play %v\" *.go", GetTags()))
// 	cmd := exec.Command("/bin/sh", cmdArgs...)
// 	cmdErr, _ := cmd.StderrPipe()
// 	startErr := cmd.Start()
// 	if startErr != nil {
// 		return
// 	}
// 	errOutput, _ := ioutil.ReadAll(cmdErr)
// 	fmt.Printf("%s", errOutput)
// }

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
				// fmt.Printf("%d FPS\n", frames)
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
