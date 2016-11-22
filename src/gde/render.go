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

// gl_Position = projection * view * model * position;
const (
	VSHADER_OPENGL_ES_2_0 = `#version 100
  attribute vec4 pos;
  attribute vec3 col;
  attribute vec2 tex;

  uniform mat4 model;
  uniform mat4 view;
  uniform mat4 projection;

  // try use in/out later
  varying vec3 colOut;
  varying vec2 texOut;

  void main() {
	gl_Position = projection * view * model * pos;
	colOut = col;
	texOut = tex;
  }`

	FSHADER_OPENGL_ES_2_0 = `#version 100
  precision mediump float;

  uniform sampler2D texture;

  varying vec3 colOut;
  varying vec2 texOut;

  void main() {
	gl_FragColor = texture2D(texture, texOut);
	// gl_FragColor = vec4(colOut, 1.0);
  }`
)
