package main

import (
	"encoding/binary"
	"log"
	"shared"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	program  gl.Program
	position gl.Attrib
	buf      gl.Buffer

	green  float32
	touchX float32
	touchY float32
)

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(glctx)
					glctx = nil
				}
			case size.Event:
				sz = e
				touchX = float32(sz.WidthPx / 2)
				touchY = float32(sz.HeightPx / 2)
			case paint.Event:
				if glctx == nil || e.External {
					continue
				}
				onPaint(glctx, sz)
				a.Publish()
				a.Send(paint.Event{})
			case touch.Event:
				touchX = e.X
				touchY = e.Y
			}
		}
	})
}

func onStart(glctx gl.Context) {
	log.Printf("----------------- OpenGL Version: %v ----------------- ", gl.Version())
	
	// Samsung S6
	// - OpenGL ES Version: GL_ES_3_0 (possibly GL_ES_3_1 soon)
	// - GLSL Version: #version 100 / #version 300 es / #version 310 es
	// Samsung S3 Mini
	// - OpenGL ES Version: GL_ES_2_0
	// - GLSL Version: #version 100
	// http://www.shaderific.com/blog/2014/3/13/tutorial-how-to-update-a-shader-for-opengl-es-30
	// http://stackoverflow.com/questions/29888213/solved-qopenglshader-cant-compile-glsl-120-on-android
	// https://github.com/mattdesl/lwjgl-basics/wiki/GLSL-Versions
	
	//if gl.Version() == "GL_ES_3_0" {
	//	program, err = glutil.CreateProgram(glctx, shared.VSHADER_OPENGL_ES_3_0, shared.FSHADER_OPENGL_ES_3_0)
	//} else if gl.Version() == "GL_ES_2_0" {
	// For now we will rather only support OPENGL_ES_2_0
	//	program, err = glutil.CreateProgram(glctx, shared.VSHADER_OPENGL_ES_2_0, shared.FSHADER_OPENGL_ES_2_0)
	//}
	
	// For now we will rather only support OPENGL_ES_2_0
	var err error
	program, err = glutil.CreateProgram(glctx, shared.VSHADER_OPENGL_ES_2_0, shared.FSHADER_OPENGL_ES_2_0)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}
	
	buf = glctx.CreateBuffer()
	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	glctx.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

	position = glctx.GetAttribLocation(program, "position")
}

func onStop(glctx gl.Context) {
	glctx.DeleteProgram(program)
	glctx.DeleteBuffer(buf)
}

func onPaint(glctx gl.Context, sz size.Event) {
	glctx.ClearColor(1, 0, 0, 1)
	glctx.Clear(gl.COLOR_BUFFER_BIT)

	glctx.UseProgram(program)

	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	glctx.EnableVertexAttribArray(position)
	glctx.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	glctx.DrawArrays(gl.TRIANGLES, 0, 3)
	glctx.DisableVertexAttribArray(position)
}
var triangleData = f32.Bytes(binary.LittleEndian,
	-0.5, 0.5,
	0.5, 0.5,
	0.5, -0.5,
)
