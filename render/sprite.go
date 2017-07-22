// +build render

package render

import (
	"github.com/autovelop/playthos/std"
)

type Sprite struct {
	image      *Image
	spriteSize *std.Vector2
	spriteX    *std.Integer
	spriteY    *std.Integer
}

func NewSprite(i *Image) *Sprite {
	return &Sprite{i, &std.Vector2{float32(i.Width), float32(i.Height)}, &std.Integer{0}, &std.Integer{0}}
}

func (s *Sprite) SetID(id uint32) {
	s.image.SetID(id)
}

func (s *Sprite) ID() uint32 {
	return s.image.ID()
}

func (s *Sprite) RGBA() []byte {
	return s.image.RGBA()
}

func (s *Sprite) Width() int32 {
	return s.image.Width
}

func (s *Sprite) Height() int32 {
	return s.image.Height
}

func (s *Sprite) SetSpriteSize(x float32, y float32) {
	s.spriteSize = &std.Vector2{x, y}
}

func (s *Sprite) SetSprite(x *std.Integer, y *std.Integer) {
	s.spriteX = x
	s.spriteY = y
	// s.offset = r
}

// func (s *Sprite) SetOffset(r *std.Vector2) {
// 	s.offset = r
// }

func (s *Sprite) Size() *std.Vector2 {
	return s.spriteSize
}

func (s *Sprite) Offset() *std.Vector2 {
	v := &std.Vector2{float32(s.spriteX.V) * s.spriteSize.X, float32(s.spriteY.V) * s.spriteSize.Y}
	return v
}
