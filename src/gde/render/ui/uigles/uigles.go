package uigles

import (
	"github.com/go-gl/mathgl/mgl32"

	"golang.org/x/mobile/gl"

	"gde/engine"
	"gde/render"
	"gde/render/ui"
	"log"
)

// Die nuwe UIGLES system is in, maar maak seker dit werk selfde as UIGL system in terme van uniforms, attributes, shaders, en orho!

type UIGLES struct {
	ui.UI
	ui.UIRoutine
	Context   gl.Context
	glProgram gl.Program
}

func (u *UIGLES) Init() {
	log.Printf("UIGLES > Init")
	u.ShaderProgram = u.NewShader(ui.VSHADER_OPENGL_ES_2_0_TEXT, ui.FSHADER_OPENGL_ES_2_0_TEXT)
}

func (u *UIGLES) Update(entities *map[string]*engine.Entity) {
	log.Printf("\n\nUIGLES UPDATE %+v\n\n\n", u.Context)
	// u.Context.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	u.Context.UseProgram(u.glProgram)

	var view mgl32.Mat4
	view = mgl32.LookAtV(mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	view_uni := u.Context.GetUniformLocation(u.glProgram, "view")

	var proj mgl32.Mat4
	// still need to understand what is going on here and how it relates to device
	proj = mgl32.Ortho(0, 1, 2, 0, 0.1, 1000)

	proj_uni := u.Context.GetUniformLocation(u.glProgram, "projection")
	model_uni := u.Context.GetUniformLocation(u.glProgram, "model")

	box_uni := u.Context.GetUniformLocation(u.glProgram, "box")
	text_uni := u.Context.GetUniformLocation(u.glProgram, "text_arr")
	text_scale_uni := u.Context.GetUniformLocation(u.glProgram, "text_scale")

	for _, v := range *entities {
		renderer := v.GetComponent(&ui.UIRenderer{})
		if renderer == nil {
			continue
		}

		vb := renderer.GetProperty("VB")
		switch vb := vb.(type) {
		case gl.Buffer:
			u.Context.BindBuffer(gl.ARRAY_BUFFER, vb)
		}

		eb := renderer.GetProperty("EB")
		switch eb := eb.(type) {
		case gl.Buffer:
			u.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, eb)
		}

		// position := renderer.GetProperty("EB")
		// switch position := position.(type) {
		// case gl.Attrib:
		// 	// Need to figure out again what this is for...
		// 	r.Context.EnableVertexAttribArray(position)
		// 	r.Context.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
		// }

		// var model mgl32.Mat4
		// trans := v.GetComponent(&render.Transform{})

		// pos := trans.GetProperty("Position")
		// switch pos := pos.(type) {
		// case render.Vector3:
		// 	model = mgl32.Translate3D(pos.X, pos.Y, pos.Z)
		// }

		// rot := trans.GetProperty("Rotation")
		// switch rot := rot.(type) {
		// case render.Vector3:
		// 	model = model.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(rot.X)).Mat4())
		// 	model = model.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(rot.Y)).Mat4())
		// 	model = model.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(rot.Z)).Mat4())
		// }

		// texture := renderer.GetProperty("TEXTURE")
		// switch texture := texture.(type) {
		// case gl.Texture:
		// 	u.Context.ActiveTexture(gl.TEXTURE0)
		// 	u.Context.BindTexture(gl.TEXTURE_2D, texture)
		// 	u.Context.Uniform1i(u.Context.GetUniformLocation(u.glProgram, "texture"), 0)
		// }

		// u.Context.UniformMatrix4fv(model_uni, model[:])
		// u.Context.UniformMatrix4fv(view_uni, view[:])
		// u.Context.UniformMatrix4fv(proj_uni, proj[:])

		// u.Context.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)

		// var model mgl32.Mat4
		model := mgl32.Ident4()
		trans := v.GetComponent(&render.Transform{})

		pos := trans.GetProperty("Position")
		switch pos := pos.(type) {
		case render.Vector3:
			scaledX := pos.X // float32(u.Platform.RenderW)
			scaledY := pos.Y // float32(u.Platform.RenderH) * u.Platform.AspectRatio
			model = model.Mul4(mgl32.Translate3D(scaledX, scaledY, pos.Z))

			scale := trans.GetProperty("Dimensions")
			switch scale := scale.(type) {
			case render.Vector2:
				scaledX = scale.X / 360 // float32(u.Platform.RenderW)
				scaledY = scale.Y / 640 // float32(u.Platform.RenderH)
				model = model.Mul4(mgl32.Scale3D(scaledX, scaledY, 0))

				text_arr := renderer.GetProperty("Text")
				switch text_arr := text_arr.(type) {
				case []render.Vector4:
					u.Context.Uniform4fv(text_uni, text_arr[0].ToUniformFloat())

					textBox := render.Vector4{pos.X /*x*/, pos.Y /*y*/, scale.X /*z*/, 640 - (pos.Y) /*w*/}

					text_scale := renderer.GetProperty("Scale")
					switch text_scale := text_scale.(type) {
					case float64:
						u.Context.Uniform1f(text_scale_uni, float32(text_scale))
					}

					text_padding := renderer.GetProperty("Padding")
					switch text_padding := text_padding.(type) {
					case render.Vector4:
						textBox[0] += text_padding[3]
						textBox[2] -= text_padding[3]
						textBox[2] -= text_padding[1]

						textBox[3] -= text_padding[0]
						// ONLY APPLIES IF ALIGNED TO BOTTOM
						textBox[1] -= text_padding[0]
						textBox[1] -= text_padding[2]
					}
					// THESE CAN BE float64?
					u.Context.Uniform4fv(box_uni, textBox.ToUniformFloat())
				}
			}
			u.Context.UniformMatrix4fv(model_uni, model[:])
			u.Context.UniformMatrix4fv(view_uni, view[:])
			u.Context.UniformMatrix4fv(proj_uni, proj[:])
		}

		u.Context.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE, 0)
	}
	u.Context.Flush()
}
func (r *UIGLES) AddUISystem(game *engine.Engine) {
}

