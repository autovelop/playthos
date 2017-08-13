// +build autovelop_playthos_render !play

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Material struct {
	engine.Component
	color   *std.Color
	texture *Texture
	sprite  *Sprite
}

func NewMaterial() *Material {
	return &Material{}
}

func (m *Material) Set(texture *Texture, col *std.Color) {
	m.texture = texture
	m.color = col
}

func (m *Material) SetColor(col *std.Color) {
	m.color = col
}

func (m *Material) Color() *std.Color {
	return m.color
}

func (m *Material) SetTexture(texture *Texture) {
	m.texture = texture
}

func (m *Material) Texture() *Texture {
	return m.texture
}

func (m *Material) SetSprite(s *Sprite) {
	m.sprite = s
}

func (m *Material) Sprite() *Sprite {
	return m.sprite
}
