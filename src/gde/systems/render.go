package systems

import (
	"fmt"

	"gde"
)

const (
	VSHADER_OPENGL_4_1 = `#version 410
in vec4 position;
void main()
{
  gl_Position = position;
}
` + "\x00"

	FSHADER_OPENGL_4_1 = `#version 410
out vec4 outColor;
void main()
{
  outColor = vec4(1.0, 1.0, 1.0, 1.0);
}
` + "\x00"

	VSHADER_OPENGL_ES_2_0 = `#version 100
attribute vec4 position;
void main() {
  gl_Position = position;
}`

	FSHADER_OPENGL_ES_2_0 = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
  gl_FragColor = color;
}`
)

type Render struct {
	gde.SystemRoutine

	Properties map[string]interface{}
}

func (r *Render) Init() {
	fmt.Println("Render.Init() executed")
	r.Properties = make(map[string]interface{})
	r.Properties["VertexShader"] = VSHADER_OPENGL_ES_2_0
	r.Properties["FragmentShader"] = FSHADER_OPENGL_ES_2_0
}
func (r *Render) Update() {
	fmt.Println("Render.Update() executed")
}
func (r *Render) End() {
	fmt.Println("Render.End() executed")
}

func (r *Render) Add(engine *gde.Engine) {
	fmt.Println("Render.Add(Engine) executed")
	engine.Systems[fmt.Sprintf("%T", Render{})] = r
}
func (r *Render) Property(key string) interface{} {
	return r.Properties[key]
}
