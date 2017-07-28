// +build animation

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

func init() {
	engine.NewSystem(&Animation{})
}

type Animation struct {
	engine.System
	clips []*AnimationClip
}

func (a *Animation) InitSystem() {}

func (a *Animation) Destroy() {}

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
		if clip.Active() {
			clip.Update()
		}
	}
}

func NewClip(s float64, c float64, v std.Animatable) *AnimationClip {
	a := &AnimationClip{}
	// not active until it has a property
	// a.SetActive(true)
	a.Set(s, c, v)
	return a
}

func (a *Animation) AddComponent(component engine.ComponentRoutine) {
	switch clip := component.(type) {
	case *AnimationClip:
		a.clips = append(a.clips, clip)
		break
	}
}

func (a *Animation) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&AnimationClip{}}
}

func (a *Animation) AddIntegrant(integrant engine.IntegrantRoutine) {}
