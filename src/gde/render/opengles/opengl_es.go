package opengles

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"

	"gde/engine"
	"gde/render"
	"gde/render/ui/uigles"
)

type OpenGLES struct {
	render.RenderRoutine

	ShaderProgram uint32
	glProgram     gl.Program
	Context       gl.Context
	Size          size.Event

	uiSystem render.RenderRoutine
}

func (r *OpenGLES) Init() {
	log.Printf("OpenGLES > Init")
	r.ShaderProgram = r.NewShader(render.VSHADER_OPENGL_ES_2_0, render.FSHADER_OPENGL_ES_2_0)
	// r.Context.Viewport(0, 0, 720, 1280)
	r.Context.Enable(gl.DEPTH_TEST)
	r.Context.Enable(gl.FRONT_AND_BACK)
}

func (r *OpenGLES) Update(entities *map[string]*engine.Entity) {
	log.Printf("\n\n\nOpenGLES UPDATE %+v\n\n\n", r.Context)
	r.Context.ClearColor(0.2, 0.3, 0.3, 1)
	r.Context.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	r.Context.UseProgram(r.glProgram)

	var view mgl32.Mat4
	view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	view_uni := r.Context.GetUniformLocation(r.glProgram, "view")

	var proj mgl32.Mat4
	// still need to understand what is going on here and how it relates to device
	proj = mgl32.Ortho(-2, 3, 6, -4, 0.1, 1000)

	proj_uni := r.Context.GetUniformLocation(r.glProgram, "projection")
	model_uni := r.Context.GetUniformLocation(r.glProgram, "model")

	for _, v := range *entities {
		renderer := v.GetComponent(&render.MeshRenderer{})
		if renderer == nil {
			continue
		}
		vb := renderer.GetProperty("VB")
		switch vb := vb.(type) {
		case gl.Buffer:
			r.Context.BindBuffer(gl.ARRAY_BUFFER, vb)
		}

		eb := renderer.GetProperty("EB")
		switch eb := eb.(type) {
		case gl.Buffer:
			r.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, eb)
		}

		// position := renderer.GetProperty("EB")
		// switch position := position.(type) {
		// case gl.Attrib:
		// 	// Need to figure out again what this is for...
		// 	r.Context.EnableVertexAttribArray(position)
		// 	r.Context.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
		// }

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
		case gl.Texture:
			r.Context.ActiveTexture(gl.TEXTURE0)
			r.Context.BindTexture(gl.TEXTURE_2D, texture)
			r.Context.Uniform1i(r.Context.GetUniformLocation(r.glProgram, "texture"), 0)
		}

		r.Context.UniformMatrix4fv(model_uni, model[:])
		r.Context.UniformMatrix4fv(view_uni, view[:])
		r.Context.UniformMatrix4fv(proj_uni, proj[:])

		r.Context.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)
	}

	// let the ui system flush if it exists
	if r.uiSystem == nil {
		r.Context.Flush()
	}
}

