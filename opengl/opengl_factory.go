// +build deploy opengl

package opengl

import (
	"fmt"
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/go-gl/gl/all-core/gl"
	"log"
	"os"
	"strings"
)

func init() {
	engine.NewIntegrant(&OpenGLFactory{})
	fmt.Println("> OpenGLFactory: Ready")
}

// OpenGLFactory used to override basic render system functions
type OpenGLFactory struct {
	engine.Integrant
}

// InitIntegrant called when the integrant plugs into the engine
func (o *OpenGLFactory) InitIntegrant() {}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (o *OpenGLFactory) AddIntegrant(integrant engine.IntegrantRoutine) {}

// Destroy called when engine is gracefully shutting down
func (o *OpenGLFactory) Destroy() {}

// NewShader creates a new vertex and fragment shader
func (o *OpenGLFactory) NewShader(vs string, fs string) uint32 {
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

	// Use this program for all upcoming render calls
	gl.UseProgram(shaderProgram)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	return shaderProgram
}

// OpenGLMesh defines a mesh (opengl)
type OpenGLMesh struct {
	m   *render.Mesh
	vao uint32
}

// SetVAO sets the VAO (opengl)
func (m *OpenGLMesh) SetVAO(vao uint32) {
	m.vao = vao
}

// VAO returns a opengl VAO
func (m *OpenGLMesh) VAO() uint32 {
	return m.vao
}

// OpenGLTexture defines a texture (opengl)
type OpenGLTexture struct {
	*render.Texture
	id uint32
}

// ID returns a opengl texture id
func (t *OpenGLTexture) ID() uint32 {
	return t.id
}

// OpenGLMaterial defines a material (opengl)
type OpenGLMaterial struct {
	*render.Material
	texture *OpenGLTexture
}

// NewOpenGLMaterial creates a meterial (opengl)
func NewOpenGLMaterial(m *render.Material) *OpenGLMaterial {
	openGLMaterial := &OpenGLMaterial{Material: m}
	return openGLMaterial
}

// OverrideTexture overrides base texture (opengl)
func (o *OpenGLMaterial) OverrideTexture(fn func(render.Textureable)) {
	o.SetTexture = fn
	o.SetTexture(o.BaseTexture().(*render.Texture))
}

// Texture returns a opengl texture
func (o *OpenGLMaterial) Texture() *OpenGLTexture {
	return o.texture
}

// ID returns a opengl material id
func (o *OpenGLMaterial) ID() uint32 {
	return o.texture.id
}
