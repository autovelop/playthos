// +build deploy webgl

package webgl

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"

	"errors"
	"fmt"
	"github.com/autovelop/playthos/platforms/web"
	"log"
)

// BUG(F): Bindings to WebGL has been removed so import from "github.com/gopherjs/webgl" and implement below functions
// func (c *Context) Str(s string) string {
// 	return string(s)
// }
// func (c *Context) BindVertexArray(vao *js.Object) {
// 	c.Call("bindVertexArray", vao)
// }
// func (c *Context) CreateVertexArray() *js.Object {
// 	return c.Call("createVertexArray")
// }
func glStr(s string) string {
	return string(s)
}
func (w *WebGL) BindVertexArray(vao *js.Object) {
	w.gl.Call("bindVertexArray", vao)
}
func (w *WebGL) CreateVertexArray() *js.Object {
	return w.gl.Call("createVertexArray")
}

func NewContext(canvas *js.Object, ca *webgl.ContextAttributes) (*webgl.Context, error) {
	if js.Global.Get("WebGLRenderingContext") == js.Undefined {
		return nil, errors.New("Your browser doesn't appear to support webgl.")
	}

	if ca == nil {
		ca = webgl.DefaultAttributes()
	}

	attrs := map[string]bool{
		"alpha":                 ca.Alpha,
		"depth":                 ca.Depth,
		"stencil":               ca.Stencil,
		"antialias":             ca.Antialias,
		"premultipliedAlpha":    ca.PremultipliedAlpha,
		"preserveDrawingBuffer": ca.PreserveDrawingBuffer,
	}
	gl := canvas.Call("getContext", "webgl2", attrs)
	if gl == nil {
		gl = canvas.Call("getContext", "webgl", attrs)
		if gl == nil {
			return nil, errors.New("Creating a webgl context has failed.")
		}
	}
	ctx := new(webgl.Context)
	ctx.Object = gl
	return ctx, nil
}

func init() {
	render.NewRenderSystem(&WebGL{})
	fmt.Println("> WebGL Added")
}

// WebGL uses gopherjs render graphics on web browsers
type WebGL struct {
	render.Render
	platform      *web.Web
	factory       *WebGLFactory
	gl            *webgl.Context
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

// InitSystem called when the system plugs into the engine
func (w *WebGL) InitSystem() {
	w.SetActive(false)
	w.settings = w.Engine().Settings()

	w.screenWidth = w.settings.ResolutionX
	w.screenHeight = w.settings.ResolutionY
	fmt.Println("> WebGL Init System")
	// This should be under the web platform and now webgl

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

	attrs := webgl.DefaultAttributes()
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

	w.gl.Enable(w.gl.DEPTH_TEST)
	w.gl.Enable(w.gl.BLEND)
	w.gl.BlendFunc(w.gl.SRC_ALPHA, w.gl.ONE_MINUS_SRC_ALPHA)

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

// Destroy called when engine is gracefully shutting down
func (w *WebGL) Destroy() {
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (w *WebGL) AddIntegrant(integrant engine.IntegrantRoutine) {
	// fmt.Printf("asdasd: %+v", integrant)
	switch integrant := integrant.(type) {
	case *WebGLFactory:
		w.factory = integrant
		break
	case *web.Web:
		w.platform = integrant
		break
	}
}

// AddComponent unorphans a component by adding it to this system
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

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (w *WebGL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&std.Transform{}, &render.Mesh{}, &render.Material{}, &render.Camera{}}
}

// requestAnimationFrame renders all entity components into scene. Called by browser
func (w *WebGL) requestAnimationFrame(float32) {
	gl := w.gl
	if w.Active() {
		if len(w.cameras) <= 0 {
			w.createDefaultCamera()
		} else {

			if len(w.cameras) <= 0 {
				log.Fatal("Your scene needs atleast one camera. Later versions of engine might allow zero (for simulations) or more than one camera.")
			}
			camera := w.cameras[0]
			clearColor := camera.ClearColor()
			gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			camSize := float32(0.2)
			camX := float32(w.settings.ResolutionX) * camSize
			camY := float32(w.settings.ResolutionY) * camSize
			proj := mgl32.Ortho(-camX/20, camX/20, -camY/20, camY/20, 0.1, 100.0)

			view := mgl32.LookAtV(
				mgl32.Vec3{
					camera.Eye().X,
					camera.Eye().Y,
					-camera.Eye().Z,
				},
				mgl32.Vec3{
					camera.Center().X,
					camera.Center().Y,
					camera.Center().Z,
				},
				mgl32.Vec3{
					camera.Up().X,
					camera.Up().Y,
					camera.Up().Z,
				})

			gl.UseProgram(w.shaderProgram)

			view_uni := gl.GetUniformLocation(w.shaderProgram, glStr("uViewMatrix"))
			gl.UniformMatrix4fv(view_uni, false, view[:])
			proj_uni := gl.GetUniformLocation(w.shaderProgram, string("uProjMatrix"))
			gl.UniformMatrix4fv(proj_uni, false, proj[:])

			for idx, mesh := range w.meshes {
				// mesh := o.meshes[idx]
				if mesh == nil {
					continue
				}
				w.BindVertexArray(mesh.VAO())

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
				model = model.Mul4(mgl32.Translate3D(position.X, position.Y, -position.Z))
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z / 1)).Mat4())
				model = model.Mul4(mgl32.Scale3D(scale.X, scale.Y, scale.Z))

				if idx >= len(w.materials) {
					w.BindVertexArray(nil)
					continue
				}
				material := w.materials[idx]
				if material == nil {
					w.BindVertexArray(nil)
					continue
				} else if !material.Active() {
					w.BindVertexArray(nil)
					continue
				}

				color := material.Color()
				if color != nil {
					gl.Uniform4f(gl.GetUniformLocation(w.shaderProgram, glStr("uColor")), color.R, color.G, color.B, color.A)
				}
				texture := material.Texture()
				if texture != nil {
					gl.ActiveTexture(gl.TEXTURE0)
					gl.BindTexture(gl.TEXTURE_2D, texture.ID())
					gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, glStr("uTexture")), 0)
					gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, glStr("uTextured")), 1)
					// } else {
					// 	gl.Uniform1i(gl.GetUniformLocation(w.shaderProgram, glStr("hasTexture")), 0)
				}

				model_uni := gl.GetUniformLocation(w.shaderProgram, glStr("uModelMatrix"))
				gl.UniformMatrix4fv(model_uni, false, model[:])

				gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)
				w.BindVertexArray(nil)
			}
		}
		js.Global.Call("requestAnimationFrame", w.requestAnimationFrame)
	}
}

