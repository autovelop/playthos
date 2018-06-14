package playthos_test

import (
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/opengl"
	"github.com/autovelop/playthos/profiling"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"testing"
	"time"
)

func TestOpenGLSquare(t *testing.T) {
	profiling.StartProfiling(false, false)

	eng := engine.New("TestOpenGLSquare", &engine.Settings{
		false,
		1024,
		576, // 16:9
		false,
	})

	newSquare := func(p *std.Vector3, r *std.Vector3, s *std.Vector3, c *std.Color) {
		ent := eng.NewEntity()

		tra := std.NewTransform()
		tra.Set(p, r, s)
		ent.AddComponent(tra)

		mat := render.NewMaterial()
		mat.Set(c)
		ent.AddComponent(mat)

		mes := render.NewMesh()
		mes.Set(std.QuadMesh)
		ent.AddComponent(mes)
	}
	// square moved and rotated
	newSquare(&std.Vector3{4, 0, 0}, &std.Vector3{0, 0, 45}, &std.Vector3{1, 1, 1}, &std.Color{0, 1, 1, 1})
	// square moved and scaled
	newSquare(&std.Vector3{-4, 0, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{2, 2, 1}, &std.Color{0, 1, 0, 1})
	// square moved and sent to front (BLUE)
	newSquare(&std.Vector3{1, 0.6, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.5, 1.5, 1}, &std.Color{0, 0, 1, 0.25})
	// square in center (YELLOW)
	newSquare(&std.Vector3{0, 0, 1}, &std.Vector3{0, 0, 0}, &std.Vector3{1.1, 1.1, 1}, &std.Color{1, 1, 0, 0.25})
	// square moved and sent to back (RED)
	newSquare(&std.Vector3{-0.6, -0.5, 2}, &std.Vector3{0, 0, 0}, &std.Vector3{0.75, 0.75, 1}, &std.Color{1, 0, 0, 0.25})
	// WTF!!!!!!!!!!!!!!!!!!  newSquare(&std.Vector3{0.5, 0.5, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1, 1, 1}, &std.Color{0, 0, 1, 0.75})

	go func(e *engine.Engine) {
		time.Sleep(5 * time.Second)
		e.Stop()
	}(eng)

	eng.Start()
	profiling.ReportUPS(eng)
	profiling.StopProfiling()
}
