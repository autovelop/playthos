// +build deploy opengl

package opengl

import (
	"github.com/autovelop/playthos"
	glfw "github.com/autovelop/playthos/glfw"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	"github.com/autovelop/playthos/platforms/linux"

	// for now we always pull in glfw if opengl is used until other window managers exist
	_ "github.com/autovelop/playthos/glfw"

	"github.com/go-gl/gl/all-core/gl"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"bytes"
	"image"
	"image/draw"
	_ "image/png"

	"fmt"
	"log"
)

func init() {
	render.NewRenderSystem(&OpenGL{})
	fmt.Println("> OpenGL Added")
}

type OpenGL struct {
	render.Render
	factory       *OpenGLFactory
	platform      *linux.Linux
	window        *glfw32.Window
	screenWidth   float32
	screenHeight  float32
	shaderProgram uint32
	cameras       []*render.Camera
	transforms    []*std.Transform
	meshes        []*OpenGLMesh
	materials     []*OpenGLMaterial
	settings      *engine.Settings
	majorVersion  int
	minorVersion  int
}

func (o *OpenGL) InitSystem() {
	o.SetActive(false)
	o.settings = o.Engine().Settings()

	o.screenWidth = o.settings.ResolutionX
	o.screenHeight = o.settings.ResolutionY

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.Viewport(0, 0, int32(o.settings.ResolutionX), int32(o.settings.ResolutionY))

	switch o.majorVersion {
	case 4:
		switch o.minorVersion {
		case 5:
			o.shaderProgram = o.factory.NewShader(render.VSHADER, render.FSHADER)
			break
		case 1:
			o.shaderProgram = o.factory.NewShader(render.VSHADER41, render.FSHADER41)
			break
		}
		break
	case 3:
		o.shaderProgram = o.factory.NewShader(render.VSHADER33, render.FSHADER33)
		break
	default:
		log.Fatalf("Playthos doesn't support OpenGL version older than v3.3")
		break

	}

	gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.FRONT_AND_BACK)

	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

}

func (o *OpenGL) Destroy() {
}

func (o *OpenGL) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		o.window = integrant.Window()
		o.majorVersion, o.minorVersion = integrant.OpenGLVersion()
		o.SetActive(true)
		fmt.Println("> OpenGL: Discovered GLFW")
		break
	case *OpenGLFactory:
		o.factory = integrant
		break
	case *linux.Linux:
		o.platform = integrant
		break
	}
}

func (o *OpenGL) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *std.Transform:
		// log.Println("AddComponent(*std.Transform)")
		o.RegisterTransform(component)
		break
	case *render.Mesh:
		// component = component.(*OpenGLMesh)
		o.RegisterMesh(component)
		break
	case *render.Material:
		// log.Println("AddComponent(*render.Material)")
		o.RegisterMaterial(component)
		break
	case *render.Camera:
		// log.Println("AddComponent(*render.Camera)")
		o.RegisterCamera(component)
		break
	}
}

func (o *OpenGL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&std.Transform{}, &render.Mesh{}, &render.Material{}, &render.Camera{}}
}

