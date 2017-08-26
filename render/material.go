// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Material struct {
	engine.Component
	color       *std.Color
	textureBase *Texture
	// sprite  *Sprite
	SetTexture func(*Texture)
}

func NewMaterial() *Material {
	m := &Material{}
	m.SetTexture = m.setTexture
	return m
}

func (m *Material) Set(texture *Texture, col *std.Color) {
	m.textureBase = texture
	m.color = col
}

func (m *Material) SetColor(col *std.Color) {
	m.color = col
}

func (m *Material) Color() *std.Color {
	return m.color
}

func (m *Material) setTexture(texture *Texture) {
	m.textureBase = texture
}

func (m *Material) Texture() *Texture {
	return m.textureBase
}

// func (m *Material) SetSprite(s *Sprite) {
// 	m.sprite = s
// }

// func (m *Material) Sprite() *Sprite {
// 	return m.sprite
// }
