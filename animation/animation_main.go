// +build animation

package animation

import (
	"github.com/autovelop/playthos"
	// "github.com/autovelop/playthos/std"
	"log"
	// "math"
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
		// if clip.autoplay && !clip.playing {
		// 	clip.Play()
		// }
		if clip.playing {

			clip.ticks += (1 / clip.duration)
			clip.percCompleted = (clip.ticks / clip.duration)

			if clip.percCompleted >= 1 {
				clip.ticks = 0
				clip.percCompleted = 0
				clip.Reset()
			}

			clip.Update()
			// log.Println(clip.percCompleted)
			log.Println()
			// log.Println(clip.value)
			// clip.Value.Sub(clip.Step)
			// log.Println(int(clip.progress * (clip.Frames)))
			// if clip.progress < float64(clip.speed) {
			// activeFrameIndex := int(math.Floor(clip.progress * float64(len(clip.frames))))
			// activeFrame := clip.frames[activeFrameIndex]
			// log.Println(activeFrame)
			// if activeFrame != nil {
			// 	if activeFrame.step != nil {
			// clip.currentStep = activeFrame.step
			// clip.value.Add(activeFrame.step)
			// }
			// clip.value = activeFrame.value
			// }
			// clip.progress += 0.01

			// get frame
			// if clip.progress >= 1 {
			// 	clip.progress = 0
			// }

			// 	if clip.reversing {
			// 		clip.Value.Sub(clip.Step)
			// 		// 	clip.From.X -= clip.step.X
			// 		// 	clip.From.Y -= clip.step.Y
			// 		// 	clip.From.Z -= clip.step.Z
			// 	} else {
			// 		// log.Println(clip.From)
			// 		clip.Value.Add(clip.Step)
			// 		// 	clip.From.X += clip.step.X
			// 		// 	clip.From.Y += clip.step.Y
			// 		// 	clip.From.Z += clip.step.Z
			// 	}
			// } else {
			// 	if clip.loop {
			// 		if clip.reverse {
			// 			clip.reversing = !clip.reversing
			// 		} else {
			// 			// n := new(std.Animatable)
			// 			// var b *std.Animatable
			// 			// b = &clip.Start
			// 			// log.Fatalf("%v", clip.Diff)
			// 			clip.Value.Sub(clip.Diff)
			// 			// log.Fatalf("%v - %v", clip.From, clip.To)
			// 			// clip.From = clip.Start
			// 		}
			// 		clip.progress = 0.0
			// 	} else {
			// 		clip.Stop()
			// 	}
			// }
		}
	}
}

func NewClip(s float64, au bool) *AnimationClip {
	a := &AnimationClip{}
	// not active until it has a property
	// a.SetActive(true)
	a.Set(s, au)
	return a
}

func (a *Animation) AddComponent(component engine.ComponentRoutine) {
	switch clip := component.(type) {
	case *AnimationClip:
		// clip.Value = &component.From
		// switch v := clip.From.(type) {
		// case *std.Vector3:
		// 	var diff std.Vector3 = v.Diff(clip.To.(*std.Vector3))
		// 	clip.step = diff.Div(float32(clip.speed))
		// }
		// diff := clip.From.Diff(clip.To)
		// clip.step = diff.Div(float32(clip.speed))
		// log.Println(clip.step)
		// log.Print(&clip.To)
		// log.Print(&clip.step)
		// log.Fatal(clip.From)

		a.clips = append(a.clips, clip)
		break
	}
}

func (a *Animation) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&AnimationClip{}}
}

func (a *Animation) AddIntegrant(integrant engine.IntegrantRoutine) {}
