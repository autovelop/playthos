// +build deploy animation

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

func init() {
	engine.NewSystem(&Animation{})
}

// Animation system used to update all clips to animation to their targets
type Animation struct {
	engine.System
	clips []*AnimationClip
}

// InitSystem called when the system plugs into the engine
func (a *Animation) InitSystem() {}

// Destroy called when engine is gracefully shutting down
func (a *Animation) Destroy() {}

// DeleteEntity removes all entity's compoents from this system
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

// Update called by engine to progress this system to the next engine loop
func (a *Animation) Update() {
	for _, clip := range a.clips {
		if clip.Active() {
			clip.Update()
		}
	}
}

// NewClip creates and sets a new orphan clip
func NewClip(s float64, c float64, v std.Animatable) *AnimationClip {
	a := &AnimationClip{}
	// not active until it has a property
	// a.SetActive(true)
	a.Set(s, c, v)
	return a
}

// AddComponent unorphans a component by adding it to this system
func (a *Animation) AddComponent(component engine.ComponentRoutine) {
	switch clip := component.(type) {
	case *AnimationClip:
		a.clips = append(a.clips, clip)
		break
	}
}

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (a *Animation) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&AnimationClip{}}
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (a *Animation) AddIntegrant(integrant engine.IntegrantRoutine) {}
