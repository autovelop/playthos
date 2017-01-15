package render

import (
	"gde/engine"
	"image"
	// "log"
)

// type Renderer struct {
// 	*engine.Component

// 	Mesh    *Mesh
// 	Texture *Texture
// }

type RendererRoutine interface {
	engine.ComponentRoutine

	LoadMesh(mesh *Mesh)
	LoadTexture(texture *Texture)

	MeshVertices() []float32
	MeshIndicies() []uint8

	// Used for OpenGLES
	MeshByteVertices() []byte
	MeshByteIndicies() []byte

	TextureRGBA() *image.RGBA
}
