package uigl

import (
	"gde/engine"
	"gde/render"
	"gde/render/ui"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type UIGL struct {
	ui.UI
	ui.UIRoutine
}

func (u *UIGL) Init() {
	// FIX HERE!!!
	// u.ShaderProgram = render.NewShader(VSHADER_OPENGL_ES_2_0_TEXT, FSHADER_OPENGL_ES_2_0_TEXT)
}

func (u *UIGL) AddSubSystem(system render.RenderRoutine) {}

func (u *UIGL) Update(entities *map[string]*engine.Entity) {
	// THIS IS PAINFULL. JUST BITE THE BULLET!
	gl.UseProgram(u.ShaderProgram)

	var view mgl32.Mat4
	view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	view_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("view\x00"))

	var proj mgl32.Mat4
	proj = mgl32.Ortho(0, 360, 640, 0, 0, 10)
	// proj = mgl32.Perspective(mgl32.DegToRad(60.0), float32(320)/640, 0.01, 1000)

	proj_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("projection\x00"))
	model_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("model\x00"))
	// dimensions_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("dimensions\x00"))
	box_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("box\x00"))
	text_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("text_arr\x00"))
	text_scale_uni := gl.GetUniformLocation(u.ShaderProgram, gl.Str("text_scale\x00"))

	for _, v := range *entities {
		uiRenderer := v.GetComponent(&ui.UIRenderer{})
		if uiRenderer == nil {
			continue
		}

		vao := uiRenderer.GetProperty("VAO")
		switch vao := vao.(type) {
		case uint32:
			gl.BindVertexArray(vao)
		}

		// var model mgl32.Mat4
		model := mgl32.Ident4()
		trans := v.GetComponent(&render.Transform{})

		pos := trans.GetProperty("Position")
		switch pos := pos.(type) {
		case render.Vector3:
			// gl.Uniform2fv(box_uni, 1, &[]float32{pos.X, pos.Y}[0])
			scaledX := pos.X // float32(u.Platform.RenderW)
			scaledY := pos.Y // float32(u.Platform.RenderH) * u.Platform.AspectRatio
			model = model.Mul4(mgl32.Translate3D(scaledX, scaledY, pos.Z))

			scale := trans.GetProperty("Dimensions")
			switch scale := scale.(type) {
			case render.Vector2:

				scaledX := scale.X // float32(u.Platform.RenderW)
				scaledY := scale.Y // float32(u.Platform.RenderH)
				model = model.Mul4(mgl32.Scale3D(scaledX, scaledY, 0))

				gl.UniformMatrix4fv(model_uni, 1, false, &model[0])
				gl.UniformMatrix4fv(view_uni, 1, false, &view[0])
				gl.UniformMatrix4fv(proj_uni, 1, false, &proj[0])

				text_arr := uiRenderer.GetProperty("Text")
				switch text_arr := text_arr.(type) {
				case []mgl32.Vec4:
					gl.Uniform4fv(text_uni, int32(len(text_arr)), &text_arr[0][0])

					textBox := render.Vector4{pos.X /*x*/, pos.Y /*y*/, scale.X /*z*/, 640 - (pos.Y) /*w*/}

					text_scale := uiRenderer.GetProperty("Scale")
					switch text_scale := text_scale.(type) {
					case float64:
						gl.Uniform1f(text_scale_uni, float32(text_scale))
					}

					text_padding := uiRenderer.GetProperty("Padding")
					switch text_padding := text_padding.(type) {
					case render.Vector4:
						textBox.X += text_padding.W
						textBox.Z -= text_padding.W
						textBox.Z -= text_padding.Y

						textBox.W -= text_padding.X
						// ONLY APPLIES IF ALIGNED TO BOTTOM
						textBox.Y -= text_padding.X
						textBox.Y -= text_padding.Z
					}
					// THESE CAN BE float64?
					gl.Uniform4fv(box_uni, 1, &[]float32{textBox.X, textBox.Y, textBox.Z, textBox.W}[0])

				}
			}
		}

		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, gl.PtrOffset(0))
	}

	gl.BindVertexArray(0)
}

func (u *UIGL) Stop() {
}

