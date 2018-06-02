// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Mesh defines the vertex and index slices of a mesh
type Mesh struct {
	engine.Component
	// vao      uint32
	vertices []float32
	indicies []uint8
}

// NewMesh creates and set a new orphan material
func NewMesh() *Mesh {
	return &Mesh{}
}

// Set used to define all the required properties
func (m *Mesh) Set(veb *std.VEB) {
	m.vertices = veb.VB
	m.indicies = veb.EB
}

// Vertices returns vertex slice
func (m *Mesh) Vertices() []float32 {
	return m.vertices
}

// Indicies returns index slice
func (m *Mesh) Indicies() []uint8 {
	return m.indicies
}

// // Meshable interface allows
// type Meshable interface {
// 	Set(veb *std.VEB)
// 	Vertices() []float32
// 	Indicies() []uint8
// }
