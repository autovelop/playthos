package main

import (
	"gde/editor"
	"gde/engine"
	"gde/render/opengl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

// ECS 2

// type Engine struct {
// 	systems  []System
// 	entities []*Entity
// }

// func (e *Engine) Update() {
// 	log.Println("Engine Update")
// 	for _, system := range e.systems {
// 		system.Update()
// 	}
// }

// func (e *Engine) Prepare() {
// 	log.Println("Engine Prepare")
// 	for _, system := range e.systems {
// 		system.Prepare()
// 	}
// }

// type Entity struct {
// 	components []Component
// }

// func (e *Entity) RegisterToSystems(engine *Engine) {
// 	for _, system := range engine.systems {
// 		for _, component_type := range system.ComponentTypes() {
// 			for _, component := range e.components {
// 				if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", component_type) {
// 					component.RegisterToSystem(system)
// 				}
// 			}
// 		}
// 	}
// }

// type System interface {
// 	Update()
// 	Prepare()
// 	ComponentTypes() []Component
// }

// type Render struct {
// 	transforms []*Transform
// }

// func (r *Render) Prepare() {
// 	log.Println("Render Prepare")
// }

// func (r *Render) ComponentTypes() []Component {
// 	return []Component{&Transform{}}
// }
// func (r *Render) Update() {
// 	log.Println("Render Update")
// 	for _, transform := range r.transforms {
// 		log.Printf("Transform Update: %v\n", transform)
// 	}
// }

// func (r *Render) RegisterTransform(transform *Transform) {
// 	r.transforms = append(r.transforms, transform)
// }

// type Component interface {
// 	RegisterToSystem(System)
// }

// type Transform struct {
// 	position string
// 	rotation string
// 	scale    string
// }

// func (t *Transform) RegisterToSystem(system System) {
// 	log.Println("Registering Transform")
// 	switch system := system.(type) {
// 	case *Render:
// 		system.RegisterTransform(t)
// 	}
// }

func init() {
	runtime.LockOSThread()
}

func main() {
	game := &engine.Engine{}

	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw")
	}
	defer glfw.Terminate()

	game.Prepare()

	// SYSTEM CREATION
	game.NewSystem(&opengl.OpenGL{})

	scene := &editor.Scene{}
	scene.LoadScene(game)

	for true {
		game.Update()
	}
	// engine.Update()
	// defer profile.Start(profile.ProfilePath(os.Getenv("HOME"))).Stop()

	// // Greate game engine
	// game := &engine.Engine{} // Set Device, OS, and OpenGL
	// game.Init()

	// // Create render system
	// render := &opengl.OpenGL{}
	// game.AddSystem(engine.SystemRender, render)
	// render.Init()
	// window := render.GetWindow()

	// // Create keyboard input system
	// // window.GetUserPointer().
	// keyInput := &keyboard.KeyListener{Window: window}
	// game.AddSystem(engine.SystemInputKeyboard, keyInput)
	// keyInput.Init()

	// // Escape
	// keyInput.BindOn(256, func() {
	// 	keyInput.Window.SetShouldClose(true)
	// })

	// // Create pointer input system
	// mouseInput := &mouse.MoveListener{Window: window}
	// game.AddSystem(engine.SystemInputPointer, mouseInput)
	// mouseInput.Init()

	// scene := &editor.Scene{}
	// scene.LoadEngine(game)
	// scene.LoadScene()

	// for true {
	// 	game.Update()
	// }
}

// OLD STUFF

// import (
// "log"
// "runtime"
// "github.com/go-gl/glfw/v3.2/glfw"
// "github.com/pkg/profile"
// "gde/editor"
// "gde/engine"
// "gde/input/keyboard"
// "gde/input/mouse"
// "gde/render/opengl"
// )

// func init() {
// 	runtime.LockOSThread()
// }

// func main() {
// defer profile.Start(profile.ProfilePath(os.Getenv("HOME"))).Stop()

// Intialize GLFW
// if err := glfw.Init(); err != nil {
// 	log.Fatalln("failed to initialize glfw:", err)
// }
// defer glfw.Terminate()

// // Greate game engine
// game := &engine.Engine{} // Set Device, OS, and OpenGL
// game.Init()

// // Create render system
// render := &opengl.OpenGL{}
// game.AddSystem(engine.SystemRender, render)
// render.Init()
// window := render.GetWindow()

// // Create keyboard input system
// // window.GetUserPointer().
// keyInput := &keyboard.KeyListener{Window: window}
// game.AddSystem(engine.SystemInputKeyboard, keyInput)
// keyInput.Init()

// // Escape
// keyInput.BindOn(256, func() {
// 	keyInput.Window.SetShouldClose(true)
// })

// // Create pointer input system
// mouseInput := &mouse.MoveListener{Window: window}
// game.AddSystem(engine.SystemInputPointer, mouseInput)
// mouseInput.Init()

// scene := &editor.Scene{}
// scene.LoadEngine(game)
// scene.LoadScene()

// for true {
// 	game.Update()
// }
// }
