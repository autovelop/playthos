// +build opengl

package opengl

import (
	"./../engine"
	"./../render"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"os"
	"strings"
)

func init() {
	render.activeRenderSystem = OpenGL
}

type OpenGL struct {
	window        *glfw.Window
	screenWidth   float32
	screenHeight  float32
	shaderProgram uint32
	transforms    []*render.Transform
	meshes        []*render.Mesh
	materials     []*render.Material
}

func (o *OpenGL) Prepare() {
	log.Println("OpenGL Prepare")
	o.screenWidth = 1024
	o.screenHeight = 768

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// 	// window, err := glfw.CreateWindow(int(r.Device.ScreenW), int(r.Device.ScreenH), "Cube", nil, nil)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	window, err := glfw.CreateWindow(int(o.screenWidth), int(o.screenHeight), "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	o.window = window

	gl.Viewport(0, 0, int32(o.screenWidth), int32(o.screenHeight))

	o.shaderProgram = o.NewShader(render.VSHADER, render.FSHADER)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.FRONT_AND_BACK)

	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.2, 0.4, 1)
}

func (o *OpenGL) ComponentTypes() []engine.Component {
	return []engine.Component{&render.Transform{}, &render.Mesh{}, &render.Material{}}
}
func (o *OpenGL) Update() {
	// log.Println("OpenGL Update")
	// for _, transform := range r.transforms {
	// 	log.Printf("Transform Update: %v\n", transform)
	// }
	if !o.window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(o.shaderProgram)

		var view mgl32.Mat4

		// 		lookat := r.Camera.GetProperty("LookAt")
		// 		switch lookat := lookat.(type) {
		// 		case render.Vector3:
		// 			lookfrom := r.Camera.GetProperty("LookFrom")
		// 			switch lookfrom := lookfrom.(type) {
		// 			case render.Vector3:
		view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0.0, 1}, mgl32.Vec3{0, 1, 0})
		// 			}
		// 		}

		view_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("view\x00"))

		var proj mgl32.Mat4
		ratio := o.screenWidth / o.screenHeight
		proj = mgl32.Ortho(0, o.screenWidth/ratio, o.screenHeight/ratio, 0, -1.0, 1000.0)

		proj_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("projection\x00"))
		model_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("model\x00"))

		for idx, mesh := range o.meshes {
			gl.BindVertexArray(mesh.GetVAO())

			transform := o.transforms[idx]
			position := transform.GetPosition()
			rotation := transform.GetRotation()
			scale := transform.GetScale()
			material := o.materials[idx]
			color := material.GetColor()
			model := mgl32.Ident4()
			model = model.Mul4(mgl32.Translate3D(position.X, position.Y, position.Z))
			model = model.Mul4(mgl32.Scale3D(100*scale.X, 100*scale.Y, scale.Z))
			model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X)).Mat4())
			model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y)).Mat4())
			model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z)).Mat4())
			// 				}

			// 				texture := renderer.GetProperty("Texture")
			// 				if texture != nil {
			// 					gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("hasTexture\x00")), 1)

			// 					switch texture := texture.(type) {
			// 					case uint32:
			// 						gl.ActiveTexture(gl.TEXTURE0)
			// 						gl.BindTexture(gl.TEXTURE_2D, texture)
			// 						gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("texture\x00")), 0)
			// 					}
			// 				} else {
			// 					gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("hasTexture\x00")), 0)
			// 				}
			gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("hasTexture\x00")), 0)

			// 				color := renderer.GetProperty("Color")
			// 				if color != nil {
			// 					switch color := color.(type) {
			// 					case *render.Color:
			gl.Uniform4fv(gl.GetUniformLocation(o.shaderProgram, gl.Str("color\x00")), 1, &color[0])
			// 					}
			// 				}

			gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
			gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
			gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])
			// 				// log.Printf("OpenGL > Draw > Model: %v", model)
			// 				// log.Printf("OpenGL > Draw > View: %v", view)
			// 				// log.Printf("OpenGL > Draw > Projection: %v", proj)

			gl.DrawElements(gl.TRIANGLES, 8, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
			// 				gl.BindVertexArray(0)
			// 			} else {
			// 				switch renderer := renderer.(type) {
			// 				case *render.MeshRenderer:
			// 					r.LoadRenderer(renderer)
			// 				}
			// 			}
			// 		}
			gl.BindVertexArray(0)
		}

		// 		// let the UI system swap the buffers and do poll events only if it exists
		// 		// if r.uiSystem == nil {
		o.window.SwapBuffers()
		glfw.PollEvents()
		// }

	} else {
		glfw.Terminate()
		// 		// when this happens, make sure render system is removed from engine so that another game loop doesn't occur
	}
}

