package ui

import (
	"gde/engine"
	"gde/render"
)

type UI struct {
	render.RenderRoutine
	ShaderProgram uint32
	Platform      *engine.Platform
	mesh          *render.Mesh
}

type UIRoutine interface {
	Init()
	LoadRenderer(render.RendererRoutine)
	NewShader(string, string) uint32
	AddUISystem(*engine.Engine)
	Stop()
	Update(*map[string]*engine.Entity)
}

func (u *UI) Init() {}

func (u *UI) AddUISystem(game *engine.Engine) {}

func (u *UI) Update(entities *map[string]*engine.Entity) {}

func (u *UI) Stop() {}

func (r *UI) LoadRenderer(renderer render.RendererRoutine) {}

const (
	VSHADER_OPENGL_ES_2_0_TEXT = `#version 100
  attribute vec4 pos;
  attribute vec3 col;

  uniform mat4 model;
  uniform mat4 view;
  uniform mat4 projection;

  varying vec3 colOut;

  void main() {
	gl_Position = projection * view * model * pos;
	colOut = col;
  }
  ` + "\x00"

	// SHADER TODO LIST
	// 1. Allow mesh color or texture
	// 2. Send container width for #3 to work
	FSHADER_OPENGL_ES_2_0_TEXT = `#version 100
	// ONLY IN GLES
  precision mediump float;

  uniform vec4 box;
  uniform vec4 text_arr[200];
  const float text_scale = 1.0;

  varying vec3 colOut;

  #define CHAR_SIZE vec2(8, 12)
  #define CHAR_SPACING vec2(8, 12)

  #define STRWIDTH(c) (c * CHAR_SPACING.x)
  #define STRHEIGHT(c) (c * CHAR_SPACING.y)

  #define NORMAL 0
  #define INVERT 1
  #define UNDERLINE 2

  int TEXT_MODE = NORMAL;

  // STILL DONT KNOW THIS ONE
  vec4 ch_lar = vec4(0x000000,0x10386C,0xC6C6FE,0x000000);

  vec2 res = vec2(360.0, 640.0) / text_scale;
  vec2 print_pos = vec2(0);

  //Extracts bit b from the given number.
  //Shifts bits right (num / 2^bit) then ANDs the result with 1 (mod(result,2.0)).
  float extract_bit(float n, float b) {
	b = clamp(b,-1.0,24.0);
	return floor(mod(floor(n / pow(2.0,floor(b))),2.0));   
  }

  //Returns the pixel at uv in the given bit-packed sprite.
  float sprite(vec4 spr, vec2 size, vec2 uv) {
	uv = floor(uv);

	//Calculate the bit to extract (x + y * width) (flipped on x-axis)
	float bit = (size.x-uv.x-1.0) + uv.y * size.x;

	//Clipping bound to remove garbage outside the sprite's boundaries.
	bool bounds = all(greaterThanEqual(uv,vec2(0))) && all(lessThan(uv,size));

	float pixels = 0.0;
	pixels += extract_bit(spr.x, bit - 72.0);
	pixels += extract_bit(spr.y, bit - 48.0);
	pixels += extract_bit(spr.z, bit - 24.0);
	pixels += extract_bit(spr.w, bit - 00.0);

	return bounds ? pixels : 0.0;
  }

  //Prints a character and moves the print position forward by 1 character width.
  float char(vec4 ch, vec2 uv)
  {
	if( TEXT_MODE == INVERT )
	{
	  //Inverts all of the bits in the character.
	  ch = pow(2.0,24.0)-1.0-ch;
	}
	if( TEXT_MODE == UNDERLINE )
	{
	  //Makes the bottom 8 bits all 1.
	  //Shifts the bottom chunk right 8 bits to drop the lowest 8 bits,
	  //then shifts it left 8 bits and adds 255 (binary 11111111).
	  ch.w = floor(ch.w/256.0)*256.0 + 255.0;  
	}

	float px = sprite(ch, CHAR_SIZE, uv - print_pos);
	print_pos.x += CHAR_SPACING.x;
	return px;
  }

  float text(vec2 uv)
  {
	float col = 0.0;
	float wrap = floor((box.z / text_scale) / CHAR_SIZE.x);

	print_pos = vec2(box.x / text_scale, (box.w / text_scale) - STRHEIGHT(1.0));

	// for(int i = 0; i < text_arr.length(); i++)
	for(int i = 0; i < 5; i++)
	{
	  if (text_arr[i].w == 1.0) {
		print_pos = vec2(box.x / text_scale, print_pos.y - STRHEIGHT(1.0));
	  } else {
		if (i > 0 && mod(float(i), wrap) == 0.0) {
		  // Warning: Expected prototype is 'mod (float, float)'
		  print_pos = vec2(box.x / text_scale, print_pos.y - STRHEIGHT(1.0));
		}
		col += char(text_arr[i],uv); 
	  }
	}

	return col;
  }

  void main()
  {
	vec2 uv = gl_FragCoord.xy / text_scale;
	vec2 duv = floor(gl_FragCoord.xy / text_scale);

	float pixel = text(duv);

	vec3 col = mix(vec3(1.0),vec3(0,0,0),pixel);

	gl_FragColor =vec4(0.6, 0.6, 0.6, 0.6) * vec4(vec3(col), 1.0);
  }
  ` + "\x00"
)
