// +build opengl
// +build glfw
// +build linux
// +build !windows
// +build !darwin

package opengl

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/opengl-glfw"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	"github.com/go-gl/gl/v4.5-core/gl"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"log"
	// "math"
	"os"
	"strings"
)

func init() {
	render.NewRenderSystem(&OpenGL{})
	log.Println("added opengl to engine2")
}

type OpenGL struct {
	render.Render
	window        *glfw32.Window
	screenWidth   float32
	screenHeight  float32
	shaderProgram uint32
	cameras       []*render.Camera
	transforms    []*std.Transform
	meshes        []*render.Mesh
	materials     []*render.Material
	settings      *engine.Settings
}

func (o *OpenGL) InitSystem() {
	o.settings = o.Engine().Settings()

	log.Println("OpenGL Prepare")
	o.screenWidth = o.settings.ResolutionX
	o.screenHeight = o.settings.ResolutionY

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.Viewport(0, 0, int32(o.settings.ResolutionX), int32(o.settings.ResolutionY))

	o.shaderProgram = o.NewShader(render.VSHADER, render.FSHADER)

	gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.FRONT_AND_BACK)

	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.3, 0.3, 0.3, 1)
}

func (o *OpenGL) NewIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		o.window = integrant.Window()
		log.Println("NewIntegrant(*glfw.GLFW)")
		break
	}
}

func (o *OpenGL) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *std.Transform:
		log.Println("NewComponent(*std.Transform)")
		o.RegisterTransform(component)
		break
	case *render.Mesh:
		log.Println("NewComponent(*render.Mesh)")
		o.RegisterMesh(component)
		break
	case *render.Material:
		log.Println("NewComponent(*render.Material)")
		o.RegisterMaterial(component)
		break
	case *render.Camera:
		log.Println("NewComponent(*render.Camera)")
		o.RegisterCamera(component)
		break
	}
}

func (o *OpenGL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&std.Transform{}, &render.Mesh{}, &render.Material{}, &render.Camera{}}
}

