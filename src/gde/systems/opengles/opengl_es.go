package opengles

import (
	"fmt"
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"

	"gde"
	"gde/components"
	"gde/systems"
)

type RenderOpenGLES struct {
	systems.Render

	ShaderProgram gl.Program
	Context       gl.Context
	Size          size.Event
}

func (r *RenderOpenGLES) Init() {
	log.Printf("----------------- OpenGLES Init ----------------- ")
	r.OpenGL = "OpenGLES 2"
	r.GLSL = "100"

	// Create program
	var shaderProgram gl.Program
	var err error
	shaderProgram, err = glutil.CreateProgram(r.Context, VSHADER_OPENGL_ES_2_0, FSHADER_OPENGL_ES_2_0)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}
	r.ShaderProgram = shaderProgram

	// Use this program for all upcoming render calls
	r.Context.UseProgram(r.ShaderProgram)
}

func (r *RenderOpenGLES) Update(entities *map[string]gde.EntityRoutine) {
	r.Context.ClearColor(0.2, 0.3, 0.3, 1)
	r.Context.Clear(gl.COLOR_BUFFER_BIT)

	r.Context.UseProgram(r.ShaderProgram)

	transformLoc := r.Context.GetUniformLocation(r.ShaderProgram, "transform")
	for _, v := range *entities {
		vb := v.Component(fmt.Sprintf("%T", &components.Renderer{})).GetProperty("VB")
		switch vb := vb.(type) {
		case gl.Buffer:
			r.Context.BindBuffer(gl.ARRAY_BUFFER, vb)
		}

		eb := v.Component(fmt.Sprintf("%T", &components.Renderer{})).GetProperty("EB")
		switch eb := eb.(type) {
		case gl.Buffer:
			r.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, eb)
		}

		position := v.Component(fmt.Sprintf("%T", &components.Renderer{})).GetProperty("EB")
		switch position := position.(type) {
		case gl.Attrib:
			// Need to figure out again what this is for...
			r.Context.EnableVertexAttribArray(position)
			r.Context.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
		}

		var transform mgl32.Mat4
		trans := v.Component(fmt.Sprintf("%T", &components.Transform{}))

		pos := trans.GetProperty("Position")
		switch pos := pos.(type) {
		case mgl32.Vec3:
			transform = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())
		}

		rot := trans.GetProperty("Rotation")
		switch rot := rot.(type) {
		case mgl32.Vec3:
			transform = transform.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rot.X())).Mat4())
			transform = transform.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rot.Y())).Mat4())
			transform = transform.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rot.Z())).Mat4())
		}

		r.Context.UniformMatrix4fv(transformLoc, transform[:])
		r.Context.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)
	}
}

func (r *RenderOpenGLES) LoadRenderer(renderer *components.Renderer) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
	var vertexBuffer = r.Context.CreateBuffer()
	r.Context.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	r.Context.BufferData(gl.ARRAY_BUFFER, renderer.Mesh.VerticesByteArray(), gl.STATIC_DRAW)
	renderer.SetProperty("VB", vertexBuffer)

	var elementBuffer = r.Context.CreateBuffer()
	r.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	r.Context.BufferData(gl.ELEMENT_ARRAY_BUFFER, renderer.Mesh.IndiciesByteArray(), gl.STATIC_DRAW)
	renderer.SetProperty("EB", elementBuffer)

	position := r.Context.GetAttribLocation(r.ShaderProgram, "position")
	r.Context.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	r.Context.EnableVertexAttribArray(position)
	renderer.SetProperty("POSITION", position)
}

func (r *RenderOpenGLES) Add(engine *gde.Engine) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
	engine.Systems[fmt.Sprintf("%T", &systems.Render{})] = systems.RenderRoutine(r)
}
func (r *RenderOpenGLES) Shutdown() {
}

const (
	VSHADER_OPENGL_ES_2_0 = `#version 100
  attribute vec4 position;
  uniform mat4 transform;
  void main() {
	gl_Position = transform * position;
  }`

	FSHADER_OPENGL_ES_2_0 = `#version 100
  precision mediump float;
  void main() {
	gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
  }`
)
