package main

import (
	// "encoding/binary"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	// "golang.org/x/mobile/event/touch"
	// "golang.org/x/mobile/exp/f32"
	// "golang.org/x/mobile/exp/gl/glutil"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/mobile/gl"
	"log"

	"gde"
	"gde/components"
	"gde/geometry"
	"gde/systems"
	"gde/systems/opengles"
)

// var (

// 	green  float32
// 	touchX float32
// 	touchY float32
// )

func main() {
	app.Main(func(a app.App) {
		// Create game engine
		engine := &gde.Engine{}
		engine.Init()

		render := &opengles.RenderOpenGLES{}
		render.Add(engine)

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

	render, err := engine.GetSystem(&systems.Render{}).(systems.RenderRoutine)
	if !err {
		log.Println(err)
		return
	}
	render.Init()

	// Simple Quad mesh renderer
	renderer := &components.Renderer{}
	renderer.Init()

	renderer.LoadMesh(&geometry.Mesh{
		Vertices: []float32{
			0.1, 0.1, 0.0,
			0.1, -0.1, 0.0,
			-0.1, -0.1, 0.0,
			-0.1, 0.1, 0.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	render.LoadRenderer(renderer)

	// Create player entity
	player := &gde.Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	transform := &components.Transform{}
	transform.Init()
	transform.SetProperty("Position", mgl32.Vec3{0.2, -0.5, 0})
	transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 45})

	player.AddComponent(transform)
	player.AddComponent(renderer)

	box := &gde.Entity{Id: "Box"}
	box.Init()
	box.Add(engine)

	box_transform := &components.Transform{}
	box_transform.Init()
	box_transform.SetProperty("Position", mgl32.Vec3{0.2, 0.5, 0})
	box_transform.SetProperty("Rotation", mgl32.Vec3{0, 0, 0})
	box.AddComponent(box_transform)

	box.AddComponent(renderer)

	// Samsung S6
	// - OpenGL ES Version: GL_ES_3_0 (possibly GL_ES_3_1 soon)
	// - GLSL Version: #version 100 / #version 300 es / #version 310 es
	// Samsung S3 Mini
	// - OpenGL ES Version: GL_ES_2_0
	// - GLSL Version: #version 100
	// http://www.shaderific.com/blog/2014/3/13/tutorial-how-to-update-a-shader-for-opengl-es-30
	// http://stackoverflow.com/questions/29888213/solved-qopenglshader-cant-compile-glsl-120-on-android
	// https://github.com/mattdesl/lwjgl-basics/wiki/GLSL-Versions
}

func onStop(engine *gde.Engine) {
  engine.
	// render.Context.DeleteProgram(program)
	// render.Context.DeleteBuffer(buf)
}

func onPaint(engine *gde.Engine) {
	engine.Update()
}
