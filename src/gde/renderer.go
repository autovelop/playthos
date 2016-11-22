package gde

import (
	"fmt"
)

type Renderer struct {
	Component
	// gde.ComponentRoutine

	Mesh    *Mesh
	Texture *Texture
}

func (r *Renderer) Init() {
	fmt.Println("Renderer.Init() executed")
	r.Properties = make(map[string]interface{})
}

// Make this happen on the Render System
func (r *Renderer) LoadMesh(mesh *Mesh) {
	r.Mesh = mesh
}

func (r *Renderer) LoadTexture(texture *Texture) {
	r.Texture = texture
}

func (r *Renderer) GetProperty(key string) interface{} {
	return r.Properties[key]
}

func (r *Renderer) SetProperty(key string, val interface{}) {
	r.Properties[key] = val
}
