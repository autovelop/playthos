// +build animation

package animation

import (
	"github.com/autovelop/playthos"
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
}

func NewClip(frames uint, speed uint) *AnimationClip {
	a := &AnimationClip{}
	a.SetActive(true)
	a.SetSpeed(speed)
	a.CreateFrames(frames)
	return a
}

func (a *Animation) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *AnimationClip:
		a.clips = append(a.clips, component)
		break
	}
}

func (a *Animation) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&AnimationClip{}}
}

func (a *Animation) NewIntegrant(integrant engine.IntegrantRoutine) {}
