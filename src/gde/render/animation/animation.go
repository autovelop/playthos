package animation

import (
	"gde/engine"
)

type Animation struct {
	engine.System
}

func (a *Animation) Init() {}

func (a *Animation) Update(entities *map[string]*engine.Entity) {
	for _, v := range *entities {
		animator := v.GetComponent(&Animator{})
		switch animator := animator.(type) {
		case AnimatorRoutine:
			animator.StepFrame()
		}
	}
}

func (a *Animation) Stop() {}
