// +build render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Material struct {
	engine.Component
	color   *std.Color
	texture *Texture
}

func (m *Material) SetColor(col *std.Color) {
	m.color = col
}

func (m *Material) GetColor() *std.Color {
	return m.color
}

func (m *Material) SetTexture(texture *Texture) {
	m.texture = texture
}

func (m *Material) GetTexture() *Texture {
	return m.texture
}
