package render

import (

	// "github.com/go-gl/gl/v4.1-core/gl"

	"./../engine"
)

const activeRenderSystem Render = nil

type Render interface {
	engine.System
	NewShader(vs string, fs string) uint32
	RegisterTransform(*Transform)
	RegisterMesh(*Mesh)
	RegisterMaterial(*Material)
}

// func (r *Render) RegisterTransform(transform *Transform) {
// 	r.transforms = append(r.transforms, transform)
// }

// type Render struct {
// 	platform *engine.Platform
// }

// type RenderRoutine interface {
// 	engine.System

// 	// LoadRenderer(RendererRoutine)
// 	NewShader(string, string) uint32
// 	AddUISystem(*engine.Engine)
// 	GetCamera() *Camera
// }

const (
	VSHADER = `#version 330
  attribute vec4 pos;
  attribute vec3 col;
  attribute vec2 tex;

  uniform mat4 model;
  uniform mat4 view;
  uniform mat4 projection;

  // try use in/out later
  varying vec3 colOut;

  // varying bool hasTexOut;
  varying vec2 texOut;

  void main( void ) {
	gl_Position = projection * view * model * pos;
	// gl_Position = vec4(1.0, 1.0, 0.0, 1.0);
	colOut = col;
	// hasTexOut = hasTex;
	texOut = tex;
  }
  ` + "\x00"

	// SHADER TODO LIST
	// 0. Better understand shaders by doing research
	// 1. Allow mesh color or texture
	// precision mediump float;
	FSHADER = `#version 330
  precision mediump float;

  uniform int hasTexture;
  uniform sampler2D texture;
  varying vec2 texOut;

  uniform vec4 color;
  varying vec3 colOut;

  void main() {
	if (hasTexture == 1) {
	  vec4 frag_texture = texture2D(texture, texOut) * color;
	  if(frag_texture.a < 0.9) {
		discard;
	  }
	  gl_FragColor = frag_texture;
	  // gl_FragColor = vec4(0.4, 0.8, 0.2, 1.0);
	} else {
	  // vec4 frag_color = vec4(1.0, 1.0, 1.0, 1.0);
	  // vec4 frag_color = color;
	  vec4 frag_color = vec4(colOut, 1.0) * color;
	  // vec4 frag_color = vec4(colOut, 1.0);
	  gl_FragColor = frag_color;
	}
  }
  ` + "\x00"
)