func (o *OpenGL) Update() {
	// log.Println(len(o.transforms), len(o.meshes), len(o.materials))

	if o.Active() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(o.shaderProgram)

		if len(o.cameras) <= 0 {
			log.Fatal("Your scene needs atleast one camera. Later versions of engine might allow no camera (for simulations)")
		}
		camera := o.cameras[0]

		ratio := float32(o.settings.ResolutionX) / float32(o.settings.ResolutionY)

		proj := mgl32.Ortho(0, float32(o.settings.ResolutionX)/ratio, 0, float32(o.settings.ResolutionY)/ratio, -1000.0, 1000.0)
		proj_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("projection\x00"))

		view := mgl32.LookAtV(mgl32.Vec3{camera.Eye().X - (o.settings.ResolutionY / 2), camera.Eye().Y - (o.settings.ResolutionX / 4), camera.Eye().Z}, mgl32.Vec3{camera.Center().X - (o.settings.ResolutionY / 2), camera.Center().Y - (o.settings.ResolutionX / 4), camera.Center().Z}, mgl32.Vec3{camera.Up().X, camera.Up().Y, camera.Up().Z})
		view_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("view\x00"))

		model_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("model\x00"))

		if len(o.meshes) != len(o.transforms) || len(o.meshes) != len(o.materials) {
			log.Println("Skew components")
			log.Fatalf("meshes: %v | transforms: %v | materials: %v", len(o.meshes), len(o.transforms), len(o.materials))
		}

		for idx, mesh := range o.meshes {
			// mesh := o.meshes[idx]
			// if mesh == nil {
			// 	continue
			// }
			gl.BindVertexArray(mesh.VAO())

			transform := o.transforms[idx]
			// this is a shortcut. should rather remove component from system registry
			if transform == nil {
				continue
			} else if !transform.Active() {
				// log.Fatal("here")
				continue
			}

			position := transform.Position()
			rotation := transform.Rotation()
			scale := transform.Scale()

			// model = model.Mul4(mgl32.Scale3D(1, 1, 1))
			model := mgl32.Ident4()
			model = model.Mul4(mgl32.Translate3D(position.X, position.Y, position.Z))
			model = model.Mul4(mgl32.Translate3D(scale.X/2, scale.Y/2, scale.Z/2))
			model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X)).Mat4())
			model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y)).Mat4())
			model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z)).Mat4())
			model = model.Mul4(mgl32.Translate3D(-scale.X/2, -scale.Y/2, -scale.Z/2))
			model = model.Mul4(mgl32.Scale3D(scale.X, scale.Y, scale.Z))
			// model = model.Mul4(mgl32.Translate3D(-(position.X / 2), -(position.Y / 2), (position.Z / 2)))
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
			material := o.materials[idx]
			color := material.Color()
			if color != nil {
				gl.Uniform4fv(gl.GetUniformLocation(o.shaderProgram, gl.Str("color\x00")), 1, &color.R)
			}

			texture := material.Texture()
			if texture != nil {
				gl.ActiveTexture(gl.TEXTURE0)
				gl.BindTexture(gl.TEXTURE_2D, texture.ID())
				gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("texture\x00")), 0)
				gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("hasTexture\x00")), 1)
			} else {
				gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("hasTexture\x00")), 0)
			}

			// 				color := renderer.GetProperty("Color")
			// 				if color != nil {
			// 					switch color := color.(type) {
			// 					case *render.Color:
			// 					}
			// 				}
			gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
			gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
			gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])
			gl.DrawElements(gl.TRIANGLES, 8, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
			// 			} else {
			// 				switch renderer := renderer.(type) {
			// 				case *render.MeshRenderer:
			// 					r.LoadRenderer(renderer)
			// 				}
			// 			}
			// 		}
			gl.BindVertexArray(0)
		}

		o.window.SwapBuffers()
		glfw32.PollEvents()
	} else {
		if o.window.ShouldClose() {
			glfw32.Terminate()
		}
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

func (o *OpenGL) RegisterTransform(transform *std.Transform) {
	o.transforms = append(o.transforms, transform)
}

func (o *OpenGL) RegisterCamera(camera *render.Camera) {
	o.cameras = append(o.cameras, camera)
}

func (o *OpenGL) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(o.transforms); i++ {
		transform := o.transforms[i]
		if transform.Entity().ID() == entity.ID() {
			copy(o.materials[i:], o.materials[i+1:])
			o.materials[len(o.materials)-1] = nil
			o.materials = o.materials[:len(o.materials)-1]

			copy(o.meshes[i:], o.meshes[i+1:])
			o.meshes[len(o.meshes)-1] = nil
			o.meshes = o.meshes[:len(o.meshes)-1]

			copy(o.transforms[i:], o.transforms[i+1:])
			o.transforms[len(o.transforms)-1] = nil
			o.transforms = o.transforms[:len(o.transforms)-1]
		}
	}
}

func (o *OpenGL) RegisterMaterial(material *render.Material) {
	texture := material.Texture()
	if texture != nil {
		// Load texture
		var tid uint32
		gl.GenTextures(1, &tid)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tid)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
		// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			texture.Width,
			texture.Height,
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(texture.RGBA()))
		texture.SetID(tid)
	}

	o.materials = append(o.materials, material)
}

func (o *OpenGL) RegisterMesh(mesh *render.Mesh) {
	var verticies []float32 = mesh.Vertices()
	var indicies []uint8 = mesh.Indicies()

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)
	// log.Printf("LoadRenderer > VAO > ID: %v", vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, gl.Ptr(verticies), gl.STATIC_DRAW)
	// log.Printf("LoadRenderer > VBO > Verticies Length: %v", len(verticies))

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)
	// log.Printf("LoadRenderer > EBO > Indicies Length: %v", len(indicies))

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
