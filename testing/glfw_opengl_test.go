package playthos_test

import (
	// "bytes"
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/opengl"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	// "log"
	"os"
	// "runtime"
	"runtime/pprof"
	"testing"
	"time"
)

// All OpenGL tests are conducted in windowed mode

func TestGLFWOpenGL(t *testing.T) {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	eng := engine.New("TestGLFWOpenGL", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	ent := eng.NewEntity()
	tr := std.NewTransform()
	tr.Set(
		&std.Vector3{0, 0, 3}, // POSITION
		&std.Vector3{0, 0, 0}, // CENTER
		&std.Vector3{0, 1, 0}, // UP
	)
	ent.AddComponent(tr)

	camera := render.NewCamera()
	cameraSize := float32(4)
	camera.Set(&cameraSize, &std.Color{0.27, 0.20, 0.54, 0})
	camera.SetTransform(tr)

	ent.AddComponent(camera)

	go func(e *engine.Engine) {
		time.Sleep(6 * time.Second)
		e.Stop()
	}(eng)

	eng.Once()
	pprof.WriteHeapProfile(os.Stdout)
}
