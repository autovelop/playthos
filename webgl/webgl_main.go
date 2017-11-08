// +build deploy webgl

package webgl

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"

	"github.com/autovelop/playthos/platforms/web"
	// "github.com/gopherjs/webgl"
	// "github.com/go-gl/gl/all-core/gl"
	// "time"
	// glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	// "github.com/go-gl/mathgl/mgl32"
	"log"
	// "math"
	"fmt"
	// "os"
	// "strings"
)

func init() {
	render.NewRenderSystem(&WebGL{})
	fmt.Println("> WebGL Added")
}

type WebGL struct {
	render.Render
	platform      *web.Web
	factory       *WebGLFactory
	gl            *Context
	screenWidth   float32
	screenHeight  float32
	shaderProgram *js.Object
	cameras       []*render.Camera
	transforms    []*std.Transform
	meshes        []*WebGLMesh
	materials     []*WebGLMaterial
	settings      *engine.Settings
	majorVersion  int
	minorVersion  int
}

func (w *WebGL) InitSystem() {
	w.SetActive(false)
	w.settings = w.Engine().Settings()

	w.screenWidth = w.settings.ResolutionX
	w.screenHeight = w.settings.ResolutionY
	fmt.Println("> WebGL Init System")

	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	canvas.Set("id", "canvas")
	canvas.Set("width", w.screenWidth*2)
	canvas.Set("height", w.screenHeight*2)
	canvas.Get("style").Set("width", "100vw")
	canvas.Get("style").Set("max-width", fmt.Sprintf("calc(100vh * (%d / %d))", int(w.screenWidth), int(w.screenHeight)))
	// canvas.Get("style").Set("height", fmt.Sprintf("%dpx", int(w.screenHeight)))
	canvas.Get("style").Set("max-height", "100vh")
	document.Get("body").Call("appendChild", canvas)

	attrs := DefaultAttributes()
	attrs.Alpha = false

	var err error
	w.gl, err = NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	w.shaderProgram = w.factory.NewShader(w.gl, render.VSHADERWEB, render.FSHADERWEB)

	w.platform.Loaded = func() {
		js.Global.Call("requestAnimationFrame", w.requestAnimationFrame)
	}

	// gl.ClearColor(0.8, 0.3, 0.01, 1)
	// gl.Clear(gl.COLOR_BUFFER_BIT)
	// os.Exit(0)

	// 	document := js.Global.Get("document")
	// 	canvas := document.Call("createElement", "canvas")
	// 	document.Get("body").Call("appendChild", canvas)

	// 	attrs := webgl.DefaultAttributes()
	// 	attrs.Alpha = false

	// 	gl, err := webgl.NewContext(canvas, attrs)
	// 	if err != nil {
	// 		js.Global.Call("alert", "Error: "+err.Error())
	// 	}

	// 	gl.ClearColor(0.8, 0.3, 0.01, 1)
	// 	gl.Clear(gl.COLOR_BUFFER_BIT)
	// 	os.Exit(0)
}

func (w *WebGL) Destroy() {
}

func (w *WebGL) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *WebGLFactory:
		w.factory = integrant
		break
	case *web.Web:
		w.platform = integrant
		break
	}
}

func (w *WebGL) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *std.Transform:
		w.RegisterTransform(component)
		break
	case *render.Mesh:
		w.RegisterMesh(component)
		break
	case *render.Material:
		w.RegisterMaterial(component)
		break
	case *render.Camera:
		w.RegisterCamera(component)
		break
	}
}

func (w *WebGL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&std.Transform{}, &render.Mesh{}, &render.Material{}, &render.Camera{}}
}

