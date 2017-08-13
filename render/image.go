// +build autovelop_playthos_render !play

package render

import (
	"github.com/autovelop/playthos"
	// "go/build"
	"image"
	"image/draw"
	// "image/color"
	"bytes"
	_ "image/png"
	"log"
	// "os"
)

type Image struct {
	id       uint32
	filePath string
	rgba     []byte
	// rgba        *image.RGBA
	Width       int32
	Height      int32
	AspectRatio float32
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

func (i *Image) RGBA() []byte {
	return i.rgba
}

func (i *Image) LoadImage(dir string, path string) bool {
	i.filePath = path

	buf, err := engine.LoadAsset(dir, path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		log.Println(err)
		return false
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Println("rgba stride error")
		return false
	}

	i.Width, i.Height = int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y)

	buffer := rgba.Pix
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	i.rgba = buffer
	return true
}
