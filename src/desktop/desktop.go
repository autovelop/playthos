package main

import (
<<<<<<< HEAD
	"fmt"
	"log"
	"runtime"

	"gde"
	components "gde/components"
	systems "gde/systems"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(360, 640, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Greate game engine
	engine := &gde.Engine{}
	engine.Init()

	// Create render system
	var render gde.SystemRoutine
	render = &systems.Render{}
	render.Init()
	render.Add(engine)

	// Create player entity
	var player gde.EntityRoutine
	player = &gde.Entity{Id: "Player"}
	player.Init()
	player.Add(engine)

	// Create transform component
	var transformComponent gde.ComponentRoutine
	transformComponent = &components.Transform{}
	transformComponent.Add(player.Get())

	fmt.Println(player.Get().Id)

	// var renderRoutine systems.RenderRoutine
	// renderRoutine = render.Get()
	// fmt.Println(render.VertexShader)

	// stop scaling and start integrating with opengl

	// ########################################
	// ########################################
	// ########################################
	// ########################################
	// ########################################
	// ########################################

	// render.SetProperty("OpenGL", gl.GoStr(gl.GetString(gl.VERSION)))

	// Create Vertex array object
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	var vertexBufferData = []float32{
		-0.5, 0.5,
		0.5, 0.5,
		0.5, -0.5,
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexBufferData)*4, gl.Ptr(vertexBufferData), gl.STATIC_DRAW)

	// Create program
	program := gl.CreateProgram()
	gl.UseProgram(program)

	// Create vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vsources, vfree := gl.Strs(render.Property("VertexShader").(string))
	gl.ShaderSource(vshader, 1, vsources, nil)
	vfree()
	gl.CompileShader(vshader)
	gl.AttachShader(program, vshader)

	// Create fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsources, ffree := gl.Strs(render.Property("FragmentShader").(string))
	gl.ShaderSource(fshader, 1, fsources, nil)
	ffree()
	gl.CompileShader(fshader)
	gl.AttachShader(program, fshader)

	// Link program
	gl.LinkProgram(program)

	// Delete all shaders
	gl.DeleteShader(vshader)
	gl.DeleteShader(fshader)

	// Start game loop
	for !window.ShouldClose() {
		gl.ClearColor(1, 1, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program)

		// engine.Update()

		gl.EnableVertexAttribArray(0)
		gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DisableVertexAttribArray(0)

		window.SwapBuffers()
		glfw.PollEvents()
	}
=======
  "fmt"
  "runtime"
  "log"
  "shared"
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
  runtime.LockOSThread()
}

func main() {
  if err := glfw.Init(); err != nil {
    log.Fatalln("failed to initialize glfw:", err)
  }
  defer glfw.Terminate()

  window, err := glfw.CreateWindow(360, 640, "Cube", nil, nil)
  if err != nil {
    panic(err)
  }
  window.MakeContextCurrent()

  // Initialize Glow
  if err := gl.Init(); err != nil {
    panic(err)
  }

  version := gl.GoStr(gl.GetString(gl.VERSION))
  fmt.Println("OpenGL version", version)


  // Create Vertex array object
  var vertexArrayID uint32
  gl.GenVertexArrays(1, &vertexArrayID)
  gl.BindVertexArray(vertexArrayID)

  var vertexBuffer uint32
  gl.GenBuffers(1, &vertexBuffer)
  var vertexBufferData = []float32{
    -0.5,  0.5,
    0.5, 0.5,
    0.5, -0.5,
  }
  gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
  gl.BufferData(gl.ARRAY_BUFFER, len(vertexBufferData)*4, gl.Ptr(vertexBufferData), gl.STATIC_DRAW)

  // Create program
  program := gl.CreateProgram()
  gl.UseProgram(program)

  // Create vertex shader
  vshader := gl.CreateShader(gl.VERTEX_SHADER)
  vsources, vfree := gl.Strs(shared.VSHADER_OPENGL_4_1)
  gl.ShaderSource(vshader, 1, vsources, nil)
  vfree()
  gl.CompileShader(vshader)
  gl.AttachShader(program, vshader)


  // Create fragment shader
  fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
  fsources, ffree := gl.Strs(shared.FSHADER_OPENGL_4_1)
  gl.ShaderSource(fshader, 1, fsources, nil)
  ffree()
  gl.CompileShader(fshader)
  gl.AttachShader(program, fshader)

  // Link program
  gl.LinkProgram(program)

  // Delete all shaders
  gl.DeleteShader(vshader)
  gl.DeleteShader(fshader)


  // Start game loop
  for !window.ShouldClose() {
    gl.ClearColor(1, 0, 0, 1)
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    gl.UseProgram(program)

    gl.EnableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)

    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
    gl.DrawArrays(gl.TRIANGLES, 0, 3)
    gl.DisableVertexAttribArray(0)

    window.SwapBuffers()
    glfw.PollEvents()
  }

>>>>>>> 2175e2767f1fb67c90f0f66e12fdb82ff9fcda7c
}
