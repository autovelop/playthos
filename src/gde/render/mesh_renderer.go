package render

import (
	"gde/engine"
	"image"
	"log"
)

type MeshRenderer struct {
	engine.Component
	RendererRoutine
	Mesh    *Mesh
	Texture *Texture
}

func (r *MeshRenderer) Init() {
	log.Printf("MeshRenderer > Init")
	r.Properties = make(map[string]interface{})
}

func (r *MeshRenderer) GetProperty(key string) interface{} {
	log.Printf("MeshRenderer > Property > Get: %v", key)
	return r.Properties[key]
}

func (r *MeshRenderer) SetProperty(key string, val interface{}) {
	log.Printf("MeshRenderer > Property > Set: %v", key)
	r.Properties[key] = val
}

func (r *MeshRenderer) LoadMesh(mesh *Mesh) {
	log.Printf("MeshRenderer > Mesh > Load: %v", mesh)
	r.Mesh = mesh
}

func (r *MeshRenderer) LoadTexture(texture *Texture) {
	log.Printf("MeshRenderer > Texture > Load: %v", texture)
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

func (r *MeshRenderer) TextureRGBA() *image.RGBA {
	log.Printf("MeshRenderer > Texture > RGBA: %T", r.Texture.RGBA)
	return r.Texture.RGBA
}
