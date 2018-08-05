// +build deploy opengl

package opengl

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	// "github.com/autovelop/playthos/platforms/linux"

	// for now we always pull in glfw if opengl is used until other window managers exist
	glfw "github.com/autovelop/playthos/glfw"

	"github.com/go-gl/gl/all-core/gl"
	glfw32 "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"os"
	"strings"
	"time"
	"unsafe"

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

// OpenGL uses GLFW window to render graphics on desktop devices
type OpenGL struct {
	render.Render
	platform          engine.Desktoper
	window            *glfw32.Window
	screenWidth       float32
	screenHeight      float32
	shader            uint32
	opaqueShader      uint32
	opaqueFBO         uint32
	opaqueTex0        uint32
	sharedDepthBuffer uint32
	transparentShader uint32
	transparentFBO    uint32
	transparentTex0   uint32
	transparentTex1   uint32
	compositeShader   uint32
	framebufferShader uint32
	cameras           []*render.Camera
	transforms        []*std.Transform
	meshes            []*OpenGLMesh
	materials         []*OpenGLMaterial
	settings          *engine.Settings
	majorVersion      int
	minorVersion      int
	fbo               uint32
	rbo               uint32
	cbo               uint32
	quadVAO           uint32
}

// InitSystem called when the system plugs into the engine
func (o *OpenGL) InitSystem() {
	o.SetActive(false)
	o.settings = o.Engine().Settings()

	o.screenWidth = o.settings.ResolutionX
	o.screenHeight = o.settings.ResolutionY

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	switch o.majorVersion {
	case 4:
		switch o.minorVersion {
		case 5:
			o.shader = newShader(render.VSHADER, render.FSHADER)
			break
		case 1:
			// o.shaderProgram = newShader(render.VSHADER41, render.FSHADER41)
			break
		}
		break
	case 3:
		// o.shaderProgram = newShader(render.VSHADER33, render.FSHADER33)
		break
	default:
		log.Fatalf("Playthos doesn't support OpenGL version older than v3.3")
		break

	}

	gl.Viewport(0, 0, int32(o.settings.ResolutionX), int32(o.settings.ResolutionY))

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	o.registerQuad()
}

// Destroy called when engine is gracefully shutting down
func (o *OpenGL) Destroy() {
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (o *OpenGL) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *glfw.GLFW:
		o.window = integrant.Window()
		o.majorVersion, o.minorVersion = integrant.OpenGLVersion()
		o.SetActive(true)
		fmt.Println("> OpenGL: Discovered GLFW")
		break
	case engine.Desktoper:
		o.platform = integrant
		break
	}
}

// AddComponent unorphans a component by adding it to this system
func (o *OpenGL) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *std.Transform:
		o.RegisterTransform(component)
		break
	case *render.Mesh:
		o.RegisterMesh(component)
		break
	case *render.Material:
		o.RegisterMaterial(component)
		break
	case *render.Camera:
		o.RegisterCamera(component)
		break
	}
}

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (o *OpenGL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&std.Transform{}, &render.Mesh{}, &render.Material{}, &render.Camera{}}
}

