package shared

const VSHADER_OPENGL_4_1 = `#version 410
in vec2 position;
void main()
{
  gl_Position = vec4(position, 0.0, 1.0);
}
` + "\x00"

const FSHADER_OPENGL_4_1 = `#version 410
out vec4 outColor;
void main()
{
  outColor = vec4(1.0, 1.0, 1.0, 1.0);
}
` + "\x00"

const VSHADER_OPENGL_ES_2_0 = `#version 100

attribute vec4 position;
void main() {
  gl_Position = position;
}`

const FSHADER_OPENGL_ES_2_0 = `#version 100
precision mediump float;
void main() {
  gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
}`


// For now we will rather only support OPENGL_ES_2_0
// const VSHADER_OPENGL_ES_3_0 = `#version 310 es
// in vec2 position;
// void main()
// {
//     gl_Position = vec4(position, 0.0, 1.0);
// }
// `
// 
// const FSHADER_OPENGL_ES_3_0 = `#version 310 es
// precision mediump float;
// out vec4 outColor;
// void main()
// {
//     outColor = vec4(1.0, 1.0, 1.0, 1.0);
// }
// `


























const VertexShaderMob = `
#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
  // offset comes in with x/y values between 0 and 1.
  // position bounds are -1 to 1.
  vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
  gl_Position = position + offset4;
}`

const FragmentShaderMob = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
  gl_FragColor = color;
}
`
