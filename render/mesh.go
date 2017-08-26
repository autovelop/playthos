// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Mesh struct {
	engine.Component
	// vao      uint32
	vertices []float32
	indicies []uint8
}

func NewMesh() *Mesh {
	return &Mesh{}
}

func (m *Mesh) Set(veb *std.VEB) {
	m.vertices = veb.VB
	m.indicies = veb.EB
}

func (m *Mesh) Vertices() []float32 {
	return m.vertices
}

func (m *Mesh) Indicies() []uint8 {
	return m.indicies
}

type Meshable interface {
	Set(veb *std.VEB)
	Vertices() []float32
	Indicies() []uint8
}
