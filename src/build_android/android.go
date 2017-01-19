package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
	"log"

	"gde/editor"
	"gde/engine"
	// "gde/input/touch"
	"gde/render"
	"gde/render/opengles"
)

// Samsung S6
// - OpenGL ES Version: GL_ES_3_0 (possibly GL_ES_3_1 soon)
// - GLSL Version: #version 100 / #version 300 es / #version 310 es
// Samsung S3 Mini
// - OpenGL ES Version: GL_ES_2_0
// - GLSL Version: #version 100
// http://www.shaderific.com/blog/2014/3/13/tutorial-how-to-update-a-shader-for-opengl-es-30
// http://stackoverflow.com/questions/29888213/solved-qopenglshader-cant-compile-glsl-120-on-android
// https://github.com/mattdesl/lwjgl-basics/wiki/GLSL-Versions

func main() {
	app.Main(func(a app.App) {

		// Determine platform automatically. Screen Dimensions, Screen Resolution, Screen Aspect Ratio, etc.
		platform := &engine.Platform{}
		platform.NewPlatform(360, 640, 360, 640)

		// Greate game engine
		game := &engine.Engine{} // Set Device, OS, and OpenGL
		game.Init(platform)

		render := &opengles.OpenGLES{}
		game.AddSystem(engine.SystemRender, render)

		// render_ui := &uigles.UIGLES{}
		// game.AddSystem(engine.SystemUI, render_ui)

		// touchSys := &touchPkg.Touch{}
		// engine.AddSystem(gde.SystemInputTouch, touchSys)

		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					render.Context, _ = e.DrawContext.(gl.Context)
					onStart(game)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(game)
					render.Context = nil
				}
			case size.Event:
				render.Size = e
				// 		touchX = float32(engine.sz.WidthPx / 2)
				// 		touchY = float32(engine.sz.HeightPx / 2)
			case paint.Event:
				if render.Context == nil || e.External {
					continue
				}
				onPaint(game)
				a.Publish()
				a.Send(paint.Event{})
			case touch.Event:
				// touchSys.Touch(0, e.X, e.Y)
				// touchX = e.X
				// touchY = e.Y
			}
		}
	})
}

func onStart(game *engine.Engine) {
	sys_render, err := game.GetSystem(engine.SystemRender).(render.RenderRoutine)
	if !err {
		log.Printf("android.go - %v", err)
		return
	}
	sys_render.Init()

	scene := &editor.Scene{}
	scene.LoadScene(game)
}

func onStop(game *engine.Engine) {
	game.Shutdown()
}

func onPaint(game *engine.Engine) {
	game.Update()
}
