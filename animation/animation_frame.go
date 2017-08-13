// +build autovelop_playthos_animation !play

package animation

import (
	"github.com/autovelop/playthos/std"
)

type AnimationKeyFrame struct {
	frameIndex float64

	// will hold the same frame for this duration
	duration float64

	target std.Animatable
	step   std.Animatable
}

func (a *AnimationKeyFrame) set(i float64, d float64, t std.Animatable) {
	a.frameIndex = i
	a.duration = d
	a.target = t
}

func (a *AnimationKeyFrame) SetStep(s std.Animatable) {
	a.step = s
}

func (a *AnimationKeyFrame) FrameIndex() float64 {
	return a.frameIndex
}

func (a *AnimationKeyFrame) Duration() float64 {
	return a.duration
}

func (a *AnimationKeyFrame) Step(value std.Animatable) {
	if a.step != nil {
		value.Add(a.step)
	}
}
