package geometry

import (
	"encoding/binary"
	"golang.org/x/mobile/exp/f32"
	"log"
)

type Mesh struct {
	Vertices []float32
	Indicies []uint8
}

func (m *Mesh) VerticesByteArray() []byte {
	// log.Printf("%v", len(m.Vertices))
	return f32.Bytes(binary.LittleEndian, m.Vertices...)
}

func (m *Mesh) IndiciesByteArray() []byte {
	log.Printf("%v", len(m.Indicies))
	return m.Indicies
	// return f32.Bytes(binary.LittleEndian, m.Indicies...)
}
