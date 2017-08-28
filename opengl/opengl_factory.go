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

type OpenGLFactory struct {
	engine.Integrant
}

func (o *OpenGLFactory) InitIntegrant() {}

func (o *OpenGLFactory) Destroy() {}

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

type OpenGLMesh struct {
	m   *render.Mesh
	vao uint32
}

func (m *OpenGLMesh) SetVAO(vao uint32) {
	m.vao = vao
}

func (m *OpenGLMesh) VAO() uint32 {
	return m.vao
}

type OpenGLTexture struct {
	*render.Texture
	id uint32
}

func (t *OpenGLTexture) ID() uint32 {
	return t.id
}

type OpenGLMaterial struct {
	*render.Material
	texture *OpenGLTexture
}

func NewOpenGLMaterial(m *render.Material) *OpenGLMaterial {
	openGLMaterial := &OpenGLMaterial{Material: m}
	return openGLMaterial
}

func (o *OpenGLMaterial) OverrideTexture(fn func(render.Textureable)) {
	o.SetTexture = fn
	o.SetTexture(o.BaseTexture().(*render.Texture))
}

func (o *OpenGLMaterial) Texture() *OpenGLTexture {
	return o.texture
}
func (o *OpenGLMaterial) ID() uint32 {
	return o.texture.id
}
