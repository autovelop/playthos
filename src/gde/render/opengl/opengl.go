package opengl

import (
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"gde/engine"
	"gde/render"
	"gde/render/ui/uigl"
)

type OpenGL struct {
	render.RenderRoutine

	window        *glfw.Window
	ShaderProgram uint32

	uiSystem render.RenderRoutine
	stepper  float32
}

func (r *OpenGL) Init() {
	log.Printf("OpenGL > Init")
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// window, err := glfw.CreateWindow(int(r.Device.ScreenW), int(r.Device.ScreenH), "Cube", nil, nil)
	window, err := glfw.CreateWindow(360, 640, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	r.window = window

	// gl.Viewport(0, 0, r.Device.RenderW, r.Device.RenderH)
	gl.Viewport(0, 0, 360, 640)

	r.ShaderProgram = r.NewShader(render.VSHADER_OPENGL_ES_2_0, render.FSHADER_OPENGL_ES_2_0)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.FRONT_AND_BACK)
}

func (r *OpenGL) Update(entities *map[string]*engine.Entity) {
	if !r.window.ShouldClose() {

		gl.ClearColor(0.2, 0.3, 0.3, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(r.ShaderProgram)

		r.stepper += 0.1
		if r.stepper > 2 {
			r.stepper = 1.0
		}

		var view mgl32.Mat4
		view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		view_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("view\x00"))

		var proj mgl32.Mat4
		proj = mgl32.Ortho(0, 1, 2, 0, 0.1, 1000)

		proj_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("projection\x00"))
		model_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("model\x00"))

		for _, v := range *entities {
			renderer := v.GetComponent(&render.MeshRenderer{})
			if renderer == nil {
				continue
			}
			vao := renderer.GetProperty("VAO")
			switch vao := vao.(type) {
			case uint32:
				gl.BindVertexArray(vao)
			}

			var model mgl32.Mat4
			trans := v.GetComponent(&render.Transform{})

			pos := trans.GetProperty("Position")
			switch pos := pos.(type) {
			case render.Vector3:
				model = mgl32.Translate3D(pos.X, pos.Y, pos.Z)
			}

			rot := trans.GetProperty("Rotation")
			switch rot := rot.(type) {
			case render.Vector3:
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rot.X)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rot.Y)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rot.Z)).Mat4())
			}

			texture := renderer.GetProperty("TEXTURE")
			switch texture := texture.(type) {
			case uint32:
				gl.ActiveTexture(gl.TEXTURE0)
				gl.BindTexture(gl.TEXTURE_2D, texture)
				gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("texture\x00")), 0)
			}

			gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
			gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
			gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])

			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
		}
		gl.BindVertexArray(0)

		// let the UI system swap the buffers and do poll events only if it exists
		if r.uiSystem == nil {
			r.window.SwapBuffers()
			glfw.PollEvents()
		}

	} else {
		glfw.Terminate()
		// when this happens, make sure render system is removed from engine so that another game loop doesn't occur
	}
}
func (r *OpenGL) LoadRenderer(renderer render.RendererRoutine) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
	// Bind vertex array object. This must wrap around the mesh creation because it is how we are going to access it later when we draw
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(renderer.MeshVertices())*4, gl.Ptr(renderer.MeshVertices()), gl.STATIC_DRAW)

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(renderer.MeshIndicies())*4, gl.Ptr(renderer.MeshIndicies()), gl.STATIC_DRAW)

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Linking texture attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	renderer.SetProperty("VAO", vertexArrayID)

	// Unbind Vertex array object
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
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(renderer.TextureRGBA().Rect.Size().X),
		int32(renderer.TextureRGBA().Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(renderer.TextureRGBA().Pix))
	renderer.SetProperty("TEXTURE", texture)
}

func (r *OpenGL) AddUISystem(game *engine.Engine) {
	log.Printf("\n\n\nOpenGL\n\n\n")
	// Create ui system
	sys_ui := &uigl.UIGL{Window: r.window}
	game.AddSystem(engine.SystemUI, sys_ui)
	sys_ui.Init()
	r.uiSystem = sys_ui
}

func (r *OpenGL) Stop() {
}

func (r *OpenGL) GetWindow() *glfw.Window {
	return r.window
}

func (r *OpenGL) NewShader(vShader string, fShader string) uint32 {
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("Render > OpenGL > Version: %v", version)
	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(vShader)
	gl.ShaderSource(vshader, 1, vsources, nil)
	vfree()
	gl.CompileShader(vshader)
	defer gl.DeleteShader(vshader)

	var vstatus int32
	gl.GetShaderiv(vshader, gl.COMPILE_STATUS, &vstatus)
	if vstatus == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vshader, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vshader, logLength, nil, gl.Str(logMsg))

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, vShader)
		os.Exit(0)
	}

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(fShader)
	gl.ShaderSource(fshader, 1, fsources, nil)
	ffree()
	gl.CompileShader(fshader)
	defer gl.DeleteShader(fshader)

	var fstatus int32
	gl.GetShaderiv(fshader, gl.COMPILE_STATUS, &fstatus)
	if fstatus == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fshader, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fshader, logLength, nil, gl.Str(logMsg))

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, fShader)
		os.Exit(0)
	}

	// Create program
	var shaderProgram uint32
	shaderProgram = gl.CreateProgram()

	gl.AttachShader(shaderProgram, vshader)
	gl.AttachShader(shaderProgram, fshader)

	// Link program
	gl.LinkProgram(shaderProgram)

	var statisLink int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &statisLink)
	if statisLink == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(logMsg))

		log.Printf("\n\n ### SHADER LINK ERROR ### \n%v\n\n", logMsg)
		os.Exit(0)
	}

	// Use this program for all upcoming render calls
	gl.UseProgram(shaderProgram)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	return shaderProgram
}
