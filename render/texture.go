// +build deploy render

package render

import (
	"github.com/autovelop/playthos/std"
)

// Texture defines the image, size, and offset
type Texture struct {
	image  *Image
	size   *std.Vector2
	offset *std.Vector2
}

// NewTexture creates and sets a new orphan texture
func NewTexture(i *Image) *Texture {
	return &Texture{i, &std.Vector2{0, 0}, &std.Vector2{0, 0}}
}

// Path returns path string of image
func (t *Texture) Path() string {
	return t.image.Path()
}

// Width returns width of image
func (t *Texture) Width() int32 {
	return t.image.Width()
}

// Height returns height of image
func (t *Texture) Height() int32 {
	return t.image.Height()
}

// SetWidth sets/changes image width
func (t *Texture) SetWidth(w int32) {
	t.image.SetWidth(w)
}

// SetHeight sets/changes image height
func (t *Texture) SetHeight(h int32) {
	t.image.SetHeight(h)
}

// SetSize sets/changes texture size
func (t *Texture) SetSize(x float32, y float32) {
	t.size = &std.Vector2{x, y}
}

// SetOffset sets/changes texture offset vector
func (t *Texture) SetOffset(o *std.Vector2) {
	t.offset = o
}

// Size returns size vector
func (t *Texture) Size() *std.Vector2 {
	return t.size
}

// SizeN returns size vector of texture relative to image dimensions
func (t *Texture) SizeN() *std.Vector2 {
	return &std.Vector2{t.size.X / float32(t.image.width), t.size.Y / float32(t.image.height)}
}

// Offset returns offset vector
func (t *Texture) Offset() *std.Vector2 {
	return t.offset
}

// Textureable interface allows for platform specific texture data
type Textureable interface {
	Path() string
	Width() int32
	Height() int32
	SetWidth(int32)
	SetHeight(int32)
}