func (r *UIGL) LoadRenderer(renderer render.RendererRoutine) {
	renderer.LoadMesh(&render.Mesh{
		Vertices: []float32{
			1.0, 1.0, -0.1, 1.0, 0.0, 0.0,
			1.0, 0.0, -0.1, 0.0, 1.0, 0.0,
			0.0, 0.0, -0.1, 0.0, 0.0, 1.0,
			0.0, 1.0, -0.1, 0.0, 1.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	// Bind vertex array object. This must wrap around the mesh creation because it is how we are going to access it later when we draw
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	// Vertex buffer
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(renderer.MeshVertices())*4, gl.Ptr(renderer.MeshVertices()), gl.STATIC_DRAW)

	// Element buffer
	var elementBuffer uint32
	gl.GenBuffers(1, &elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(renderer.MeshIndicies())*4, gl.Ptr(renderer.MeshIndicies()), gl.STATIC_DRAW)

	// Linking vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Linking fragment attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Unbind Vertex array object
	renderer.SetProperty("VAO", vertexArrayID)
	// fmt.Printf("VAO Load: %v\n", vertexArrayID)
	gl.BindVertexArray(0)
}

const (
	VSHADER_OPENGL_ES_2_0_TEXT = `#version 120
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
	FSHADER_OPENGL_ES_2_0_TEXT = `#version 120
  uniform vec4 box;
  uniform vec4 text_arr[200];
  uniform float text_scale = 1.0;

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
  float extract_bit(float n, float b)
  {
	b = clamp(b,-1.0,24.0);
	return floor(mod(floor(n / pow(2.0,floor(b))),2.0));   
  }

  //Returns the pixel at uv in the given bit-packed sprite.
  float sprite(vec4 spr, vec2 size, vec2 uv)
  {
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

	for(int i = 0; i < text_arr.length(); i++)
	{
	  if (text_arr[i].w == 1) {
		print_pos = vec2(box.x / text_scale, print_pos.y - STRHEIGHT(1.0));
	  } else {
		if (i > 0.0 && mod(i, wrap) == 0.0) {
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

// precision mediump float;

//   uniform vec2 text_arr[100];
//   uniform vec4 textstart;

//   varying vec3 colOut;

//   #define CHAR_SIZE vec2(6, 7)
//   #define CHAR_SPACING vec2(6, 9)

//   #define text_scale 1.0

//   vec2 start_pos = vec2(0,0);
//   vec2 print_pos = vec2(0,0);
//   vec2 print_pos_pre_move = vec2(0, 0);
//   vec3 text_color = vec3(1, 0, 0);

//   //Text coloring
//   #define HEX(i) text_color = mod(vec3(i / 65536,i / 256,i),vec3(256.0))/255.0;
//   #define RGB(r,g,b) text_color = vec3(r,g,b);

//   #define STRWIDTH(c) (c * CHAR_SPACING.x)
//   #define STRHEIGHT(c) (c * CHAR_SPACING.y)
//   #define BEGIN_TEXT(x,y) print_pos = floor(vec2(x,y)); start_pos = floor(vec2(x,y));

//   //Automatically generated from the sprite sheet here: http://uzebox.org/wiki/index.php?title=File:Font6x8.png
//   #define _ col+=char(vec2(0.0,0.0),uv);
//   #define _spc col+=char(vec2(0.0,0.0),uv)*text_color;
//   #define _exc col+=char(vec2(276705.0,32776.0),uv)*text_color;
//   #define _quo col+=char(vec2(1797408.0,0.0),uv)*text_color;
//   #define _hsh col+=char(vec2(10738.0,1134484.0),uv)*text_color;
//   #define _dol col+=char(vec2(538883.0,19976.0),uv)*text_color;
//   #define _pct col+=char(vec2(1664033.0,68006.0),uv)*text_color;
//   #define _amp col+=char(vec2(545090.0,174362.0),uv)*text_color;
//   #define _apo col+=char(vec2(798848.0,0.0),uv)*text_color;
//   #define _lbr col+=char(vec2(270466.0,66568.0),uv)*text_color;
//   #define _rbr col+=char(vec2(528449.0,33296.0),uv)*text_color;
//   #define _ast col+=char(vec2(10471.0,1688832.0),uv)*text_color;
//   #define _crs col+=char(vec2(4167.0,1606144.0),uv)*text_color;
//   #define _per col+=char(vec2(0.0,1560.0),uv)*text_color;
//   #define _dsh col+=char(vec2(7.0,1572864.0),uv)*text_color;
//   #define _com col+=char(vec2(0.0,1544.0),uv)*text_color;
//   #define _lsl col+=char(vec2(1057.0,67584.0),uv)*text_color;
//   #define _0 col+=char(vec2(935221.0,731292.0),uv)*text_color;
//   #define _1 col+=char(vec2(274497.0,33308.0),uv)*text_color;
//   #define _2 col+=char(vec2(934929.0,1116222.0),uv)*text_color;
//   #define _3 col+=char(vec2(934931.0,1058972.0),uv)*text_color;
//   #define _4 col+=char(vec2(137380.0,1302788.0),uv)*text_color;
//   #define _5 col+=char(vec2(2048263.0,1058972.0),uv)*text_color;
//   #define _6 col+=char(vec2(401671.0,1190044.0),uv)*text_color;
//   #define _7 col+=char(vec2(2032673.0,66576.0),uv)*text_color;
//   #define _8 col+=char(vec2(935187.0,1190044.0),uv)*text_color;
//   #define _9 col+=char(vec2(935187.0,1581336.0),uv)*text_color;
//   #define _col col+=char(vec2(195.0,1560.0),uv)*text_color;
//   #define _scl col+=char(vec2(195.0,1544.0),uv)*text_color;
//   #define _les col+=char(vec2(135300.0,66052.0),uv)*text_color;
//   #define _equ col+=char(vec2(496.0,3968.0),uv)*text_color;
//   #define _grt col+=char(vec2(528416.0,541200.0),uv)*text_color;
//   #define _que col+=char(vec2(934929.0,1081352.0),uv)*text_color;
//   #define _ats col+=char(vec2(935285.0,714780.0),uv)*text_color;
//   #define _A col+=char(vec2(935188.0,780450.0),uv)*text_color;
//   #define _B col+=char(vec2(1983767.0,1190076.0),uv)*text_color;
//   #define _C col+=char(vec2(935172.0,133276.0),uv)*text_color;
//   #define _D col+=char(vec2(1983764.0,665788.0),uv)*text_color;
//   #define _E col+=char(vec2(2048263.0,1181758.0),uv)*text_color;
//   #define _F col+=char(vec2(2048263.0,1181728.0),uv)*text_color;
//   #define _G col+=char(vec2(935173.0,1714334.0),uv)*text_color;
//   #define _H col+=char(vec2(1131799.0,1714338.0),uv)*text_color;
//   #define _I col+=char(vec2(921665.0,33308.0),uv)*text_color;
//   #define _J col+=char(vec2(66576.0,665756.0),uv)*text_color;
//   #define _K col+=char(vec2(1132870.0,166178.0),uv)*text_color;
//   #define _L col+=char(vec2(1065220.0,133182.0),uv)*text_color;
//   #define _M col+=char(vec2(1142100.0,665762.0),uv)*text_color;
//   #define _N col+=char(vec2(1140052.0,1714338.0),uv)*text_color;
//   #define _O col+=char(vec2(935188.0,665756.0),uv)*text_color;
//   #define _P col+=char(vec2(1983767.0,1181728.0),uv)*text_color;
//   #define _Q col+=char(vec2(935188.0,698650.0),uv)*text_color;
//   #define _R col+=char(vec2(1983767.0,1198242.0),uv)*text_color;
//   #define _S col+=char(vec2(935171.0,1058972.0),uv)*text_color;
//   #define _T col+=char(vec2(2035777.0,33288.0),uv)*text_color;
//   #define _U col+=char(vec2(1131796.0,665756.0),uv)*text_color;
//   #define _V col+=char(vec2(1131796.0,664840.0),uv)*text_color;
//   #define _W col+=char(vec2(1131861.0,699028.0),uv)*text_color;
//   #define _X col+=char(vec2(1131681.0,84130.0),uv)*text_color;
//   #define _Y col+=char(vec2(1131794.0,1081864.0),uv)*text_color;
//   #define _Z col+=char(vec2(1968194.0,133180.0),uv)*text_color;
//   #define _lsb col+=char(vec2(925826.0,66588.0),uv)*text_color;
//   #define _rsl col+=char(vec2(16513.0,16512.0),uv)*text_color;
//   #define _rsb col+=char(vec2(919584.0,1065244.0),uv)*text_color;
//   #define _pow col+=char(vec2(272656.0,0.0),uv)*text_color;
//   #define _usc col+=char(vec2(0.0,62.0),uv)*text_color;
//   #define _a col+=char(vec2(224.0,649374.0),uv)*text_color;
//   #define _b col+=char(vec2(1065444.0,665788.0),uv)*text_color;
//   #define _c col+=char(vec2(228.0,657564.0),uv)*text_color;
//   #define _d col+=char(vec2(66804.0,665758.0),uv)*text_color;
//   #define _e col+=char(vec2(228.0,772124.0),uv)*text_color;
//   #define _f col+=char(vec2(401543.0,1115152.0),uv)*text_color;
//   #define _g col+=char(vec2(244.0,665474.0),uv)*text_color;
//   #define _h col+=char(vec2(1065444.0,665762.0),uv)*text_color;
//   #define _i col+=char(vec2(262209.0,33292.0),uv)*text_color;
//   #define _j col+=char(vec2(131168.0,1066252.0),uv)*text_color;
//   #define _k col+=char(vec2(1065253.0,199204.0),uv)*text_color;
//   #define _l col+=char(vec2(266305.0,33292.0),uv)*text_color;
//   #define _m col+=char(vec2(421.0,698530.0),uv)*text_color;
//   #define _n col+=char(vec2(452.0,1198372.0),uv)*text_color;
//   #define _o col+=char(vec2(228.0,665756.0),uv)*text_color;
//   #define _p col+=char(vec2(484.0,667424.0),uv)*text_color;
//   #define _q col+=char(vec2(244.0,665474.0),uv)*text_color;
//   #define _r col+=char(vec2(354.0,590904.0),uv)*text_color;
//   #define _s col+=char(vec2(228.0,114844.0),uv)*text_color;
//   #define _t col+=char(vec2(8674.0,66824.0),uv)*text_color;
//   #define _u col+=char(vec2(292.0,1198868.0),uv)*text_color;
//   #define _v col+=char(vec2(276.0,664840.0),uv)*text_color;
//   #define _w col+=char(vec2(276.0,700308.0),uv)*text_color;
//   #define _x col+=char(vec2(292.0,1149220.0),uv)*text_color;
//   #define _y col+=char(vec2(292.0,1163824.0),uv)*text_color;
//   #define _z col+=char(vec2(480.0,1148988.0),uv)*text_color;
//   #define _lpa col+=char(vec2(401542.0,66572.0),uv)*text_color;
//   #define _bar col+=char(vec2(266304.0,33288.0),uv)*text_color;
//   #define _rpa col+=char(vec2(788512.0,1589528.0),uv)*text_color;
//   #define _tid col+=char(vec2(675840.0,0.0),uv)*text_color;
//   #define _lar col+=char(vec2(8387.0,1147904.0),uv)*text_color;
//   #define _nl print_pos = start_pos - vec2(0,CHAR_SPACING.y);

//   //Extracts bit b from the given number.
//   float extract_bit(float n, float b)
//   {
// 	b = clamp(b,-1.0,22.0);
// 	return floor(mod(floor(n / pow(2.0,floor(b))),2.0));
//   }

//   //Returns the pixel at uv in the given bit-packed sprite.
//   float sprite(vec2 spr, vec2 size, vec2 uv)
//   {
// 	uv = floor(uv);
// 	float bit = (size.x-uv.x-1.0) + uv.y * size.x;
// 	bool bounds = all(greaterThanEqual(uv,vec2(0)))&& all(lessThan(uv,size));
// 	return bounds ? extract_bit(spr.x, bit - 21.0) + extract_bit(spr.y, bit) : 0.0;
//   }

//   //Prints a character and moves the print position forward by 1 character width.
//   vec3 char(vec2 ch, vec2 uv)
//   {
// 	float px = sprite(ch, CHAR_SIZE, uv - print_pos);
// 	print_pos.x += CHAR_SPACING.x;
// 	return vec3(px);
//   }

//   vec3 Text(vec2 uv)
//   {
// 	vec3 col = vec3(0.0);

// 	// BEGIN_TEXT(posOut.x, posOut.y * 320)
// 	// BEGIN_TEXT(posOut.x, posOut.y * ((640  / DOWN_SCALE) - CHAR_SIZE.y))
// 	BEGIN_TEXT(textstart.x, textstart.y)
// 	for(int i = 0; i < text_arr.length(); i++)
// 	{
// 	  col+=char(text_arr[i],uv)*text_color;
// 	}

// 	return col;
//   }

//   void main( void ) {
// 	vec2 uv = gl_FragCoord.xy / DOWN_SCALE;
// 	vec2 duv = floor(gl_FragCoord.xy / DOWN_SCALE);

// 	gl_FragColor = vec4(Text(duv), 1.0);
//   }