// Draw renders all entity components into scene
func (o *OpenGL) Draw() {
	if o.Active() {
		if len(o.cameras) <= 0 {
			o.createDefaultCamera()
		}
		if o.window.ShouldClose() {
			o.window.Destroy()
			defer glfw32.Terminate()
		} else {
			if len(o.cameras) <= 0 {
				log.Fatal("Your scene needs atleast one camera. Later versions of engine might allow zero (for simulations) or more than one camera.")
			}
			camera := o.cameras[0]
			clearColor := camera.ClearColor()
			gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			camSize := float32(0.2)
			camX := float32(o.settings.ResolutionX) * camSize
			camY := float32(o.settings.ResolutionY) * camSize
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

			if gl.GetError() != 0 {
				log.Fatal("OpenGL error")
			}

			gl.UseProgram(o.shader)
			if gl.GetError() != 0 {
				log.Fatal("OpenGL error")
			}
			gl.UniformMatrix4fv(uniformLocation(o.shader, "uViewMatrix\x00"), 1, false, &view[0])
			if gl.GetError() != 0 {
				log.Fatal("OpenGL error")
			}
			gl.UniformMatrix4fv(uniformLocation(o.shader, "uProjMatrix\x00"), 1, false, &proj[0])
			if gl.GetError() != 0 {
				log.Fatal("OpenGL error")
			}
			for idx, mesh := range o.meshes {
				if mesh == nil {
					continue
				}
				gl.BindVertexArray(mesh.VAO())
				if gl.GetError() != 0 {
					log.Fatal("OpenGL error")
				}

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
				model = model.Mul4(mgl32.Translate3D(position.X, position.Y, -position.Z))
				model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rotation.X / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rotation.Y / 1)).Mat4())
				model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rotation.Z / 1)).Mat4())
				model = model.Mul4(mgl32.Scale3D(scale.X, scale.Y, scale.Z))

				if idx >= len(o.materials) {
					gl.BindVertexArray(0)
					continue
				}
				material := o.materials[idx]
				if material == nil {
					gl.BindVertexArray(0)
					continue
				} else if !material.Active() {
					gl.BindVertexArray(0)
					continue
				}

				color := material.Color()
				if color != nil {
					gl.Uniform4fv(uniformLocation(o.shader, "uColor\x00"), 1, &color.R)
					if gl.GetError() != 0 {
						log.Fatal("OpenGL error")
					}
				}

				texture := material.Texture()
				if texture != nil {
					gl.ActiveTexture(gl.TEXTURE0)
					gl.BindTexture(gl.TEXTURE_2D, texture.ID())
					gl.Uniform1i(uniformLocation(o.shader, "uTexture\x00"), 0)
					gl.Uniform1i(uniformLocation(o.shader, "uTextured\x00"), 1)
					if gl.GetError() != 0 {
						log.Fatal("OpenGL error")
					}
				}

				gl.UniformMatrix4fv(uniformLocation(o.shader, "uModelMatrix\x00"), 1, false, &model[0])
				if gl.GetError() != 0 {
					log.Fatal("OpenGL error")
				}

				gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
				if gl.GetError() != 0 {
					log.Fatal("OpenGL error")
				}
				gl.BindVertexArray(0)
			}
			gl.UseProgram(0)

			o.window.SwapBuffers()
			glfw32.PollEvents()
		}
	}
}

// createDefaultCamera creates a camera if not set explicitly
func (o *OpenGL) createDefaultCamera() {
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
	o.cameras = append(o.cameras, camera)
}

// RegisterTransform tells opengl about a given transform component
func (o *OpenGL) RegisterTransform(transform *std.Transform) {
	if o.platform == nil {
		return
	}
	o.transforms = append(o.transforms, transform)
}

// RegisterCamera tells opengl about a given camera component
func (o *OpenGL) RegisterCamera(camera *render.Camera) {
	if o.platform == nil {
		return
	}
	clearColor := camera.ClearColor()
	gl.ClearColor(clearColor.R, clearColor.G, clearColor.B, clearColor.A)
	camera.SetWindow(o.settings.ResolutionX, o.settings.ResolutionY)
	o.cameras = append(o.cameras, camera)
	o.materials = append(o.materials, nil)
	o.meshes = append(o.meshes, nil)
}

// DeleteEntity removes all entity's compoents from this system
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

// RegisterMaterial tells opengl about a given material component
func (o *OpenGL) RegisterMaterial(material *render.Material) {
	if o.platform == nil {
		return
	}
	texture := material.BaseTexture()
	openGLMaterial := NewOpenGLMaterial(material)
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

			tid := newTexture(gl.TEXTURE_2D, gl.RGBA, t.Width(), t.Height(), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pix))

			openGLMaterial.texture = &OpenGLTexture{t.(*render.Texture), tid}
		})
	}
	o.materials = append(o.materials, openGLMaterial)
}

// RegisterMesh tells opengl about a given mesh component
func (o *OpenGL) RegisterMesh(mesh *render.Mesh) {
	if o.platform == nil {
		return
	}
	var vertices []float32 = mesh.Vertices()
	var indicies []uint8 = mesh.Indicies()

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking texture attributes
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(2*4))
	gl.EnableVertexAttribArray(1)

	var openGLMesh *OpenGLMesh
	openGLMesh = &OpenGLMesh{m: mesh}

	openGLMesh.SetVAO(vertexArrayID)

	// Unbind Vertex array object
	gl.BindVertexArray(0)

	o.meshes = append(o.meshes, openGLMesh)
}