func (o *OpenGL) NewShader(vs string, fs string) uint32 {
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("Render > OpenGL > Version: %v", version)
	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(vs)
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

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, vs)
		os.Exit(0)
	}

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(fs)
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

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, fs)
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

func (o *OpenGL) RegisterTransform(transform *render.Transform) {
	o.transforms = append(o.transforms, transform)
}

func (o *OpenGL) RegisterMaterial(material *render.Material) {
	o.materials = append(o.materials, material)
}

func (o *OpenGL) RegisterMesh(mesh *render.Mesh) {
	var verticies []float32 = mesh.GetVertices()
	var indicies []uint8 = mesh.GetIndicies()

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)
	log.Printf("LoadRenderer > VAO > ID: %v", vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, gl.Ptr(verticies), gl.STATIC_DRAW)
	log.Printf("LoadRenderer > VBO > Verticies Length: %v", len(verticies))

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)
	log.Printf("LoadRenderer > EBO > Indicies Length: %v", len(indicies))

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Linking texture attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	// if renderer.HasTexture() {
	// 	// Load texture
	// 	var texture uint32
	// 	gl.GenTextures(1, &texture)
	// 	gl.ActiveTexture(gl.TEXTURE0)
	// 	gl.BindTexture(gl.TEXTURE_2D, texture)
	// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	// 	gl.TexImage2D(
	// 		gl.TEXTURE_2D,
	// 		0,
	// 		gl.RGBA,
	// 		int32(renderer.GetTextureRGBA().Rect.Size().X),
	// 		int32(renderer.GetTextureRGBA().Rect.Size().Y),
	// 		0,
	// 		gl.RGBA,
	// 		gl.UNSIGNED_BYTE,
	// 		gl.Ptr(renderer.GetTextureRGBA().Pix))
	// 	log.Printf("LoadRenderer > Texture > ID: %v", texture)
	// 	log.Printf("LoadRenderer > Texture > Width: %v", int32(renderer.GetTextureRGBA().Rect.Size().X))
	// 	log.Printf("LoadRenderer > Texture > Height: %v", int32(renderer.GetTextureRGBA().Rect.Size().Y))
	// 	log.Printf("LoadRenderer > Texture > Pix Length: %v", len(renderer.GetTextureRGBA().Pix))
	// 	renderer.SetProperty("Texture", texture)

	// 	color := renderer.GetColor()
	// 	if color != nil {
	// 		renderer.SetProperty("Color", color)
	// 	}
	// } else {
	// 	color := renderer.GetColor()
	// 	if color != nil {
	// 		renderer.SetProperty("Color", color)
	// 	}
	// }

	// renderer.SetProperty("VAO", vertexArrayID)
	mesh.SetVAO(vertexArrayID)

	// Unbind Vertex array object
	gl.BindVertexArray(0)
	o.meshes = append(o.meshes, mesh)
}

// import (
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/go-gl/gl/v4.1-core/gl"
// 	"github.com/go-gl/glfw/v3.2/glfw"
// 	"github.com/go-gl/mathgl/mgl32"

// 	"gde/engine"
// 	"gde/render"
// 	"gde/render/ui/uigl"
// )

// type OpenGL struct {
// 	render.RenderRoutine

// 	Window        *glfw.Window
// 	ShaderProgram uint32

// 	screenWidth  float32
// 	screenHeight float32

// 	Camera *render.Camera

// 	uiSystem render.RenderRoutine

// 	registeredRenderer func()
// }

// func (r *OpenGL) Init() {
// 	log.Printf("OpenGL > Init")

// 	r.screenWidth = 1024
// 	r.screenHeight = 768
// 	// r.screenWidth = 768
// 	// r.screenHeight = 1024

// 	log.Printf("OpenGL > Screen > %v x %v", r.screenWidth, r.screenHeight)
// 	// Initialize Glow
// 	if err := gl.Init(); err != nil {
// 		panic(err)
// 	}

// 	r.Camera = &render.Camera{}
// 	r.Camera.Init()

