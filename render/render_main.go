// +build render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Render interface {
	NewShader(vs string, fs string) uint32
	UnRegisterEntity(*engine.Entity)
	RegisterTransform(*std.Transform)
	RegisterMesh(*Mesh)
	RegisterMaterial(*Material)
	RegisterCamera(*Camera)
}

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
	  if(frag_texture.a < 0.1) {
		discard;
	  }
	  gl_FragColor = frag_texture;
	} else {
	  vec4 frag_color = vec4(colOut, 1.0) * color;
	  gl_FragColor = frag_color;
	}
  }
  ` + "\x00"
)