func (w *WebGL) requestAnimationFrame(float32) {
	gl := w.gl
	if w.Active() {
		if len(w.cameras) <= 0 {
			w.createDefaultCamera()
		} else {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			gl.UseProgram(w.shaderProgram)

			if len(w.cameras) <= 0 {
				log.Fatal("Your scene needs atleast one camera. Later versions of engine might allow zero (for simulations) or more than one camera.")
			}
			camera := w.cameras[0]
			clearColor := camera.ClearColor()
			gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)

			proj := mgl32.Ortho(0, float32(w.settings.ResolutionX)/(*camera.Scale()), 0, float32(w.settings.ResolutionY)/(*camera.Scale()), -1000.0, 1000.0)
			proj_uni := gl.GetUniformLocation(w.shaderProgram, string("projection"))

			view := mgl32.LookAtV(
				mgl32.Vec3{
					camera.Eye().X - ((w.settings.ResolutionX / 2) / (*camera.Scale())),
					camera.Eye().Y - ((w.settings.ResolutionY / 2) / (*camera.Scale())),
					camera.Eye().Z,
				},
				mgl32.Vec3{
					camera.Center().X - ((w.settings.ResolutionX / 2) / (*camera.Scale())),
					camera.Center().Y - ((w.settings.ResolutionY / 2) / (*camera.Scale())),
					camera.Center().Z,
				},
				mgl32.Vec3{
					camera.Up().X,
					camera.Up().Y,
					camera.Up().Z,
				})
			// float32(math.Cos(float64(mgl32.DegToRad(camera.Direction().X))) * math.Cos(float64(mgl32.DegToRad(camera.Direction().Y)))),
			// float32(math.Sin(float64(mgl32.DegToRad(camera.Direction().Y)))),
			// float32(math.Sin(float64(mgl32.DegToRad(camera.Direction().X))) * math.Cos(float64(mgl32.DegToRad(camera.Direction().Y)))),

			// front.x = cos(glm::radians(yaw)) * cos(glm::radians(pitch));
			// front.y = sin(glm::radians(pitch));
			// front.z = sin(glm::radians(yaw)) * cos(glm::radians(pitch));

			view_uni := gl.GetUniformLocation(w.shaderProgram, gl.Str("view"))

			model_uni := gl.GetUniformLocation(w.shaderProgram, gl.Str("model"))

			// if len(o.meshes) != len(o.transforms) || len(o.meshes) != len(o.materials) {
			// 	log.Println("Skew components")
			// 	log.Fatalf("meshes: %v | transforms: %v | materials: %v", len(o.meshes), len(o.transforms), len(o.materials))
			// }

			for idx, mesh := range w.meshes {
				// mesh := o.meshes[idx]
				if mesh == nil {
					continue
				}
				gl.BindVertexArray(mesh.VAO())

				transform := w.transforms[idx]
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

				model := mgl32.Ident4()
				model = model.Mul4(mgl32.Translate3D(position.X, position.Y, position.Z))
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z / 1)).Mat4())
				model = model.Mul4(mgl32.Translate3D(-scale.X/2, -scale.Y/2, -scale.Z/2))
				model = model.Mul4(mgl32.Scale3D(scale.X, scale.Y, scale.Z))

				material := w.materials[idx]
				if material == nil {
					continue
				} else if !material.Active() {
					continue
				}

				color := material.Color()
				if color != nil {
					gl.Uniform4f(gl.GetUniformLocation(w.shaderProgram, gl.Str("color")), color.R, color.G, color.B, color.A)
				}
				texture := material.Texture()
				if texture != nil {
					gl.Uniform2f(gl.GetUniformLocation(w.shaderProgram, gl.Str("spriteScaler")), texture.SizeN().X, texture.SizeN().Y)
					gl.Uniform2f(gl.GetUniformLocation(w.shaderProgram, gl.Str("spriteOffset")), texture.Offset().X, texture.Offset().Y)
					gl.ActiveTexture(gl.TEXTURE0)
					gl.BindTexture(gl.TEXTURE_2D, texture.ID())
					gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, gl.Str("texture")), 0)
					gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, gl.Str("hasTexture")), 1)
				} else {
					gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, gl.Str("hasTexture")), 0)
				}

				gl.UniformMatrix4fv(model_uni, false, model[:])
				gl.UniformMatrix4fv(view_uni, false, view[:])
				gl.UniformMatrix4fv(proj_uni, false, proj[:])
				gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)
				gl.BindVertexArray(nil)
			}
		}
		js.Global.Call("requestAnimationFrame", w.requestAnimationFrame)
	}
}

