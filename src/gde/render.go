package gde

type Render struct {
	OpenGL string
	GLSL   string
}

type RenderRoutine interface {
	SystemRoutine

	LoadRenderer(*Renderer)
}

func (r *Render) Init()                                     {}
func (r *Render) End()                                      {}
func (r *Render) Update(entities *map[string]EntityRoutine) {}

const (
	VSHADER_OPENGL_ES_2_0 = `#version 100
  attribute vec4 position;
  uniform mat4 transform;
  void main() {
	gl_Position = transform * position;
  }`

	FSHADER_OPENGL_ES_2_0 = `#version 100
  precision mediump float;
  void main() {
	gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
  }`
)
