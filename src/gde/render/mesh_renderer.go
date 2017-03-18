package render

import (
	// "encoding/binary"
	"gde/engine"
	// "golang.org/x/mobile/exp/f32"
	"image"
	"log"
)

type MeshRenderer struct {
	engine.Component
	RendererRoutine
	Mesh    *Mesh
	Texture *Texture
	Color   *Color
}

func (r *MeshRenderer) Init() {
	// log.Printf("MeshRenderer > Init")
	r.Properties = make(map[string]interface{})
	r.Color = &Color{0, 0, 0, 1}
}

func (r *MeshRenderer) GetProperty(key string) interface{} {
	// log.Printf("MeshRenderer > Property > Get: %v", key)
	return r.Properties[key]
}

func (r *MeshRenderer) SetProperty(key string, val interface{}) {
	// log.Printf("MeshRenderer > Property > Set: %v", key)
	r.Properties[key] = val
}

func (r *MeshRenderer) LoadMesh(mesh *Mesh) {
	log.Printf("MeshRenderer > Mesh > Load: %v", mesh)
	r.Mesh = mesh
}

func (r *MeshRenderer) LoadTexture(texture *Texture) {
	// log.Printf("MeshRenderer > Texture > Load: %v", texture)
	r.Texture = texture
}

// Make this happen on the Render System
func (r *MeshRenderer) MeshVertices() []float32 {
	log.Printf("MeshRenderer > Mesh > Vectices: %v", len(r.Mesh.Vertices))
	return r.Mesh.Vertices
}

func (r *MeshRenderer) MeshIndicies() []uint8 {
	log.Printf("MeshRenderer > Mesh > Indicies: %v", len(r.Mesh.Indicies))
	return r.Mesh.Indicies
}

// split this to opengles package
// func (r *MeshRenderer) MeshByteVertices() []byte {
// log.Printf("MeshRenderer > Mesh > Vectices: %v", len(r.Mesh.Vertices))
// return f32.Bytes(binary.LittleEndian, r.mesh.Vertices...)
// }

// func (r *MeshRenderer) MeshByteIndicies() []byte {
// 	log.Printf("MeshRenderer > Mesh > Indicies: %v", len(r.Mesh.Indicies))
// 	return r.Mesh.Indicies
// }

func (r *MeshRenderer) SetColor(color *Color) {
	log.Printf("MeshRenderer > SetColor: %T", color)
	r.Color = color
}
func (r *MeshRenderer) GetColor() *Color {
	log.Printf("MeshRenderer > GetColor: %T", r.Color)
	return r.Color
}

func (r *MeshRenderer) HasTexture() bool {
	log.Printf("MeshRenderer > HasTexture: %v", r.Texture != nil)
	return r.Texture != nil
}

func (r *MeshRenderer) GetTextureRGBA() *image.RGBA {
	log.Printf("MeshRenderer > Texture > RGBA: %T", r.Texture.GetRGBA())
	return r.Texture.GetRGBA()
}
