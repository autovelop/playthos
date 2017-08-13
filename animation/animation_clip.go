// +build autovelop_playthos_animation !play

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	"log"
)

type AnimationClip struct {
	engine.Component
	frames          []*AnimationKeyFrame
	currentKeyFrame *AnimationKeyFrame
	value           std.Animatable
	running         bool
	paused          bool
	autoplay        bool
	loop            bool
	speed           float64
	frameCount      float64
	tick            float64
}

func (a *AnimationClip) Set(s float64, f float64, o std.Animatable) {
	a.speed = s
	a.frameCount = f
	a.value = o
	a.autoplay = true
	a.running = true
}

func (a *AnimationClip) Update() {
	if a.running {
		if a.tick > a.frameCount {
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

func (a *AnimationClip) Stop() {
	a.running = false
	a.paused = false
	a.tick = 0
}

func (a *AnimationClip) Start() {
	a.running = true
	a.paused = false
}

func (a *AnimationClip) Pause() {
	a.paused = true
	a.running = true
}

func (a *AnimationClip) SetAutoplay(ap bool) {
	a.autoplay = ap
	a.running = ap
	// a.running = true
}

func (a *AnimationClip) Loop() {
	a.loop = true
}

func (a *AnimationClip) AddKeyFrame(i float64, d float64, t std.Animatable) {
	// calculate and set keyframe step
	k := &AnimationKeyFrame{}
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
