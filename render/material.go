package render

import (
	"./../engine"
	"log"
)

type Material struct {
	color *Color
}

func (m *Material) RegisterToSystem(system engine.System) {
	log.Println("Registering Material")
	switch system := system.(type) {
	case Render:
		system.RegisterMaterial(m)
	}
}

func (m *Material) Set(col *Color) {
	m.color = col
}

func (m *Material) GetColor() *Color {
	return m.color
}
