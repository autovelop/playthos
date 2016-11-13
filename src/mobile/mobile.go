package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	// "golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
	"log"

	"gde"
	"gde/opengles"
)

// var (

// 	green  float32
// 	touchX float32
// 	touchY float32
// )

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
		// Create game engine
		engine := &gde.Engine{RenderSystem: "OpenGLES"}
		engine.Init()

		render := &opengles.RenderOpenGLES{}
		// render.Add(engine)
		engine.AddSystem(gde.SystemRender, render)

		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					render.Context, _ = e.DrawContext.(gl.Context)
					onStart(engine)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(engine)
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
				onPaint(engine)
				a.Publish()
				a.Send(paint.Event{})
				// 	case touch.Event:
				// 		touchX = e.X
				// 		touchY = e.Y
			}
		}
	})
}

func onStart(engine *gde.Engine) {
	log.Printf("----------------- OpenGL Version: %v ----------------- ", gl.Version())

	render, err := engine.GetSystem(gde.SystemRender).(gde.RenderRoutine)
	if !err {
		log.Println(err)
		return
	}
	render.Init()

	engine.LoadScene(&gde.Scene{})
}

func onStop(engine *gde.Engine) {
	engine.Shutdown()
}

func onPaint(engine *gde.Engine) {
	engine.Update()
}
