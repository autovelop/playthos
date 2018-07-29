package playthos_test

import (
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/opengl"
	_ "github.com/autovelop/playthos/platforms/linux"
	_ "github.com/autovelop/playthos/platforms/web"
	_ "github.com/autovelop/playthos/platforms/windows"
	"github.com/autovelop/playthos/profiling"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"testing"
	// "time"
)

func TestOpenGLTexture(t *testing.T) {
	profiling.StartProfiling(false, false)

	eng := engine.New("TestOpenGLTexture", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	newTextureObject := func(p *std.Vector3, s *std.Vector3, a float32, t string) {
		ent := eng.NewEntity()

		tra := std.NewTransform()
		tra.Set(p, &std.Vector3{0, 0, 0}, s)
		ent.AddComponent(tra)

		mat := render.NewMaterial()
		mat.Set(&std.Color{1, 1, 1, a})
		img := render.NewImage()
		img.LoadImage(t)
		tex := render.NewTexture(img)
		mat.SetTexture(tex)

		ent.AddComponent(mat)

		mes := render.NewMesh()
		mes.Set(std.QuadMesh)
		ent.AddComponent(mes)
	}
	// opaque background texture in center
	newTextureObject(&std.Vector3{0, 0, 5}, &std.Vector3{5, 5, 1}, 1, "../background.png")
	// texture moved and transparent
	newTextureObject(&std.Vector3{3, 0, 4}, &std.Vector3{1, 1, 1}, 1, "../texture.png")
	// texture in center
	newTextureObject(&std.Vector3{0, 0, 3}, &std.Vector3{1, 1, 1}, 1, "../texture.png")
	// texture moved, transparent, and overlapping
	newTextureObject(&std.Vector3{4, 1, 2}, &std.Vector3{1, 1, 1}, 1, "../texture.png")
	// texture moved, semi-transparent, and full coverage
	newTextureObject(&std.Vector3{-4, 4, 1}, &std.Vector3{2, 2, 1}, 1, "../window.png")

	eng.Start()
	profiling.ReportUPS(eng)
	profiling.StopProfiling()
}
