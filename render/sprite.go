// +build autovelop_playthos_render !play

package render

import (
	"github.com/autovelop/playthos/std"
)

type Sprite struct {
	image        *Image
	spriteSize   *std.Vector2
	spriteOffset *std.Vector2
}

func NewSprite(i *Image) *Sprite {
	return &Sprite{i, &std.Vector2{float32(i.Width), float32(i.Height)}, &std.Vector2{0, 0}}
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

func (s *Sprite) SetSpriteOffset(v *std.Vector2) {
	s.spriteOffset = v
}

func (s *Sprite) Size() *std.Vector2 {
	return s.spriteSize
}

func (s *Sprite) SizeN() *std.Vector2 {
	return &std.Vector2{s.spriteSize.X / float32(s.Width()), s.spriteSize.Y / float32(s.Height())}
}

func (s *Sprite) Offset() *std.Vector2 {
	return s.spriteOffset
}
