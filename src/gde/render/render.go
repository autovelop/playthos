package render

import (
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"

	"gde/engine"
)

type Render struct {
	platform *engine.Platform
}

type RenderRoutine interface {
	engine.System

	GetPlatform() *engine.Platform
	LoadRenderer(RendererRoutine)
	NewShader(string, string) uint32
	AddSubSystem(RenderRoutine)
}

func (r *Render) GetPlatform() *engine.Platform {
	return r.platform
}

func NewShader(vShader string, fShader string) uint32 {
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("Render > OpenGL > Version: %v", version)
	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(vShader)
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

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, vShader)
		os.Exit(0)
	}

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(fShader)
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

		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", logMsg, vShader)
		os.Exit(0)
	}

	// Create program
	var shaderProgram uint32
	shaderProgram = gl.CreateProgram()

	gl.AttachShader(shaderProgram, fshader)
	gl.AttachShader(shaderProgram, vshader)

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

const (
	VSHADER_OPENGL_ES_2_0 = `#version 120
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
	FSHADER_OPENGL_ES_2_0 = `#version 120
  // precision mediump float;

  uniform sampler2D texture;

  varying vec3 colOut;
  varying vec2 texOut;

  void main() {
	gl_FragColor = texture2D(texture, texOut);
  }
  ` + "\x00"
)
