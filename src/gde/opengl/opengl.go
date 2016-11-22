package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"gde"
)

type RenderOpenGL struct {
	gde.Render

	window        *glfw.Window
	VertexArrayId uint32
	ShaderProgram uint32
}

func (r *RenderOpenGL) Init() {
	r.OpenGL = "OpenGL 4"
	r.GLSL = "100"

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// var window *glfw.Window
	window, err := glfw.CreateWindow(360, 640, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	r.window = window

	gl.Viewport(0, 0, 360, 640)

	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(gde.VSHADER_OPENGL_ES_2_0)
	gl.ShaderSource(vshader, 1, vsources, nil)
	vfree()
	gl.CompileShader(vshader)

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(gde.FSHADER_OPENGL_ES_2_0)
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

	if glError := gl.GetError(); glError != 0 {
		fmt.Printf("gl.GetError: %v", glError)
		return
	}
	// gl.Enable(gl.DEPTH_TEST)
}

func (r *RenderOpenGL) Update(entities *map[string]gde.EntityRoutine) {
	// fmt.Println("Render.Update() started")
	if !r.window.ShouldClose() {
		// fmt.Println("Render.Update() executed")

		gl.ClearColor(0.2, 0.3, 0.3, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(r.ShaderProgram)

		var view mgl32.Mat4
		view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		view_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("view\x00"))

		var proj mgl32.Mat4
		proj = mgl32.Ortho(0, 1, 2, 0, 0.1, 1000)
		// proj = mgl32.Perspective(mgl32.DegToRad(60.0), float32(320)/640, 0.01, 1000)
		proj_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("projection\x00"))

		model_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("model\x00"))

		for _, v := range *entities {
			vao := v.Component(fmt.Sprintf("%T", &gde.Renderer{})).GetProperty("VAO")
			switch vao := vao.(type) {
			case uint32:
				gl.BindVertexArray(vao)
			}

			var model mgl32.Mat4
			trans := v.Component(fmt.Sprintf("%T", &gde.Transform{}))

			pos := trans.GetProperty("Position")
			switch pos := pos.(type) {
			case gde.Vector3:
				model = mgl32.Translate3D(pos.X, pos.Y, pos.Z)
			}

			rot := trans.GetProperty("Rotation")
			switch rot := rot.(type) {
			case gde.Vector3:
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rot.X)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rot.Y)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rot.Z)).Mat4())
			}

			gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
			gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
			gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])

			texture := v.Component(fmt.Sprintf("%T", &gde.Renderer{})).GetProperty("TEXTURE")
			switch texture := texture.(type) {
			case uint32:
				// fmt.Printf("here %v \n\n\n", texture)
				gl.ActiveTexture(gl.TEXTURE0)
				gl.BindTexture(gl.TEXTURE_2D, texture)
				gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("texture\x00")), 0)
			}

			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
		}

		gl.BindVertexArray(0)

		r.window.SwapBuffers()
		glfw.PollEvents()
	} else {
		glfw.Terminate()
		// when thid happens, make sure render system is removed from engine so that another game loop doesn't occur
	}
}
func (r *RenderOpenGL) LoadRenderer(renderer *gde.Renderer) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
	// Bind vertex array object. This must wrap around the mesh creation because it is how we are going to access it later when we draw
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(renderer.Mesh.Vertices)*4, gl.Ptr(renderer.Mesh.Vertices), gl.STATIC_DRAW)

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(renderer.Mesh.Indicies)*4, gl.Ptr(renderer.Mesh.Indicies), gl.STATIC_DRAW)

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Linking texture attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	// Unbind Vertex array object
	renderer.SetProperty("VAO", vertexArrayID)
	// fmt.Printf("VAO Load: %v\n", vertexArrayID)
	gl.BindVertexArray(0)

	// Load texture
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	fmt.Println(renderer.Texture)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(renderer.Texture.RGBA.Rect.Size().X),
		int32(renderer.Texture.RGBA.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(renderer.Texture.RGBA.Pix))
	renderer.SetProperty("TEXTURE", texture)
	fmt.Printf("here %v \n\n\n", texture)
}

func (r *RenderOpenGL) Stop() {
}

func (r *RenderOpenGL) GetWindow() *glfw.Window {
	return r.window
}
