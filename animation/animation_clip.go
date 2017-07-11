// +build animation

package animation

import (
	"github.com/autovelop/playthos"
)

type AnimationClip struct {
	engine.Component
	speed  uint
	frames []engine.ComponentRoutine
}

func (a *AnimationClip) SetSpeed(speed uint) {
	a.speed = speed
}

func (a *AnimationClip) CreateFrames(size uint) {
	a.frames = make([]engine.ComponentRoutine, size)
}

func (a *AnimationClip) SetFrame(index uint, component engine.ComponentRoutine) {
	a.frames[index] = component
}