func (w *WebGL) createDefaultCamera() {
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
	w.cameras = append(w.cameras, camera)
}

func (w *WebGL) RegisterTransform(transform *std.Transform) {
	w.transforms = append(w.transforms, transform)
}

func (w *WebGL) RegisterCamera(camera *render.Camera) {
	clearColor := camera.ClearColor()
	w.gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
	// camera.SetWindow(w.settings.ResolutionX, w.settings.ResolutionY)
	// gl.Viewport(0, 0, int32(), int32(w.settings.ResolutionY))
	w.cameras = append(w.cameras, camera)
	w.materials = append(w.materials, nil)
	w.meshes = append(w.meshes, nil)
}

func (w *WebGL) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(w.transforms); i++ {
		transform := w.transforms[i]
		if transform.Entity().ID() == entity.ID() {
			copy(w.materials[i:], w.materials[i+1:])
			w.materials[len(w.materials)-1] = nil
			w.materials = w.materials[:len(w.materials)-1]

			copy(w.meshes[i:], w.meshes[i+1:])
			w.meshes[len(w.meshes)-1] = nil
			w.meshes = w.meshes[:len(w.meshes)-1]

			copy(w.transforms[i:], w.transforms[i+1:])
			w.transforms[len(w.transforms)-1] = nil
			w.transforms = w.transforms[:len(w.transforms)-1]
		}
	}
}

func (w *WebGL) RegisterMaterial(material *render.Material) {
	gl := w.gl
	texture := material.BaseTexture()
	webGLMaterial := &WebGLMaterial{Material: material}
	if texture != nil {
		webGLMaterial.OverrideTexture(func(t render.Textureable) {
			raw := w.platform.Asset(t.Path())
			if raw != nil {

				t.SetHeight(int32(raw.Get("height").Int()))
				t.SetWidth(int32(raw.Get("width").Int()))

				tid := gl.CreateTexture()
				gl.ActiveTexture(gl.TEXTURE0)
				gl.BindTexture(gl.TEXTURE_2D, tid)
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
				gl.TexImage2D(
					gl.TEXTURE_2D,
					0,
					gl.RGBA,
					gl.RGBA,
					gl.UNSIGNED_BYTE,
					raw)
				webGLMaterial.texture = &WebGLTexture{t.(*render.Texture), tid}
			}
		})
	}
	w.materials = append(w.materials, webGLMaterial)
}

func (w *WebGL) RegisterMesh(mesh *render.Mesh) {
	gl := w.gl
	var verticies []float32 = mesh.Vertices()
	var indicies []uint8 = mesh.Indicies()

	vertexArrayID := gl.CreateVertexArray()
	gl.BindVertexArray(vertexArrayID)
	// log.Printf("LoadRenderer > VAO > ID: %v", vertexArrayID)

	// Vertex buffer
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, verticies, gl.STATIC_DRAW)
	// log.Printf("LoadRenderer > VBO > Verticies Length: %v", len(verticies))

	// Element buffer
	elementBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indicies, gl.STATIC_DRAW)
	// log.Printf("LoadRenderer > EBO > Indicies Length: %v", len(indicies))

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, 0)
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, 3*4)
	gl.EnableVertexAttribArray(1)

	// Linking texture attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, 6*4)
	gl.EnableVertexAttribArray(2)

	webGLMesh := &WebGLMesh{m: mesh}

	webGLMesh.SetVAO(vertexArrayID)

	// Unbind Vertex array object
	gl.BindVertexArray(nil)
	w.meshes = append(w.meshes, webGLMesh)
}