// 	// window, err := glfw.CreateWindow(int(r.Device.ScreenW), int(r.Device.ScreenH), "Cube", nil, nil)
// 	glfw.WindowHint(glfw.ContextVersionMajor, 3)
// 	glfw.WindowHint(glfw.ContextVersionMinor, 3)
// 	window, err := glfw.CreateWindow(int(r.screenWidth), int(r.screenHeight), "Cube", nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	window.MakeContextCurrent()
// 	r.Window = window

// 	// gl.Viewport(0, 0, r.Device.RenderW, r.Device.RenderH)
// 	// gl.Viewport(0, 0, r.screenWidth, r.screenHeight)
// 	gl.Viewport(0, 0, int32(r.screenWidth), int32(r.screenHeight))

// 	r.ShaderProgram = r.NewShader(render.VSHADER_OPENGL_ES_2_0, render.FSHADER_OPENGL_ES_2_0)

// 	gl.Enable(gl.DEPTH_TEST)
// 	gl.DepthFunc(gl.LEQUAL)
// 	gl.Enable(gl.FRONT_AND_BACK)

// 	gl.Enable(gl.BLEND)
// 	gl.BlendEquation(gl.FUNC_ADD)
// 	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

// 	// gl.AlphaFunc(gl.GREATER, 0.005)
// 	gl.ClearColor(0.1, 0.2, 0.4, 1)
// }

// func (r *OpenGL) Update(entities *map[string]*engine.Entity) {
// 	if !r.Window.ShouldClose() {

// 		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

// 		gl.UseProgram(r.ShaderProgram)

// 		var view mgl32.Mat4

// 		lookat := r.Camera.GetProperty("LookAt")
// 		switch lookat := lookat.(type) {
// 		case render.Vector3:
// 			lookfrom := r.Camera.GetProperty("LookFrom")
// 			switch lookfrom := lookfrom.(type) {
// 			case render.Vector3:
// 				view = mgl32.LookAtV(mgl32.Vec3{lookat.X, lookat.Y, lookat.Z}, mgl32.Vec3{lookfrom.X, lookfrom.Y, lookfrom.Z}, mgl32.Vec3{0, 1, 0})
// 			}
// 		}

// 		view_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("view\x00"))

// 		var proj mgl32.Mat4
// 		ratio := r.screenWidth / r.screenHeight
// 		// ratio := r.screenHeight / r.screenWidth
// 		// proj = mgl32.Ortho(-float32(r.screenWidth/r.screenHeight), float32(r.screenWidth/r.screenHeight), -float32(r.screenWidth/r.screenHeight), float32(r.screenWidth/r.screenHeight), 0.1, 1000)
// 		proj = mgl32.Ortho(0, r.screenWidth/ratio, r.screenHeight/ratio, 0, -1.0, 1000.0)
// 		// proj = mgl32.Ortho(0, r.screenWidth, r.screenHeight, 0, -1.0, 10.0)

// 		proj_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("projection\x00"))
// 		model_uni := gl.GetUniformLocation(r.ShaderProgram, gl.Str("model\x00"))

// 		for _, v := range *entities {
// 			renderer := v.GetComponent("MeshRenderer")
// 			if renderer == nil {
// 				continue
// 			}

// 			vao := renderer.GetProperty("VAO")
// 			if vao != nil {
// 				switch vao := vao.(type) {
// 				case uint32:
// 					gl.BindVertexArray(vao)
// 					// log.Printf("OpenGL > Draw > VAO: %v", vao)
// 				}

// 				// var model mgl32.Mat4
// 				model := mgl32.Ident4()
// 				trans := v.GetComponent("Transform")
// 				if trans == nil {
// 					continue
// 				}

// 				// defaults

// 				pos := trans.GetProperty("Position")
// 				switch pos := pos.(type) {
// 				case render.Vector3:
// 					model = model.Mul4(mgl32.Translate3D(pos.X, pos.Y, pos.Z))
// 					// log.Printf("OpenGL > Draw > Position: %v", pos)
// 				}

// 				scale := trans.GetProperty("Scale")
// 				switch scale := scale.(type) {
// 				case render.Vector3:
// 					model = model.Mul4(mgl32.Scale3D(100*scale.X, 100*scale.Y, scale.Z))
// 					// log.Printf("OpenGL > Draw > Position: %v", pos)
// 				}

