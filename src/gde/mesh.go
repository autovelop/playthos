package gde

import (
	"encoding/binary"
	"golang.org/x/mobile/exp/f32"
)

type Mesh struct {
	Vertices []float32
	Indicies []uint8
}

func (m *Mesh) VerticesByteArray() []byte {
	return f32.Bytes(binary.LittleEndian, m.Vertices...)
}

func (m *Mesh) IndiciesByteArray() []byte {
	return m.Indicies
}
