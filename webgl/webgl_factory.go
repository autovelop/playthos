// +build deploy webgl

package webgl

import (
	"fmt"
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	engine.NewIntegrant(&WebGLFactory{})
	fmt.Println("> WebGLFactory: Ready")
}

type WebGLFactory struct {
	engine.Integrant
}

func (w *WebGLFactory) InitIntegrant() {}

func (w *WebGLFactory) Destroy() {}

func (w *WebGLFactory) NewShader(gl *Context, vs string, fs string) *js.Object {
	version := gl.GetParameter(gl.VERSION)
	fmt.Printf("> WebGLFactory: Profile = %v\n", version)

	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vshader, string(vs))
	gl.CompileShader(vshader)
	if !gl.GetShaderParameterb(vshader, gl.COMPILE_STATUS) {
		fmt.Printf("> WebGLFactory: Vertex Shader %v\n%v\n", gl.GetShaderInfoLog(vshader), vs)
		return nil
	}

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fshader, string(fs))
	gl.CompileShader(fshader)
	if !gl.GetShaderParameterb(fshader, gl.COMPILE_STATUS) {
		fmt.Printf("> WebGLFactory: Fragment Shader %v\n", gl.GetShaderInfoLog(fshader))
		return nil
	}

	// Create program
	var shaderProgram *js.Object
	shaderProgram = gl.CreateProgram()

	gl.AttachShader(shaderProgram, vshader)
	gl.AttachShader(shaderProgram, fshader)

	// Link program
	gl.LinkProgram(shaderProgram)

	if !gl.GetProgramParameterb(shaderProgram, gl.LINK_STATUS) {
		fmt.Printf("> WebGLFactory: Shader Linking %v\n", gl.GetProgramInfoLog(shaderProgram))
		return nil
	}

	// var statisLink int32
	// gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &statisLink)
	// if statisLink == gl.FALSE {
	// 	var logLength int32
	// 	gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

	// 	logMsg := strings.Repeat("\x00", int(logLength+1))
	// 	gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(logMsg))

	// 	log.Printf("\n\n ### SHADER LINK ERROR ### \n%v\n\n", logMsg)
	// 	os.Exit(0)
	// }

	// Use this program for all upcoming render calls
	gl.UseProgram(shaderProgram)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	return shaderProgram
}

type WebGLMesh struct {
	m   *render.Mesh
	vao *js.Object
}

func (m *WebGLMesh) SetVAO(vao *js.Object) {
	m.vao = vao
}

func (m *WebGLMesh) VAO() *js.Object {
	return m.vao
}

type WebGLTexture struct {
	*render.Texture
	id *js.Object
}

func (t *WebGLTexture) ID() *js.Object {
	return t.id
}

type WebGLMaterial struct {
	*render.Material
	texture *WebGLTexture
}

func NewWebGLMaterial(m *render.Material) *WebGLMaterial {
	webGLMaterial := &WebGLMaterial{Material: m}
	return webGLMaterial
}

func (w *WebGLMaterial) OverrideTexture(fn func(render.Textureable)) {
	w.SetTexture = fn
	w.SetTexture(w.BaseTexture().(*render.Texture))
}

func (w *WebGLMaterial) Texture() *WebGLTexture {
	return w.texture
}
func (w *WebGLMaterial) ID() *js.Object {
	return w.texture.id
}
