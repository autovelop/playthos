// +build animation

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	// "log"
)

func init() {
	engine.NewSystem(&Animation{})
}

type Animation struct {
	engine.System
	clips []*AnimationClip
}

func (a *Animation) InitSystem() {}

func (a *Animation) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(a.clips); i++ {
		clip := a.clips[i]
		if clip.Entity().ID() == entity.ID() {
			copy(a.clips[i:], a.clips[i+1:])
			a.clips[len(a.clips)-1] = nil
			a.clips = a.clips[:len(a.clips)-1]
		}
	}
}

func (a *Animation) Update() {
	for _, clip := range a.clips {
		if clip.progress < float64(clip.speed) {
			clip.progress += 0.01
			if clip.reversing {
				clip.From.X -= clip.step.X
				clip.From.Y -= clip.step.Y
				clip.From.Z -= clip.step.Z
			} else {
				clip.From.X += clip.step.X
				clip.From.Y += clip.step.Y
				clip.From.Z += clip.step.Z
			}
		} else if clip.loop {
			if clip.reverse {
				clip.reversing = !clip.reversing
			} else {
				*clip.From = clip.Start
			}
			clip.progress = 0.0
		}
	}
}

func NewClip(frames uint, speed float64) *AnimationClip {
	a := &AnimationClip{}
	a.SetActive(true)
	a.SetSpeed(speed)
	// a.CreateFrames(frames)
	return a
}

func (a *Animation) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *AnimationClip:
		// component.Value = &component.From

		var diff std.Vector3 = component.From.Diff(component.To)
		component.step = diff.Div(float32(component.speed))
		a.clips = append(a.clips, component)
		break
	}
}

func (a *Animation) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&AnimationClip{}}
}

func (a *Animation) NewIntegrant(integrant engine.IntegrantRoutine) {}
