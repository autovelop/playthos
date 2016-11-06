package systems

import (
	"fmt"

	"gde"
	components "gde/components"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	VSHADER_OPENGL_4_1 = `#version 410
  in vec4 position;
  void main()
  {
	gl_Position = position;
  }
  ` + "\x00"

	FSHADER_OPENGL_4_1 = `#version 410
  out vec4 outColor;
  void main()
  {
	outColor = vec4(1.0, 1.0, 1.0, 1.0);
  }
  ` + "\x00"

	// layout (location = 0) attribute vec4 position;
	VSHADER_OPENGL_ES_2_0 = `#version 100
  attribute vec4 position;
  uniform mat4 transform;
  void main() {
	gl_Position = transform * position;
  }`

	FSHADER_OPENGL_ES_2_0 = `#version 100
  precision mediump float;
  uniform vec4 color;
  void main() {
	gl_FragColor = color;
  }`
)

type Render struct {
	gde.SystemRoutine

	Program       uint32
	Window        *glfw.Window
	VertexArrayId uint32
	ShaderProgram uint32
}

func key_callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	fmt.Println("Render.Update() executed")
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func (r *Render) Init() {
	fmt.Println("Render.Init() executed")
	// https://github.com/viking/gl-tutorial
	// https://github.com/gdm85/wolfengo

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	var window *glfw.Window
	window, err := glfw.CreateWindow(360, 640, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(key_callback)
	r.Window = window

	// gl.Viewport(0, 0, 360, 360)

	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(VSHADER_OPENGL_ES_2_0)
	gl.ShaderSource(vshader, 1, vsources, nil)
	vfree()
	gl.CompileShader(vshader)

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(FSHADER_OPENGL_ES_2_0)
	gl.ShaderSource(fshader, 1, fsources, nil)
	ffree()
	gl.CompileShader(fshader)

	// Create program
	var shaderProgram uint32
	shaderProgram = gl.CreateProgram()

	gl.AttachShader(shaderProgram, fshader)
	gl.AttachShader(shaderProgram, vshader)

	// Link program
	gl.LinkProgram(shaderProgram)

	// Use this program for all upcoming render calls
	gl.UseProgram(shaderProgram)

	r.ShaderProgram = shaderProgram

	// Delete all shaders
	gl.DeleteShader(vshader)
	gl.DeleteShader(fshader)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	// gl.Enable(gl.DEPTH_TEST)
}
func (r *Render) Update(entities *map[string]gde.EntityRoutine) {
	if !r.Window.ShouldClose() {
		fmt.Println("Render.Update() executed")

		gl.ClearColor(0.2, 0.3, 0.3, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(r.ShaderProgram)

		transformLoc := gl.GetUniformLocation(r.ShaderProgram, gl.Str("transform\x00"))
		for _, v := range *entities {

			vao := v.Component(fmt.Sprintf("%T", &components.Renderer{})).GetProperty("VAO")
			switch vao := vao.(type) {
			case uint32:
				gl.BindVertexArray(vao)
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

			gl.UniformMatrix4fv(transformLoc, 1, false, &transform[0])

			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		}

		gl.BindVertexArray(0)

		r.Window.SwapBuffers()
		glfw.PollEvents()
	} else {
		glfw.Terminate()
		// once thid happens, make sure render system is removed from engine so that another game loop doesn't occur
	}
}
func (r *Render) End() {
	fmt.Println("Render.End() executed")
}

func (r *Render) Add(engine *gde.Engine) {
	fmt.Println("Render.Add(Engine) executed")
	engine.Systems[fmt.Sprintf("%T", Render{})] = r
}