func (r *OpenGLES) LoadRenderer(renderer render.RendererRoutine) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
	// Vertex buffer
	var vertexBuffer = r.Context.CreateBuffer()
	r.Context.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	r.Context.BufferData(gl.ARRAY_BUFFER, renderer.MeshByteVertices(), gl.STATIC_DRAW)
	renderer.SetProperty("VB", vertexBuffer)

	// Element buffer
	var elementBuffer = r.Context.CreateBuffer()
	r.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	r.Context.BufferData(gl.ELEMENT_ARRAY_BUFFER, renderer.MeshByteIndicies(), gl.STATIC_DRAW)
	renderer.SetProperty("EB", elementBuffer)

	// Linking vertex attributes
	posAttrib := r.Context.GetAttribLocation(r.glProgram, "pos")
	r.Context.VertexAttribPointer(posAttrib, 3, gl.FLOAT, false, 8*4, 0)
	r.Context.EnableVertexAttribArray(posAttrib)

	// Linking fragment attributes
	colAttrib := r.Context.GetAttribLocation(r.glProgram, "col")
	r.Context.VertexAttribPointer(colAttrib, 3, gl.FLOAT, false, 8*4, 3*4)
	r.Context.EnableVertexAttribArray(colAttrib)

	// Linking texture attributes
	texAttrib := r.Context.GetAttribLocation(r.glProgram, "tex")
	r.Context.VertexAttribPointer(texAttrib, 2, gl.FLOAT, false, 8*4, 6*4)
	r.Context.EnableVertexAttribArray(texAttrib)

	texture := r.Context.CreateTexture()

	r.Context.BindTexture(gl.TEXTURE_2D, texture)
	r.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	r.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	r.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	r.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	r.Context.TexImage2D(
		gl.TEXTURE_2D,
		0,
		int(renderer.TextureRGBA().Rect.Size().X),
		int(renderer.TextureRGBA().Rect.Size().Y),
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		renderer.TextureRGBA().Pix)
	renderer.SetProperty("TEXTURE", texture)
}

func (r *OpenGLES) AddUISystem(game *engine.Engine) {
	// Create ui system
	log.Printf("\n\n\nOpenGLES: %+v\n\n\n", r.Context)
	sys_ui := &uigles.UIGLES{Context: r.Context}
	game.AddSystem(engine.SystemUI, sys_ui)
	sys_ui.Init()
	r.uiSystem = sys_ui
}

func (r *OpenGLES) Stop() {
	r.Context.DeleteProgram(r.glProgram)
	// CLEAN UP BUFFERS
	// r.Context.DeleteBuffer(buf)
}

func (r *OpenGLES) NewShader(vShader string, fShader string) uint32 {
	version := gl.Version()
	glsl_version := r.Context.GetString(gl.SHADING_LANGUAGE_VERSION)
	log.Printf("Render > OpenGL > Version: %v", version)
	log.Printf("Render > OpenGL > GLSL Version: %v", glsl_version)
	// Create vertex shader
	vshader := r.Context.CreateShader(gl.VERTEX_SHADER)
	if vshader.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "Could not create VERTEXT_SHADER", vShader)
	}
	r.Context.ShaderSource(vshader, vShader)
	r.Context.CompileShader(vshader)
	defer r.Context.DeleteShader(vshader)
	if r.Context.GetShaderi(vshader, gl.COMPILE_STATUS) == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", r.Context.GetShaderInfoLog(vshader), vShader)
	}

	// Create fragment shader
	fshader := r.Context.CreateShader(gl.FRAGMENT_SHADER)
	if fshader.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "Could not create FRAGMENT_SHADER", fShader)
	}
	r.Context.ShaderSource(fshader, fShader)
	r.Context.CompileShader(fshader)
	defer r.Context.DeleteShader(fshader)
	if r.Context.GetShaderi(fshader, gl.COMPILE_STATUS) == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", r.Context.GetShaderInfoLog(fshader), fShader)
	}

	shaderProgram := r.Context.CreateProgram()
	r.glProgram = shaderProgram
	if shaderProgram.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "No GLES program available")
	}

	r.Context.AttachShader(shaderProgram, vshader)
	r.Context.AttachShader(shaderProgram, fshader)

	r.Context.LinkProgram(shaderProgram)

	// Flag shaders for deletion when program is unlinked.
	r.Context.DeleteShader(vshader)
	r.Context.DeleteShader(fshader)

	if r.Context.GetProgrami(shaderProgram, gl.LINK_STATUS) == 0 {
		defer r.Context.DeleteProgram(shaderProgram)
		log.Printf("\n\n ### SHADER LINK ERROR ### \n%v\n\n", r.Context.GetProgramInfoLog(shaderProgram))
	}

	r.Context.UseProgram(shaderProgram)

	return shaderProgram.Value
}
