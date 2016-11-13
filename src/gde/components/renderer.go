package components

import (
	"fmt"

	"gde"
	geometry "gde/geometry"
	// "github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	gde.Component
	// gde.ComponentRoutine

	Mesh *geometry.Mesh
}

func (r *Renderer) Init() {
	fmt.Println("Renderer.Init() executed")
	r.Properties = make(map[string]interface{})
}

// Make this happen on the Render System
func (r *Renderer) LoadMesh(mesh *geometry.Mesh) {
	r.Mesh = mesh
	// Bind vertex array object. This must wrap around the mesh creation because it is how we are going to access it later when we draw
	// var vertexArrayID uint32
	// gl.GenVertexArrays(1, &vertexArrayID)
	// gl.BindVertexArray(vertexArrayID)

	// // Vertex buffer
	// var vertexBuffer uint32
	// gl.GenBuffers(1, &vertexBuffer)
	// gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	// gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Vertices)*4, gl.Ptr(mesh.Vertices), gl.STATIC_DRAW)

	// // Element buffer
	// var elementBuffer uint32
	// gl.GenBuffers(1, &elementBuffer)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBuffer)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.Indicies)*4, gl.Ptr(mesh.Indicies), gl.STATIC_DRAW)

	// // Linking vertex attributes
	// gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	// gl.EnableVertexAttribArray(0)

	// // Unbind Vertex array object
	// r.SetProperty("VAO", vertexArrayID)
	// gl.BindVertexArray(0)
}

func (r *Renderer) GetProperty(key string) interface{} {
	return r.Properties[key]
}

func (r *Renderer) SetProperty(key string, val interface{}) {
	r.Properties[key] = val
}
