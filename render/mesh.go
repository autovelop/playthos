package render

import (
	"github.com/autovelop/golang-gde"
	"log"
)

type Mesh struct {
	vao      uint32
	vertices []float32
	indicies []uint8
}

func (m *Mesh) RegisterToSystem(system engine.System) {
	log.Println("Registering Mesh")
	switch system := system.(type) {
	case Render:
		system.RegisterMesh(m)
	}
}

func (m *Mesh) SetVAO(vao uint32) {
	m.vao = vao
}

func (m *Mesh) GetVAO() uint32 {
	return m.vao
}

func (m *Mesh) Set(verts []float32, inds []uint8) {
	m.vertices = verts
	m.indicies = inds
}

func (m *Mesh) GetVertices() []float32 {
	return m.vertices
}

func (m *Mesh) GetIndicies() []uint8 {
	return m.indicies
}
