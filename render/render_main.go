// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Render defines an empty system that will be overwritten by a platform specific system
type Render struct {
	engine.System
}

// NewRenderSystem instructs engine to inject render system
func NewRenderSystem(render RenderRoutine) {
	engine.NewSystem(render)
}

// RenderRoutine interface to allow implemented systems to register drawer components
type RenderRoutine interface {
	engine.SystemRoutine
	// engine.Drawer
	// NewShader(vs string, fs string) uint32
	// UnRegisterEntity(*engine.Entity)
	RegisterTransform(*std.Transform)
	RegisterMesh(*Mesh)
	RegisterMaterial(*Material)
	RegisterCamera(*Camera)
}

const (
	VSHADEROPAQUE = `#version 450 core
	layout(location=0) in vec3 inVertexPosition;
	layout(location=1) in vec2 inTexCoord;

	layout(location=0) out vec2 vTexCoord;

	uniform mat4 uModelMatrix;
	uniform mat4 uViewMatrix;
	uniform mat4 uProjMatrix;

	void main(void)
	{
		gl_Position = uProjMatrix * uViewMatrix * uModelMatrix * vec4(inVertexPosition, 1.0);
		vTexCoord = inTexCoord;
	}` + "\x00"

	FSHADEROPAQUE = `#version 450 core
	layout(location=0) in vec2 vTexCoord;

	uniform vec4 uColor;
	// uniform int uTextured;
	// uniform sampler2D uTexture;

	layout (location = 0) out vec4 fragColor;

	void main() {
		// vec4 Ci = uColor;
		// if (uTextured == 1) {
		// 	Ci *= texture(uTexture, vTexCoord);
		// }
		fragColor = uColor;
	}` + "\x00"
	VSHADERTRANSPARENT = `#version 450 core
	layout(location=0) in vec3 inVertexPosition;
	layout(location=1) in vec2 inTexCoord;

	layout(location=0) out vec2 vTexCoord;

	uniform mat4 uModelMatrix;
	uniform mat4 uViewMatrix;
	uniform mat4 uProjMatrix;

	void main(void)
	{
		gl_Position = uProjMatrix * uViewMatrix * uModelMatrix * vec4(inVertexPosition, 1.0);
		vTexCoord = inTexCoord;
	}` + "\x00"

	FSHADERTRANSPARENT = `#version 450 core
	layout(location=0) in vec2 vTexCoord;

	uniform vec4 uColor;
	uniform int uTextured;
	uniform sampler2D uTexture;

	layout(location=0, index=0) out vec4 oSumColor;
	layout(location=0, index=1) out vec4 oSumWeight;

	void main() {
		vec4 color = uColor;
		if (uTextured == 1) {
			color = texture(uTexture, vTexCoord);
		}

		float viewDepth = abs(1.0 / gl_FragCoord.w);

		// Tuned to work well with FP16 accumulation buffers and 0.001 < linearDepth < 2.5
		// See Equation (9) from http://jcgt.org/published/0002/02/09/
		// float linearDepth = viewDepth * uDepthScale;
		float linearDepth = viewDepth * 0.4;

		float weight = clamp(0.03 / (1e-5 + pow(linearDepth, 4.0)), 1e-2, 3e3);

		oSumColor = vec4(color.rgb * color.a, color.a);
		// oSumColor = vec4(color.rgb * color.a, color.a) * weight;
		oSumWeight = vec4(color.a);
	}` + "\x00"

	VSHADERCOMPOSITE = `#version 450 core
	layout(location=0) in vec3 inVertexPosition;

	uniform mat4 uModelMatrix;
	uniform mat4 uViewMatrix;
	uniform mat4 uProjMatrix;

	void main(void)
	{
		gl_Position = vec4(inVertexPosition, 1.0);
	}` + "\x00"

	FSHADERCOMPOSITE = `#version 450 core
	layout(location=0) out vec4 outColor;

	uniform sampler2DRect ColorTex0;
	uniform sampler2DRect ColorTex1;
	uniform sampler2D ColorTex2; // opaque objects

	out vec4 oFragColor;

	void main() {
		vec4 sumColor = texture(ColorTex0, gl_FragCoord.xy);
		float transmittance = texture(ColorTex1, gl_FragCoord.xy).g;
		vec3 averageColor = sumColor.rgb / max(sumColor.a, 0.00001);
		// oFragColor.rgb = averageColor * (1 - transmittance) + vec3(0.0, 0.0, 0.0) * transmittance;
		// oFragColor.rgb = averageColor * (1 - transmittance) + vec3(0.0, 0.0, 0.0) * transmittance;
		oFragColor.rgb = averageColor * (1 - transmittance) + vec3(0.75, 0.75, 0.75) * transmittance;
		// oFragColor.rgb = averageColor * (1 - transmittance);
		// oFragColor.rgb = averageColor;
		// oFragColor.rgb = vec3(transmittance);
	}
	` + "\x00"
	VSHADERFRAMEBUFFER = `#version 450 core
	layout(location=0) in vec2 inVertexPosition;
	layout(location=1) in vec2 inTexCoord;

	layout(location=0) out vec2 vTexCoord;

	void main(void)
	{
		gl_Position = vec4(inVertexPosition, 0.0, 1.0);
		vTexCoord = inTexCoord;
	}` + "\x00"

	FSHADERFRAMEBUFFER = `#version 450 core
	layout(location=0) in vec2 vTexCoord;

	layout(location=0) out vec4 oFragColor;

	uniform sampler2D Texture;

	void main() {
		oFragColor = texture(Texture, vTexCoord);
	}
	` + "\x00"
	// VSHADER OpenGL 4.5 Vertex Shader
	VSHADER = `#version 450 core
	layout (location = 0) in vec4 iVertexPosition;
	layout (location = 1) in vec2 iTexCoord;

	uniform mat4 uModelMatrix;
	uniform mat4 uViewMatrix;
	uniform mat4 uProjMatrix;

	layout (location = 0) out vec2 oTexCoord;

	void main( void ) {
		gl_Position = uProjMatrix * uViewMatrix * uModelMatrix * iVertexPosition;
		oTexCoord = iTexCoord;
	}
	` + "\x00"

	// FSHADER OpenGL 4.5 Fragment Shader
	FSHADER = `#version 450 core

	uniform vec4 uColor;

	uniform int uTextured;
	uniform sampler2D uTexture;

	layout (location = 0) in vec2 iTexCoord;

	uniform vec2 uTextureOffset;
	uniform vec2 uTextureTiling;

	layout (location = 0) out vec4 oColor;

	void main(void) {
		oColor = uColor;
		if (uTextured == 1) {
			oColor *= texture(uTexture, (iTexCoord * (1.0 - uTextureTiling)) + (uTextureOffset * (1.0 - uTextureTiling)));
		}
	}
	` + "\x00"

	// VSCREENSHADER OpenGL 4.5 Vertex Screen Shader (Multipass Rendering)
	VSCREENSHADER = `#version 450 core
	layout (location = 0) in vec2 pos;
	layout (location = 1) in vec2 tex;

	layout (location = 0) out vec2 texOut;

	void main( void ) {
		texOut = tex;
		gl_Position = vec4(pos.x, pos.y, 0.0, 1.0);
	}
	` + "\x00"

	// FSCREENSHADER OpenGL 4.5 Fragment Shader (Multipass Rendering)
	FSCREENSHADER = `#version 450 core
	uniform sampler2D textu;
	layout (location = 0) in vec2 texOut;

	layout (location = 0) out vec4 fragColor;

	void main() {
		vec3 frag_texture = texture(textu, texOut).rgb;
		fragColor = vec4(frag_texture, 1.0);
	}
	` + "\x00"

	// VSHADERWEB OpenGLES 3.0 Vertex Shader
	VSHADERWEB = `#version 300 es

	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	out vec3 colOut;

	out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		colOut = col;
		texOut = tex;
	}
	`

	// FSHADERWEB OpenGLES 3.0 Fragment Shader
	FSHADERWEB = `#version 300 es

	precision mediump float;

	uniform vec4 color;
	in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	in vec2 texOut;
	uniform vec2 spriteScaler;
	uniform vec2 spriteOffset;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec4 frag_texture = texture(textu, (texOut * spriteScaler) + (spriteOffset * spriteScaler)) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	`
	// VSHADER41 OpenGL 4.1 Vertex Shader
	VSHADER41 = `#version 410 core
	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	// try use in/out later
	layout (location = 0) out vec3 colOut;

	// varying bool hasTexOut;
	layout (location = 1) out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		colOut = col;
		texOut = tex;
	}
	` + "\x00"

	// FSHADER41 OpenGL 4.1 Fragment Shader
	FSHADER41 = `#version 410 core

	uniform vec4 color;
	layout (location = 0) in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	layout (location = 1) in vec2 texOut;
	uniform vec2 spriteScaler;
	uniform vec2 spriteOffset;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec4 frag_texture = texture(textu, (texOut * spriteScaler) + (spriteOffset * spriteScaler)) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	` + "\x00"

	// VSHADER33 OpenGL 3.3 Vertex Shader
	VSHADER33 = `#version 330 core
	#extension GL_ARB_separate_shader_objects : enable

	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	// try use in/out later
	layout (location = 0) out vec3 colOut;

	// varying bool hasTexOut;
	layout (location = 1) out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		colOut = col;
		texOut = tex;
	}
	` + "\x00"

	// FSHADER33 OpenGL 3.3 Fragment Shader
	FSHADER33 = `#version 330 core
	#extension GL_ARB_separate_shader_objects : enable

	uniform vec4 color;
	layout (location = 0) in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	layout (location = 1) in vec2 texOut;
	uniform vec2 spriteScaler;
	uniform vec2 spriteOffset;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec4 frag_texture = texture(textu, (texOut * spriteScaler) + (spriteOffset * spriteScaler)) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	` + "\x00"
)
