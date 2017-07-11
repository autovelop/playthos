package engine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	PlatformLinux   = "linux"
	PlatformWindows = "windows"
	PlatformMacOS   = "macos"
	PlatformAndroid = "android"
	PlatformIOS     = "ios"
)

var packages []string
var systems []SystemRoutine
var updaters []Updater
var listeners []Listener
var integrants []IntegrantRoutine

var play bool = false

func init() {
	log.Println("init engine")
}

func RegisterPackage(tags ...string) {
	packages = append(packages, tags...)
}

func GetTags() string {
	packages = removeDuplicates(packages)
	return strings.Join(packages[:], " ")
}

type Engine struct {
	entities       []*Entity
	updaters       []Updater
	settings       *Settings
	running        bool
	runtimePackage string

	newTime     time.Time
	currentTime time.Time
	accumulator int64
	deltaTime   int64
	frames      uint64
}

func New(p string, s *Settings) *Engine {
	game := &Engine{}
	game.SetSettings(s)
	game.Init()
	game.runtimePackage = p
	return game
}

func (e *Engine) Start() {
	e.running = true
	e.update()
}

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
	if updater, ok := s.(Updater); ok {
		updaters = append(updaters, updater)
	}
	if listener, ok := s.(Listener); ok {
		listeners = append(listeners, listener)
	}
}

func NewIntegrant(integrant IntegrantRoutine) {
	integrants = append(integrants, integrant)
}

func (e *Engine) SetSettings(settings *Settings) {
	e.settings = settings
}

func (e *Engine) Settings() *Settings {
	return e.settings
}

func Play() bool {
	return play
}

func (e *Engine) Deploy(platforms ...string) {

	for _, platform := range platforms {
		cmdArgs := []string{
			"install",
			"-tags",
			fmt.Sprintf("play %v", GetTags()),
			"github.com/autovelop/jumper",
		}
		cmd := exec.Command("go", cmdArgs...)
		cmd.Env = os.Environ()
		switch platform {
		case PlatformLinux:
			cmd.Env = append(cmd.Env, "GOOS=linux")
			cmd.Env = append(cmd.Env, "GOARCH=amd64")
			break
		case PlatformMacOS:
			cmd.Env = append(cmd.Env, "GOOS=darwin")
			cmd.Env = append(cmd.Env, "GOARCH=amd64")
		case PlatformWindows:
			cmd.Env = append(cmd.Env, "GOOS=windows")
			cmd.Env = append(cmd.Env, "GOARCH=amd64")
			break
		default:
			continue
			break
		}

		cmdOut, _ := cmd.StdoutPipe()
		cmdErr, _ := cmd.StderrPipe()

		startErr := cmd.Start()
		if startErr != nil {
			log.Println("here")
			log.Println(startErr)
			// cmd.Wait()
			return
		}

		// read stdout and stderr
		stdOutput, _ := ioutil.ReadAll(cmdOut)
		errOutput, _ := ioutil.ReadAll(cmdErr)

		fmt.Printf("STDOUT: %s\n", stdOutput)
		fmt.Printf("ERROUT: %s\n", errOutput)

		// cmd.Wait()
		// log.Fatal(err)
	}
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

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
