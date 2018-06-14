package playthos_test

import (
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/glfw"
	_ "github.com/autovelop/playthos/render"
	"testing"
)

func TestGLFWWindowed(t *testing.T) {
	eng := engine.New("TestGLFWWindowed", &engine.Settings{
		false,
		1024,
		768,
		false,
	})
	eng.Once()
}

// TODO(F): When a fullscreen test runs, then all subsequent GLFW instances are fullscreen
func TestGLFWFullscreen(t *testing.T) {
	eng := engine.New("TestGLFWFullscreen", &engine.Settings{
		true,
		1024,
		768,
		false,
	})
	eng.Once()
}
