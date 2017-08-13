// +build autovelop_playthos_render !play

package render

type Texture struct {
	image *Image
}

func NewTexture(i *Image) *Texture {
	return &Texture{i}
}

func (t *Texture) SetID(id uint32) {
	t.image.SetID(id)
}

func (t *Texture) ID() uint32 {
	return t.image.ID()
}

func (t *Texture) RGBA() []byte {
	return t.image.RGBA()
}

func (t *Texture) Width() int32 {
	return t.image.Width
}

func (t *Texture) Height() int32 {
	return t.image.Height
}
