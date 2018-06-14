// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Material defines the color and texture properties
type Material struct {
	engine.Component
	color       *std.Color
	baseTexture Textureable
	SetTexture  func(Textureable)
}

// NewMaterial creates and sets a new orphan material
func NewMaterial() *Material {
	m := &Material{}
	m.SetTexture = func(t Textureable) {
		m.baseTexture = t
	}
	return m
}

// BaseTexture returns base texture interface
func (m *Material) BaseTexture() Textureable {
	return m.baseTexture
}

// SetTexture sets/changes texture
// func (m *Material) SetTexture(t Textureable) {
// 	m.baseTexture = t
// }

// Set used to define all the required properties
func (m *Material) Set(col *std.Color) {
	m.color = col
}

// Color returns the material color
func (m *Material) Color() *std.Color {
	return m.color
}
