package engine

import (
	"fmt"
	"github.com/jteeuwen/go-bindata"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	PlatformLinux   = "lin"
	PlatformWindows = "win"
	PlatformMacOS   = "mac"
	PlatformAndroid = "and"
	PlatformIOS     = "ios"
)

var packages []string
var systems []SystemRoutine
var updaters []Updater
var drawer Drawer
var listeners []Listener
var integrants []IntegrantRoutine

var play bool = true
var deploy bool = false

func init() {
	fmt.Println("> Engine: Initializing")
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
}

func New(n string, p string, s ...*Settings) *Engine {
	game := &Engine{}
	game.gamePackage = p
	if !deploy && !play {
		play = false
		game.Run()
		return game
	}
	var osDetect string
	switch runtime.GOOS {
	case "windows":
		osDetect = PlatformWindows
		break
	case "linux":
		osDetect = PlatformLinux
		break
	case "darwin":
		osDetect = PlatformMacOS
		break
	}
	if len(s) > 0 {
		settings := s[0]
		if len(settings.Platforms) <= 0 {
			settings.Platforms = []string{osDetect}
		}
		game.SetSettings(settings)
	} else {
		game.SetSettings(&Settings{false, 800, 600, true, []string{osDetect}})
	}
	game.Init()
	game.gameName = n

	if deploy {
		game.Deploy(game.settings.Platforms...)
		os.Exit(0)
	}

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
	if l, ok := s.(Listener); ok {
		listeners = append(listeners, l)
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

func (e *Engine) Run() {
	// cmd := exec.Command("pwd")
	cmdArgs := []string{
		"-c",
		fmt.Sprintf("go run -tags=\"play %v\" *.go", GetTags()),
	}
	cmd := exec.Command("/bin/sh", cmdArgs...)
	cmdErr, _ := cmd.StdoutPipe()
	startErr := cmd.Start()
	if startErr != nil {
		return
	}
	errOutput, _ := ioutil.ReadAll(cmdErr)
	fmt.Printf("%s", errOutput)
}

// TODO: This function is still a mess. Need to make Deploy system in order to tidy it up.
func (e *Engine) Deploy(platforms ...string) {
	fmt.Printf("Deploying...\n")

	c := bindata.NewConfig()
	c.Input = []bindata.InputConfig{bindata.InputConfig{
		Path:      filepath.Clean("assets"),
		Recursive: true,
	}}
	c.Package = "engine"
	c.Tags = "deployed"
	c.Output = fmt.Sprintf("%v/assets.go", "../playthos")
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

		cmdArgs := []string{
			"build",
			"-v",
			"-o",
			fmt.Sprintf("bin/%v_%v%v", strings.ToLower(e.gameName), simpleName, fileExtension),
			"-tags",
			fmt.Sprintf("deployed %v %v", simpleName, GetTags()),
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

		cmdErr, _ := cmd.StderrPipe()

		startErr := cmd.Start()
		if startErr != nil {
			return
		}
		errOutput, _ := ioutil.ReadAll(cmdErr)
		fmt.Printf("%s", errOutput)
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
