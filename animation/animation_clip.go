// +build autovelop_playthos_animation !play

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	"log"
)

// Clip is the sequence of all the frames of a animation. It also controls the state, direction, repetition, and speed of the animation.
type Clip struct {
	engine.Component
	frames          []*KeyFrame
	currentKeyFrame *KeyFrame
	value           std.Animatable
	running         bool
	paused          bool
	autoplay        bool
	loop            bool
	speed           float64
	frameCount      float64
	tick            float64
}

// NewClip creates and sets a new orphan clip
func NewClip(s float64, f float64, v std.Animatable) *Clip {
	a := &Clip{}
	// not active until it has a property
	// a.SetActive(true)
	a.Set(s, f, v)
	return a
}

// Set used to define all the require properties of a Clip
func (a *Clip) Set(s float64, f float64, o std.Animatable) {
	a.speed = s
	a.frameCount = f
	a.value = o
	a.autoplay = true
	a.running = true
}

// Update is called on every tick of the game loop if running is set to true
func (a *Clip) Update() {
	if a.running {
		if a.tick >= a.frameCount {
			a.tick = 0

			// do a proper uncalculated reset without math
			diff := a.value.Copy()
			diff.Sub(a.frames[0].target)
			a.value.Sub(diff)
		}
		for _, f := range a.frames {
			if a.tick == f.FrameIndex() {
				// just change current key frame to this frame
				a.currentKeyFrame = f
				break
			}
		}
		if a.tick > (a.currentKeyFrame.FrameIndex() + a.currentKeyFrame.Duration() - 1) {
			a.currentKeyFrame.Step(a.value)
		}
		a.tick++
	}
}

// Stop and reset the clip
func (a *Clip) Stop() {
	a.running = false
	a.paused = false
	a.tick = 0
}

// Start or resume the clip
func (a *Clip) Start() {
	a.running = true
	a.paused = false
}

// Pause the clip
func (a *Clip) Pause() {
	a.paused = true
	a.running = true
}

// SetAutoplay set the autoplay property of the clip
func (a *Clip) SetAutoplay(ap bool) {
	a.autoplay = ap
	a.running = ap
	// a.running = true
}

// Loop set the clip to loop
func (a *Clip) Loop() {
	a.loop = true
}

// AddKeyFrame appends an animation frame to the clip.
func (a *Clip) AddKeyFrame(i float64, d float64, t std.Animatable) {
	// calculate and set keyframe step
	k := &KeyFrame{}
	k.set(i, d, t)
	a.frames = append(a.frames, k)
	if len(a.frames) > 1 {
		prevFrame := a.frames[len(a.frames)-2]

		// get difference between this frame and prev
		// also minus the duration of the prevFrame
		diffFrames := k.FrameIndex() - prevFrame.FrameIndex()
		if prevFrame.Duration() > diffFrames {
			log.Fatal("frame duration cannot go over to the next keyframe. Reduce duration or move next keyframe up")
		}

		diffValue := k.target.Copy()
		diffValue.Sub(prevFrame.target)
		diffValue.Div(float32(diffFrames - prevFrame.Duration()))

		prevFrame.SetStep(diffValue)
	}
}
