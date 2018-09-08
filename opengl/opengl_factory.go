// +build deploy opengl

package opengl

import (
	// "fmt"
	// "github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
)

// OpenGLMesh defines a mesh (opengl)
type OpenGLMesh struct {
	m   *render.Mesh
	vao uint32
}

// SetVAO sets the VAO (opengl)
func (m *OpenGLMesh) SetVAO(vao uint32) {
	m.vao = vao
}

// VAO returns a opengl VAO
func (m *OpenGLMesh) VAO() uint32 {
	return m.vao
}

// OpenGLTexture defines a texture (opengl)
type OpenGLTexture struct {
	*render.Texture
	id uint32
}

// ID returns a opengl texture id
func (t *OpenGLTexture) ID() uint32 {
	return t.id
}

// OpenGLMaterial defines a material (opengl)
type OpenGLMaterial struct {
	*render.Material
	texture *OpenGLTexture
}

// NewOpenGLMaterial creates a meterial (opengl)
func NewOpenGLMaterial(m *render.Material) *OpenGLMaterial {
	openGLMaterial := &OpenGLMaterial{Material: m}
	return openGLMaterial
}

// OverrideTexture overrides base texture (opengl)
func (o *OpenGLMaterial) OverrideTexture(fn func(render.Textureable)) {
	o.SetTexture = fn
	baseTexture := o.BaseTexture()
	if baseTexture != nil {
		o.SetTexture(baseTexture.(*render.Texture))
	}
}

// Texture returns a opengl texture
func (o *OpenGLMaterial) Texture() *OpenGLTexture {
	return o.texture
}

// ID returns a opengl material id
func (o *OpenGLMaterial) ID() uint32 {
	return o.texture.id
}
