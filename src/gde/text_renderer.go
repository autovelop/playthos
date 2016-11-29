package gde

import (
	"fmt"
)

// here what to do
// implement http://glslsandbox.com/e#24076.0 shader
// create a shader program for the textrenderer
// send an uniform array of the vec2 macros defined in shader
// calc things like new lines etc.
type TextRenderer struct {
	Component

	Mesh *Mesh
	Text *Text
}

func (r *TextRenderer) Init() {
	fmt.Println("TextRenderer.Init() executed")
	r.Properties = make(map[string]interface{})
}

// Make this happen on the Render System
func (r *TextRenderer) LoadMesh(mesh *Mesh) {
	r.Mesh = mesh
}

func (r *TextRenderer) LoadText(text *Text) {
	r.Text = text
}

func (r *TextRenderer) GetProperty(key string) interface{} {
	return r.Properties[key]
}

func (r *TextRenderer) SetProperty(key string, val interface{}) {
	r.Properties[key] = val
}
