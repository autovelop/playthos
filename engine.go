package engine

import (
	"fmt"
	"github.com/jteeuwen/go-bindata"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	gameName    string
	gamePackage string

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

func New(n string, p string, s *Settings) *Engine {
	game := &Engine{}
	game.SetSettings(s)
	game.Init()
	game.gameName = n
	game.gamePackage = p
	return game
}

func (e *Engine) Start() {
	e.running = true
	e.update()
}

func (e *Engine) Stop() {
	for _, system := range systems {
		system.SetActive(false)
	}
	for _, integrant := range integrants {
		integrant.DeleteIntegrant()
	}
	e.running = false
}

func (e *Engine) Init() {
	if play {
		for _, integrant := range integrants {
			integrant.initUnit(e)
			integrant.InitIntegrant()
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
	fmt.Printf("Deploying...\n")

	c := bindata.NewConfig()
	c.Input = []bindata.InputConfig{bindata.InputConfig{
		Path:      filepath.Clean("assets"),
		Recursive: true,
	}}
	c.Package = "engine"
	c.Tags = "deploy play"
	c.Output = fmt.Sprintf("%v/assets.go", "github.com/autovelop/playthos")
	// c.Output = fmt.Sprintf("%v/assets.go", e.gamePackage)
	bindata.Translate(c)

	for _, platform := range platforms {
		simpleName := "linux"
		fileExtension := ""
		cgo := false
		var cc string
		arch386 := false
		switch platform {
		case PlatformLinux:
			fmt.Printf("- Linux\n- Requirements: libgl1-mesa-dev, xorg-dev\n\n")
			break
		case PlatformMacOS:
			fmt.Printf("- MacOS\n- Requirements: xcode 7.3, cmake, libxml2, fuse, osxcross\n- Full details: https://github.com/tpoechtrager/osxcross#packaging-the-sdk\n\n")
			simpleName = "darwin"
			cgo = true
			// cc = "CC=i386-apple-darwin15-g++"
			cc = "CC=o32-clang"
		case PlatformWindows:
			fmt.Printf("- Windows (32-bit only)\n- Requirements: mingw-w64-gcc\n\n")
			simpleName = "windows"
			cgo = true
			arch386 = true
			fileExtension = ".exe"
			cc = "CC=i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp"
			break
		default:
			continue
			break
		}
		// log.Fatalf("%v/bin/%v_%v", e.gamePackage, strings.ToLower(e.gameName), simpleName)
		cmdArgs := []string{
			"build",
			"-v",
			"-o",
			fmt.Sprintf("%v/bin/%v_%v%v", e.gamePackage, strings.ToLower(e.gameName), simpleName, fileExtension),
			"-tags",
			fmt.Sprintf("deploy play %v %v", simpleName, GetTags()),
			e.gamePackage,
		}
		cmd := exec.Command("go", cmdArgs...)
		cmd.Env = os.Environ()

		if cgo {
			cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
		}

		if arch386 {
			cmd.Env = append(cmd.Env, "GOARCH=386")
		} else {
			cmd.Env = append(cmd.Env, "GOARCH=amd64")
		}
		if len(cc) > 0 {
			cmd.Env = append(cmd.Env, cc)
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", simpleName))

		// cmdOut, _ := cmd.StdoutPipe()
		cmdErr, _ := cmd.StderrPipe()

		startErr := cmd.Start()
		if startErr != nil {
			log.Println("here")
			log.Println(startErr)
			// cmd.Wait()
			return
		}

		// read stdout and stderr
		// stdOutput, _ := ioutil.ReadAll(cmdOut)
		errOutput, _ := ioutil.ReadAll(cmdErr)

		// fmt.Printf("STDOUT: %s\n", stdOutput)
		fmt.Printf("%s", errOutput)

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
			updater.Update()
		}
		if !e.running {
			os.Exit(0)
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
