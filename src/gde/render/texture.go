package render

import (
	"go/build"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/mobile/asset"
)

type Texture struct {
	filePath    string
	RGBA        *image.RGBA
	Width       int
	Height      int
	AspectRatio float32
}

// TODO. OCD clean this fanus!

func (t *Texture) NewTexture(dir string, path string) bool {
	t.filePath = path

	dir, err := importPathToDir(dir)
	if err != nil {
		log.Println("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
		return false
	}
	err = os.Chdir(dir)
	if err != nil {
		log.Println("os.Chdir:", err)
		return false
	}

	imgFile, err := os.Open(t.filePath)
	if err != nil {
		log.Println(err)
		return false
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println(err)
		return false
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Println("rgba stride error")
		return false
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
	return true
}

func (t *Texture) NewTextureMobile(path string) bool {
	a, err := asset.Open(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// log.Printf("\n%+v\n\n%+v", a, m)
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Println("rgba stride error")
		return false
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
	return true
}

func importPathToDir(importPath string) (string, error) {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, nil
}
