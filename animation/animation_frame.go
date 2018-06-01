// +build autovelop_playthos_animation !play

package animation

import (
	"github.com/autovelop/playthos/std"
)

// Keyframe holds the duration that a animation frame as well as the targeted value it is animating towards
type KeyFrame struct {
	frameIndex float64

	// will hold the same frame for this duration
	duration float64

	target std.Animatable
	step   std.Animatable
}

func (a *KeyFrame) set(i float64, d float64, t std.Animatable) {
	a.frameIndex = i
	a.duration = d
	a.target = t
}

// SetStep set/overwrites the current stop of the keyframe
func (a *KeyFrame) SetStep(s std.Animatable) {
	a.step = s
}

// FrameIndex returns current frame index
func (a *KeyFrame) FrameIndex() float64 {
	return a.frameIndex
}

// FrameIndex returns frame duration
func (a *KeyFrame) Duration() float64 {
	return a.duration
}

// Step steps forward towards the target
func (a *KeyFrame) Step(value std.Animatable) {
	if a.step != nil {
		value.Add(a.step)
	}
}
