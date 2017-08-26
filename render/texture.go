// +build deploy render

package render

import (
	"github.com/autovelop/playthos/std"
)

type Texture struct {
	image  *Image
	size   *std.Vector2
	offset *std.Vector2
}

func NewTexture(i *Image) *Texture {
	return &Texture{i, &std.Vector2{0, 0}, &std.Vector2{0, 0}}
}

func (t *Texture) Path() string {
	return t.image.Path()
}

func (t *Texture) Width() int32 {
	return t.image.Width()
}

func (t *Texture) Height() int32 {
	return t.image.Height()
}

func (t *Texture) SetWidth(w int32) {
	t.image.SetWidth(w)
}

func (t *Texture) SetHeight(h int32) {
	t.image.SetHeight(h)
}

func (t *Texture) SetSize(x float32, y float32) {
	t.size = &std.Vector2{x, y}
}

func (t *Texture) SetOffset(o *std.Vector2) {
	t.offset = o
}

func (t *Texture) Size() *std.Vector2 {
	return t.size
}

func (t *Texture) SizeN() *std.Vector2 {
	return &std.Vector2{t.size.X / float32(t.image.width), t.size.Y / float32(t.image.height)}
}

func (t *Texture) Offset() *std.Vector2 {
	return t.offset
}
