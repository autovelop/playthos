// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
)

type Image struct {
	id   uint32
	path string
	// filePath string
	// rgba []byte
	// rgba        *image.RGBA
	width       int32
	height      int32
	aspectRatio float32
}

func NewImage() *Image {
	return &Image{}
}

func (i *Image) SetID(id uint32) {
	i.id = id
}

func (i *Image) ID() uint32 {
	return i.id
}

func (i *Image) Path() string {
	return i.path
}

func (i *Image) Width() int32 {
	return i.width
}

func (i *Image) Height() int32 {
	return i.height
}

func (i *Image) SetWidth(w int32) {
	i.width = w
}

func (i *Image) SetHeight(h int32) {
	i.height = h
}

func (i *Image) LoadImage(p string) {
	i.path = p
	engine.LoadAsset(p)
}