// registerQuad has to exist in every rendering system to allow for multipass rendering
func (o *OpenGL) registerQuad() {
	vertices := []float32{
		1.0, 1.0, 1.0, 0.0,
		1.0, -1.0, 1.0, 1.0,
		-1.0, -1.0, 0.0, 1.0,
		-1.0, 1.0, 0.0, 0.0,
	}
	indicies := []uint8{
		0, 1, 3,
		1, 2, 3,
	}

	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking texture attributes
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(2*4))
	gl.EnableVertexAttribArray(1)

	o.quadVAO = vertexArrayID

	// Unbind Vertex array object
	gl.BindVertexArray(0)
}

func (o *OpenGL) renderFramebuffer(fbo uint32, tid uint32) {
	// FBO DEBUGGER
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.UseProgram(o.framebufferShader)

	if gl.GetError() != 0 {
		log.Fatal("OpenGL error")
	}
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_RECTANGLE, tid)
	gl.Uniform1i(uniformLocation(o.framebufferShader, "Texture\x00"), int32(0))
	if gl.GetError() != 0 {
		log.Fatal("OpenGL error")
	}
	gl.BindVertexArray(o.quadVAO)
	if gl.GetError() != 0 {
		log.Fatal("OpenGL error")
	}
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
	if gl.GetError() != 0 {
		log.Fatal("OpenGL error")
	}
	gl.BindVertexArray(0)
	gl.UseProgram(0)
	o.window.SwapBuffers()
	time.Sleep(100 * time.Second)
}

// newShader creates a new vertex and fragment shader
func newShader(vs string, fs string) uint32 {
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("> OpenGL: Profile = %v\n", version)

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

	return shaderProgram
}

// newTexture creates a new texture
// (target, internal format, width, height, format)
func newTexture(ta uint32, i int32, w int32, h int32, f uint32, ty uint32, p unsafe.Pointer) uint32 {
	var id uint32
	gl.GenTextures(1, &id)
	gl.BindTexture(ta, id)
	gl.TexParameteri(ta, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(ta, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(ta, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(ta, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(ta, 0, i, w, h, 0, f, ty, p)
	return id
}

func uniformLocation(p uint32, n string) int32 {
	l := gl.GetUniformLocation(p, gl.Str(n))
	if l == -1 {
		log.Fatalf("> OpenGL: Uniform %vnot found in shader %v (remember that if you don't use the uniform, it wil automatically get removed by the compiler)\n", n, p)
	}
	return l
}

// attachFramebufferTextures attaches textures to an existing framebuffer
func (o *OpenGL) attachFramebufferTextures(fib uint32, a []attachment) {
	for _, att := range a {
		if *att.textureID <= 0 {
			gl.GenTextures(1, att.textureID)
			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_RECTANGLE, *att.textureID)
			gl.TexImage2D(gl.TEXTURE_RECTANGLE, 0, att.internalFormat, int32(o.screenWidth), int32(o.screenHeight), 0, att.format, att.typ, nil)
			gl.TexParameteri(gl.TEXTURE_RECTANGLE, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
			gl.TexParameteri(gl.TEXTURE_RECTANGLE, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		} else {
			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_RECTANGLE, *att.textureID)
		}
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, att.point, gl.TEXTURE_RECTANGLE, *att.textureID, 0)
	}

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		log.Fatalf("Framebuffer Error Hex: 0x%x\nNow go look it up here: https://raw.githubusercontent.com/go-gl/gl/master/v4.5-core/gl/package.go", gl.CheckFramebufferStatus(gl.FRAMEBUFFER))
	}
	// return err here someday
}

// newFramebuffer creates a new framebuffer
func newFramebuffer() uint32 {
	var fid uint32
	gl.GenFramebuffers(1, &fid)
	return fid
}

func bindFramebuffer(id uint32) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, id)
	gl.Viewport(0, 0, int32(1024), int32(576))
}

type attachment struct {
	textureID      *uint32
	internalFormat int32
	format         uint32
	typ            uint32
	point          uint32
}