func (u *UIGLES) Stop() {
	u.Context.DeleteProgram(u.glProgram)
}

func (u *UIGLES) LoadRenderer(renderer render.RendererRoutine) {
	renderer.LoadMesh(&render.Mesh{
		Vertices: []float32{
			1.0, 1.0, -0.2, 1.0, 0.0, 0.0,
			1.0, 0.0, -0.2, 0.0, 1.0, 0.0,
			0.0, 0.0, -0.2, 0.0, 0.0, 1.0,
			0.0, 1.0, -0.2, 0.0, 1.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	// Vertex buffer
	var vertexBuffer = u.Context.CreateBuffer()
	u.Context.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	u.Context.BufferData(gl.ARRAY_BUFFER, renderer.MeshByteVertices(), gl.STATIC_DRAW)
	renderer.SetProperty("VB", vertexBuffer)

	// Element buffer
	var elementBuffer = u.Context.CreateBuffer()
	u.Context.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	u.Context.BufferData(gl.ELEMENT_ARRAY_BUFFER, renderer.MeshByteIndicies(), gl.STATIC_DRAW)
	renderer.SetProperty("EB", elementBuffer)

	// Linking vertex attributes
	posAttrib := u.Context.GetAttribLocation(u.glProgram, "pos")
	u.Context.VertexAttribPointer(posAttrib, 3, gl.FLOAT, false, 6*4, 0)
	u.Context.EnableVertexAttribArray(posAttrib)

	// Linking fragment attributes
	colAttrib := u.Context.GetAttribLocation(u.glProgram, "col")
	u.Context.VertexAttribPointer(colAttrib, 3, gl.FLOAT, false, 6*4, 3*4)
	u.Context.EnableVertexAttribArray(colAttrib)

	// Linking texture attributes
	// texAttrib := u.Context.GetAttribLocation(u.glProgram, "tex")
	// u.Context.VertexAttribPointer(texAttrib, 2, gl.FLOAT, false, 8*4, 6*4)
	// u.Context.EnableVertexAttribArray(texAttrib)

	// texture := u.Context.CreateTexture()

	// u.Context.BindTexture(gl.TEXTURE_2D, texture)
	// u.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	// u.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	// u.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	// u.Context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	// u.Context.TexImage2D(
	// 	gl.TEXTURE_2D,
	// 	0,
	// 	int(renderer.TextureRGBA().Rect.Size().X),
	// 	int(renderer.TextureRGBA().Rect.Size().Y),
	// 	gl.RGBA,
	// 	gl.UNSIGNED_BYTE,
	// 	renderer.TextureRGBA().Pix)
	// renderer.SetProperty("TEXTURE", texture)
}

func (u *UIGLES) NewShader(vShader string, fShader string) uint32 {
	version := gl.Version()
	log.Printf("Render > UIGLES > Version: %v", version)
	// Create vertex shader
	vshader := u.Context.CreateShader(gl.VERTEX_SHADER)
	if vshader.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "Could not create VERTEXT_SHADER", vShader)
	}
	u.Context.ShaderSource(vshader, vShader)
	u.Context.CompileShader(vshader)
	defer u.Context.DeleteShader(vshader)
	if u.Context.GetShaderi(vshader, gl.COMPILE_STATUS) == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", u.Context.GetShaderInfoLog(vshader), vShader)
	}

	// Create fragment shader
	fshader := u.Context.CreateShader(gl.FRAGMENT_SHADER)
	if fshader.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "Could not create FRAGMENT_SHADER", fShader)
	}
	u.Context.ShaderSource(fshader, fShader)
	u.Context.CompileShader(fshader)
	defer u.Context.DeleteShader(fshader)
	if u.Context.GetShaderi(fshader, gl.COMPILE_STATUS) == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", u.Context.GetShaderInfoLog(fshader), fShader)
	}

	shaderProgram := u.Context.CreateProgram()
	u.glProgram = shaderProgram
	if shaderProgram.Value == 0 {
		log.Printf("\n\n ### SHADER ERROR ### \n%v\n%v\n\n", "No GLES program available")
	}

	u.Context.AttachShader(shaderProgram, vshader)
	u.Context.AttachShader(shaderProgram, fshader)

	u.Context.LinkProgram(shaderProgram)

	// Flag shaders for deletion when program is unlinked.
	u.Context.DeleteShader(vshader)
	u.Context.DeleteShader(fshader)

	if u.Context.GetProgrami(shaderProgram, gl.LINK_STATUS) == 0 {
		defer u.Context.DeleteProgram(shaderProgram)
		log.Printf("\n\n ### SHADER LINK ERROR ### \n%v\n\n", u.Context.GetProgramInfoLog(shaderProgram))
	}

	u.Context.UseProgram(shaderProgram)

	return shaderProgram.Value
}

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
