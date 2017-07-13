// +build render

package render

import (
	"github.com/autovelop/playthos"
	"go/build"
	"image"
	"image/draw"
	// "image/color"
	"bytes"
	_ "image/png"
	"log"
	// "os"
)

type Texture struct {
	id       uint32
	filePath string
	rgba     []byte
	// rgba        *image.RGBA
	Width       int32
	Height      int32
	AspectRatio float32
}

func NewTexture() *Texture {
	return &Texture{}
}

func (t *Texture) SetID(id uint32) {
	t.id = id
}

func (t *Texture) ID() uint32 {
	return t.id
}

func (t *Texture) RGBA() []byte {
	return t.rgba
}

func (t *Texture) NewTexture(dir string, path string) bool {
	t.filePath = path

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

	t.Width, t.Height = int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y)

	buffer := rgba.Pix
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	t.rgba = buffer
	return true
}

// func (t *Texture) NewTextureMobile(path string) bool {
// 	a, err := asset.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	defer a.Close()

// 	img, _, err := image.Decode(a)
// 	if err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	// log.Printf("\n%+v\n\n%+v", a, m)
// 	rgba := image.NewRGBA(img.Bounds())
// 	if rgba.Stride != rgba.Rect.Size().X*4 {
// 		log.Println("rgba stride error")
// 		return false
// 	}

// 	t.Width = rgba.Rect.Size().X
// 	t.Height = rgba.Rect.Size().Y
// 	if t.Width > t.Height {
// 		t.AspectRatio = float32(t.Width) / float32(t.Height)
// 	} else {
// 		t.AspectRatio = float32(t.Height) / float32(t.Width)
// 	}

// 	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
// 	t.rgba = rgba
// 	return true
// }

func importPathToDir(importPath string) (string, error) {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, nil
}
