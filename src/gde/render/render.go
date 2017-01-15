package render

import (

	// "github.com/go-gl/gl/v4.1-core/gl"

	"gde/engine"
)

type Render struct {
	platform *engine.Platform
}

type RenderRoutine interface {
	engine.System

	LoadRenderer(RendererRoutine)
	NewShader(string, string) uint32
	AddSubSystem(RenderRoutine)
}

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

  void main( void ) {
	gl_Position = projection * view * model * pos;
	colOut = col;
	texOut = tex;
  }
  ` + "\x00"

	// SHADER TODO LIST
	// 0. Better understand shaders by doing research
	// 1. Allow mesh color or texture
	FSHADER_OPENGL_ES_2_0 = `#version 100
  precision mediump float;

  uniform sampler2D texture;

  varying vec3 colOut;
  varying vec2 texOut;

  void main() {
	// gl_FragColor = vec4(.8, .3, .3, 1.);
	gl_FragColor = texture2D(texture, texOut);
  }
  ` + "\x00"
)
