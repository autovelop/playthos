package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	// "github.com/pkg/profile"

	"gde/editor"
	"gde/engine"
	"gde/input/keyboard"
	"gde/input/mouse"
	"gde/render/opengl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// defer profile.Start(profile.ProfilePath(os.Getenv("HOME"))).Stop()

	// Intialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Determine platform automatically. Screen Dimensions, Screen Resolution, Screen Aspect Ratio, etc.
	platform := &engine.Platform{}
	platform.NewPlatform(360, 640, 360, 640)

	// Greate game engine
	game := &engine.Engine{} // Set Device, OS, and OpenGL
	game.Init(platform)

	// Create render system
	render := &opengl.OpenGL{}
	game.AddSystem(engine.SystemRender, render)
	render.Init()
	window := render.GetWindow()

	// Create keyboard input system
	// window.GetUserPointer().
	keyInput := &keyboard.KeyListener{Window: window}
	game.AddSystem(engine.SystemInputKeyboard, keyInput)
	keyInput.Init()

	// Escape
	keyInput.BindOn(256, func() {
		keyInput.Window.SetShouldClose(true)
	})

	// Create pointer input system
	mouseInput := &mouse.MoveListener{Window: window}
	game.AddSystem(engine.SystemInputPointer, mouseInput)
	mouseInput.Init()

	scene := &editor.Scene{}
	scene.LoadScene(game)

	for true {
		game.Update()
	}
}
