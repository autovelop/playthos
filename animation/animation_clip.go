// +build animation

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type AnimationClip struct {
	engine.Component
	speed float64
	// Value *float64
	// From *struct{ float64 }
	From     *std.Vector3
	Start    std.Vector3
	To       std.Vector3
	step     std.Vector3
	progress float64
	// From *float64
	// To    float32
	loop      bool
	reverse   bool
	reversing bool
	// frames []engine.ComponentRoutine
}

func (a *AnimationClip) SetSpeed(speed float64) {
	a.speed = speed
}

func (a *AnimationClip) Loop() {
	a.loop = true
	a.Start = *a.From
}

func (a *AnimationClip) Reverse() {
	a.reverse = true
}

// func (a *AnimationClip) CreateFrames(size uint) {
// 	a.frames = make([]engine.ComponentRoutine, size)
// }

func (a *AnimationClip) SetFrame(index uint, component engine.ComponentRoutine) {
	// a.frames[index] = component
}
