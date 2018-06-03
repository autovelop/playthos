package playthos_test

import (
	"github.com/autovelop/playthos"
	// _ "github.com/autovelop/playthos/opengl"
	"github.com/autovelop/playthos/profiling"
	// "github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"testing"
	"time"
)

func TestHeadless(t *testing.T) {
	profiling.StartProfiling(false, false)

	eng := engine.New("TestHeadless")

	ent := eng.NewEntity()
	tr := std.NewTransform()
	tr.Set(
		&std.Vector3{0, 0, 3}, // POSITION
		&std.Vector3{0, 0, 0}, // CENTER
		&std.Vector3{0, 1, 0}, // UP
	)
	ent.AddComponent(tr)

	go func(e *engine.Engine) {
		time.Sleep(5 * time.Second)
		e.Stop()
	}(eng)

	eng.Start()
	profiling.ReportUPS(eng)
	profiling.StopProfiling()
}