// 				rot := trans.GetProperty("Rotation")
// 				switch rot := rot.(type) {
// 				case render.Vector3:
// 					model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rot.X)).Mat4())
// 					model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rot.Y)).Mat4())
// 					model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rot.Z)).Mat4())
// 				}

// 				texture := renderer.GetProperty("Texture")
// 				if texture != nil {
// 					gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("hasTexture\x00")), 1)

// 					switch texture := texture.(type) {
// 					case uint32:
// 						gl.ActiveTexture(gl.TEXTURE0)
// 						gl.BindTexture(gl.TEXTURE_2D, texture)
// 						gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("texture\x00")), 0)
// 					}
// 				} else {
// 					gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("hasTexture\x00")), 0)
// 				}
// 				// gl.Uniform1i(gl.GetUniformLocation(r.ShaderProgram, gl.Str("hasTexture\x00")), 0)

// 				color := renderer.GetProperty("Color")
// 				if color != nil {
// 					switch color := color.(type) {
// 					case *render.Color:
// 						gl.Uniform4fv(gl.GetUniformLocation(r.ShaderProgram, gl.Str("color\x00")), 1, &color[0])
// 					}
// 				}

// 				gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
// 				gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
// 				gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])
// 				// log.Printf("OpenGL > Draw > Model: %v", model)
// 				// log.Printf("OpenGL > Draw > View: %v", view)
// 				// log.Printf("OpenGL > Draw > Projection: %v", proj)

// 				gl.DrawElements(gl.TRIANGLES, 8, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
// 				gl.BindVertexArray(0)
// 			} else {
// 				switch renderer := renderer.(type) {
// 				case *render.MeshRenderer:
// 					r.LoadRenderer(renderer)
// 				}
// 			}
// 		}

// 		// let the UI system swap the buffers and do poll events only if it exists
// 		// if r.uiSystem == nil {
// 		r.Window.SwapBuffers()
// 		glfw.PollEvents()
// 		// }

// 	} else {
// 		glfw.Terminate()
// 		// when this happens, make sure render system is removed from engine so that another game loop doesn't occur
// 	}
// }
// func (r *OpenGL) NewShader(vShader string, fShader string) uint32 {
// 	version := gl.GoStr(gl.GetString(gl.VERSION))
// 	log.Printf("Render > OpenGL > Version: %v", version)
// 	// Create vertex shader
// 	vshader := gl.CreateShader(gl.VERTEX_SHADER)
// 	vsources, vfree := gl.Strs(vShader)
// 	gl.ShaderSource(vshader, 1, vsources, nil)
// 	vfree()
// 	gl.CompileShader(vshader)
// 	defer gl.DeleteShader(vshader)

// 	var vstatus int32
// 	gl.GetShaderiv(vshader, gl.COMPILE_STATUS, &vstatus)
// 	if vstatus == gl.FALSE {
// 		var logLength int32
// 		gl.GetShaderiv(vshader, gl.INFO_LOG_LENGTH, &logLength)

// 		logMsg := strings.Repeat("\x00", int(logLength+1))
// 		gl.GetShaderInfoLog(vshader, logLength, nil, gl.Str(logMsg))

// 		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, vShader)
// 		os.Exit(0)
// 	}

// 	// Create fragment shader
// 	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
// 	fsources, ffree := gl.Strs(fShader)
// 	gl.ShaderSource(fshader, 1, fsources, nil)
// 	ffree()
// 	gl.CompileShader(fshader)
// 	defer gl.DeleteShader(fshader)

// 	var fstatus int32
// 	gl.GetShaderiv(fshader, gl.COMPILE_STATUS, &fstatus)
// 	if fstatus == gl.FALSE {
// 		var logLength int32
// 		gl.GetShaderiv(fshader, gl.INFO_LOG_LENGTH, &logLength)

// 		logMsg := strings.Repeat("\x00", int(logLength+1))
// 		gl.GetShaderInfoLog(fshader, logLength, nil, gl.Str(logMsg))

// 		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, fShader)
// 		os.Exit(0)
// 	}

// 	// Create program
// 	var shaderProgram uint32
// 	shaderProgram = gl.CreateProgram()

// 	gl.AttachShader(shaderProgram, vshader)
// 	gl.AttachShader(shaderProgram, fshader)

// 	// Link program
// 	gl.LinkProgram(shaderProgram)

// 	var statisLink int32
// 	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &statisLink)
// 	if statisLink == gl.FALSE {
// 		var logLength int32
// 		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

