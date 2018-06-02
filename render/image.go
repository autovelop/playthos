// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
)

// Image defines the dimension data of an image stored on an operating system or in a binary file
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

// NewImage creates and set a new orphan image
func NewImage() *Image {
	return &Image{}
}

// SetID sets/changes the id
func (i *Image) SetID(id uint32) {
	i.id = id
}

// ID returns image id
func (i *Image) ID() uint32 {
	return i.id
}

// Path returns image path string
func (i *Image) Path() string {
	return i.path
}

// Width returns image width
func (i *Image) Width() int32 {
	return i.width
}

// Height returns image height
func (i *Image) Height() int32 {
	return i.height
}

// SetWidth sets/changes image width value
func (i *Image) SetWidth(w int32) {
	i.width = w
}

// SetHeight sets/changes image height value
func (i *Image) SetHeight(h int32) {
	i.height = h
}

// LoadImage instructs engine to load iamge from path
func (i *Image) LoadImage(p string) {
	i.path = p
	engine.LoadAsset(p)
}