func (o *OpenGL) Draw() {
	if o.Active() {
		if len(o.cameras) <= 0 {
			o.createDefaultCamera()
		}
		if o.window.ShouldClose() {
			o.window.Destroy()
			defer glfw32.Terminate()
			// log.Fatal("here")
			// glfw32.Terminate()
		} else {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			gl.UseProgram(o.shaderProgram)

			if len(o.cameras) <= 0 {
				log.Fatal("Your scene needs atleast one camera. Later versions of engine might allow zero (for simulations) or more than one camera.")
			}
			camera := o.cameras[0]
			clearColor := camera.ClearColor()
			gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)

			proj := mgl32.Ortho(0, float32(o.settings.ResolutionX)/(*camera.Scale()), 0, float32(o.settings.ResolutionY)/(*camera.Scale()), -1000.0, 1000.0)
			proj_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("projection\x00"))

			view := mgl32.LookAtV(
				mgl32.Vec3{
					camera.Eye().X - ((o.settings.ResolutionX / 2) / (*camera.Scale())),
					camera.Eye().Y - ((o.settings.ResolutionY / 2) / (*camera.Scale())),
					camera.Eye().Z,
				},
				mgl32.Vec3{
					camera.Center().X - ((o.settings.ResolutionX / 2) / (*camera.Scale())),
					camera.Center().Y - ((o.settings.ResolutionY / 2) / (*camera.Scale())),
					camera.Center().Z,
				},
				mgl32.Vec3{
					camera.Up().X,
					camera.Up().Y,
					camera.Up().Z,
				})

			view_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("view\x00"))

			model_uni := gl.GetUniformLocation(o.shaderProgram, gl.Str("model\x00"))

			// if len(o.meshes) != len(o.transforms) || len(o.meshes) != len(o.materials) {
			// 	log.Println("Skew components")
			// 	log.Fatalf("meshes: %v | transforms: %v | materials: %v", len(o.meshes), len(o.transforms), len(o.materials))
			// }

			for idx, mesh := range o.meshes {
				if mesh == nil {
					continue
				}
				gl.BindVertexArray(mesh.VAO())

				transform := o.transforms[idx]

				// this is a shortcut. should rather remove component from system registry
				if transform == nil {
					continue
				} else if !transform.Active() {
					continue
				}

				position := transform.Position()
				rotation := transform.Rotation()
				scale := transform.Scale()

				model := mgl32.Ident4()
				model = model.Mul4(mgl32.Translate3D(position.X, position.Y, position.Z))
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z / 1)).Mat4())
				model = model.Mul4(mgl32.Translate3D(-scale.X/2, -scale.Y/2, -scale.Z/2))
				model = model.Mul4(mgl32.Scale3D(scale.X, scale.Y, scale.Z))

				material := o.materials[idx]
				if material == nil {
					continue
				} else if !material.Active() {
					continue
				}

				color := material.Color()
				if color != nil {
					gl.Uniform4fv(gl.GetUniformLocation(o.shaderProgram, gl.Str("color\x00")), 1, &color.R)
				}

				texture := material.Texture()
				if texture != nil {
					gl.Uniform2fv(gl.GetUniformLocation(o.shaderProgram, gl.Str("spriteScaler\x00")), 1, &texture.SizeN().X)
					gl.Uniform2fv(gl.GetUniformLocation(o.shaderProgram, gl.Str("spriteOffset\x00")), 1, &texture.Offset().X)
					gl.ActiveTexture(gl.TEXTURE0)
					gl.BindTexture(gl.TEXTURE_2D, texture.ID())
					gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("texture\x00")), 0)
					gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("hasTexture\x00")), 1)
				} else {
					gl.Uniform1i(gl.GetUniformLocation(o.shaderProgram, gl.Str("hasTexture\x00")), 0)
				}
				gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
				gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
				gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])
				gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
				gl.BindVertexArray(0)
			}

			o.window.SwapBuffers()
			glfw32.PollEvents()
		}
	}
}

func (o *OpenGL) createDefaultCamera() {
	camera_transform := std.NewTransform()
	camera_transform.Set(
		&std.Vector3{0, 0, 3}, // POSITION
		&std.Vector3{0, 0, 0}, // CENTER
		&std.Vector3{0, 1, 0}, // UP
	)
	camera := render.NewCamera()
	cameraSize := float32(1.0)
	camera.Set(&cameraSize, &std.Color{0, 0, 0, 0})
	camera.SetTransform(camera_transform)
	o.cameras = append(o.cameras, camera)
}

func (o *OpenGL) RegisterTransform(transform *std.Transform) {
	o.transforms = append(o.transforms, transform)
}

func (o *OpenGL) RegisterCamera(camera *render.Camera) {
	clearColor := camera.ClearColor()
	gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
	camera.SetWindow(o.settings.ResolutionX, o.settings.ResolutionY)
	// gl.Viewport(0, 0, int32(), int32(o.settings.ResolutionY))
	o.cameras = append(o.cameras, camera)
	o.materials = append(o.materials, nil)
	o.meshes = append(o.meshes, nil)
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
	texture := material.BaseTexture()
	openGLMaterial := &OpenGLMaterial{Material: material}
	if texture != nil {
		openGLMaterial.OverrideTexture(func(t render.Textureable) {
			raw := o.platform.Asset(t.Path())

			img, _, err := image.Decode(bytes.NewReader(raw))
			if err != nil {
				log.Println(err)
				return
			}

			rgba := image.NewRGBA(img.Bounds())
			if rgba.Stride != rgba.Rect.Size().X*4 {
				log.Println("rgba stride error")
				return
			}

			t.SetHeight(int32(rgba.Rect.Size().Y))
			t.SetWidth(int32(rgba.Rect.Size().X))

			pix := rgba.Pix
			draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

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
				t.Width(),
				t.Height(),
				0,
				gl.RGBA,
				gl.UNSIGNED_BYTE,
				gl.Ptr(pix))

			openGLMaterial.texture = &OpenGLTexture{t.(*render.Texture), tid}
		})
	}
	o.materials = append(o.materials, openGLMaterial)
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
	// var openGLMesh *OpenGLMesh
	// openGLMesh = mesh.(*OpenGLMesh)

	var openGLMesh *OpenGLMesh
	openGLMesh = &OpenGLMesh{m: mesh}

	openGLMesh.SetVAO(vertexArrayID)

	// Unbind Vertex array object
	gl.BindVertexArray(0)
	o.meshes = append(o.meshes, openGLMesh)
}