// 		logMsg := strings.Repeat("\x00", int(logLength+1))
// 		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(logMsg))

// 		log.Printf("\n\n ### SHADER LINK ERROR ### \n%v\n\n", logMsg)
// 		os.Exit(0)
// 	}

// 	// Use this program for all upcoming render calls
// 	gl.UseProgram(shaderProgram)

// 	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

// 	return shaderProgram
// }

// func (r *OpenGL) LoadRenderer(renderer render.RendererRoutine) { // USE ENGINE VARIABLES TO SEND RENDER SYSTEM VARIABLES
// 	// Bind vertex array object. This must wrap around the mesh creation because it is how we are going to access it later when we draw
// 	var vertexArrayID uint32
// 	gl.GenVertexArrays(1, &vertexArrayID)
// 	gl.BindVertexArray(vertexArrayID)
// 	log.Printf("LoadRenderer > VAO > ID: %v", vertexArrayID)

// 	// Vertex buffer
// 	var vertexBuffer uint32
// 	gl.GenBuffers(1, &vertexBuffer)
// 	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
// 	gl.BufferData(gl.ARRAY_BUFFER, len(renderer.MeshVertices())*4, gl.Ptr(renderer.MeshVertices()), gl.STATIC_DRAW)
// 	log.Printf("LoadRenderer > VBO > Verticies Length: %v", len(renderer.MeshVertices()))

// 	// Element buffer
// 	var elementBuffer uint32
// 	gl.GenBuffers(1, &elementBuffer)
// 	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
// 	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(renderer.MeshIndicies())*4, gl.Ptr(renderer.MeshIndicies()), gl.STATIC_DRAW)
// 	log.Printf("LoadRenderer > EBO > Indicies Length: %v", len(renderer.MeshIndicies()))

// 	// Linking vertex attributes
// 	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
// 	gl.EnableVertexAttribArray(0)

// 	// Linking fragment attributes
// 	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
// 	gl.EnableVertexAttribArray(1)

// 	// Linking texture attributes
// 	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
// 	gl.EnableVertexAttribArray(2)

// 	if renderer.HasTexture() {
// 		// Load texture
// 		var texture uint32
// 		gl.GenTextures(1, &texture)
// 		gl.ActiveTexture(gl.TEXTURE0)
// 		gl.BindTexture(gl.TEXTURE_2D, texture)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
// 		gl.TexImage2D(
// 			gl.TEXTURE_2D,
// 			0,
// 			gl.RGBA,
// 			int32(renderer.GetTextureRGBA().Rect.Size().X),
// 			int32(renderer.GetTextureRGBA().Rect.Size().Y),
// 			0,
// 			gl.RGBA,
// 			gl.UNSIGNED_BYTE,
// 			gl.Ptr(renderer.GetTextureRGBA().Pix))
// 		log.Printf("LoadRenderer > Texture > ID: %v", texture)
// 		log.Printf("LoadRenderer > Texture > Width: %v", int32(renderer.GetTextureRGBA().Rect.Size().X))
// 		log.Printf("LoadRenderer > Texture > Height: %v", int32(renderer.GetTextureRGBA().Rect.Size().Y))
// 		log.Printf("LoadRenderer > Texture > Pix Length: %v", len(renderer.GetTextureRGBA().Pix))
// 		renderer.SetProperty("Texture", texture)

// 		color := renderer.GetColor()
// 		if color != nil {
// 			renderer.SetProperty("Color", color)
// 		}
// 	} else {
// 		color := renderer.GetColor()
// 		if color != nil {
// 			renderer.SetProperty("Color", color)
// 		}
// 	}

// 	renderer.SetProperty("VAO", vertexArrayID)

// 	// Unbind Vertex array object
// 	gl.BindVertexArray(0)
// }

// func (r *OpenGL) AddUISystem(game *engine.Engine) {
// 	// Create ui system
// 	sys_ui := &uigl.UIGL{Window: r.Window}
// 	game.AddSystem(engine.SystemUI, sys_ui)
// 	sys_ui.Init()
// 	r.uiSystem = sys_ui
// }

// func (r *OpenGL) Stop() {
// }

// func (r *OpenGL) GetWindow() *glfw.Window {
// 	return r.Window
// }

// func (r *OpenGL) GetCamera() *render.Camera {
// 	return r.Camera
// }
