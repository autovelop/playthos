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

// Set used to define all the required properties
func (m *Material) Set(t Textureable, col *std.Color) {
	m.baseTexture = t
	m.color = col
}

// SetColor sets/changes material color
func (m *Material) SetColor(col *std.Color) {
	m.color = col
}

// Color returns the material color
func (m *Material) Color() *std.Color {
	return m.color
}
