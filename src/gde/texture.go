package gde

import (
	"fmt"
	"go/build"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
)

type Texture struct {
	filePath    string
	RGBA        *image.RGBA
	Width       int
	Height      int
	AspectRatio float32
}

func (t *Texture) NewTexture(path string) {
	t.filePath = path

	dir, err := importPathToDir("gde/resources")
	if err != nil {
		log.Fatalln("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
	}
	err = os.Chdir(dir)
	if err != nil {
		log.Panicln("os.Chdir:", err)
	}

	imgFile, err := os.Open(t.filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Println("rgba stride error")
		return
	}

	t.Width = rgba.Rect.Size().X
	t.Height = rgba.Rect.Size().Y
	if t.Width > t.Height {
		t.AspectRatio = float32(t.Width) / float32(t.Height)
	} else {
		t.AspectRatio = float32(t.Height) / float32(t.Width)
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	t.RGBA = rgba
}
func importPathToDir(importPath string) (string, error) {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, nil
}
