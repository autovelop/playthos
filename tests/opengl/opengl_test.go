package playthos_test

import (
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/opengl"
	"github.com/autovelop/playthos/profiling"
	"testing"
	"time"
)

func TestOpenGL(t *testing.T) {
	profiling.StartProfiling(false, false)

	eng := engine.New("TestOpenGL", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	go func(e *engine.Engine) {
		time.Sleep(5 * time.Second)
		e.Stop()
	}(eng)

	eng.Start()
	profiling.ReportUPS(eng)
	profiling.StopProfiling()
}