// createDefaultCamera creates a camera if not set explicitly
func (w *WebGL) createDefaultCamera() {
	camera_transform := std.NewTransform()
	camera_transform.Set(
		&std.Vector3{0, 0, -5}, // POSITION
		&std.Vector3{0, 0, 0},  // CENTER
		&std.Vector3{0, 1, 0},  // UP
	)
	camera := render.NewCamera()
	cameraSize := float32(40.0)
	camera.Set(&cameraSize, &std.Color{0, 0, 0, 1})
	camera.SetTransform(camera_transform)
	w.cameras = append(w.cameras, camera)
}

// RegisterTransform tells webgl about a given transform component
func (w *WebGL) RegisterTransform(transform *std.Transform) {
	if w.platform == nil {
		return
	}
	w.transforms = append(w.transforms, transform)
}

// RegisterCamera tells webgl about a given camera component
func (w *WebGL) RegisterCamera(camera *render.Camera) {
	if w.platform == nil {
		return
	}
	clearColor := camera.ClearColor()
	w.gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
	// camera.SetWindow(w.settings.ResolutionX, w.settings.ResolutionY)
	// gl.Viewport(0, 0, int32(), int32(w.settings.ResolutionY))
	w.cameras = append(w.cameras, camera)
	w.materials = append(w.materials, nil)
	w.meshes = append(w.meshes, nil)
}

// DeleteEntity removes all entity's compoents from this system
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

// RegisterMaterial tells webgl about a given material component
func (w *WebGL) RegisterMaterial(material *render.Material) {
	if w.platform == nil {
		return
	}
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

// RegisterMesh tells webgl about a given mesh component
func (w *WebGL) RegisterMesh(mesh *render.Mesh) {
	if w.platform == nil {
		return
	}
	gl := w.gl
	var verticies []float32 = mesh.Vertices()
	var indicies []uint8 = mesh.Indicies()

	vertexArrayID := w.CreateVertexArray()
	w.BindVertexArray(vertexArrayID)
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
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 4*4, 0)
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 4*4, 2*4)
	gl.EnableVertexAttribArray(1)

	// Linking texture attributes
	// gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, 6*4)
	// gl.EnableVertexAttribArray(2)

	webGLMesh := &WebGLMesh{m: mesh}

	webGLMesh.SetVAO(vertexArrayID)

	// Unbind Vertex array object
	w.BindVertexArray(nil)
	w.meshes = append(w.meshes, webGLMesh)
}
