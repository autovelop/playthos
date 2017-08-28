// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Material struct {
	engine.Component
	color       *std.Color
	baseTexture Textureable
	SetTexture  func(Textureable)
}

func NewMaterial() *Material {
	m := &Material{}
	m.SetTexture = func(t Textureable) {
		m.baseTexture = t
	}
	return m
}

func (m *Material) BaseTexture() Textureable {
	return m.baseTexture
}

func (m *Material) Set(t Textureable, col *std.Color) {
	m.baseTexture = t
	m.color = col
}

func (m *Material) SetColor(col *std.Color) {
	m.color = col
}

func (m *Material) Color() *std.Color {
	return m.color
}

// func (m *Material) setTexture(texture *Texture) {
// 	m.textureBase = texture
// }

// func (m *Material) Texture() *Texture {
// 	return m.textureBase
// }

// func (m *Material) SetSprite(s *Sprite) {
// 	m.sprite = s
// }

// func (m *Material) Sprite() *Sprite {
// 	return m.sprite
// }
