package physics

import (
	"gde/engine"
)

type Physics struct {
	engine.System
}

func (p *Physics) Init() {}

func (p *Physics) Update(entities *map[string]*engine.Entity) {
	for _, v := range *entities {
		dynamic := v.GetComponent(&Dynamic{})
		switch dynamic := dynamic.(type) {
		case PhysicsRoutine:
			dynamic.UpdateGravity()
			dynamic.UpdateVelocity()
			// dynamic.UpdateForce()
			dynamic.UpdateCollision()
			dynamic.UpdateTrigger()
		}
		static := v.GetComponent(&Static{})
		switch static := static.(type) {
		case PhysicsRoutine:
			static.UpdateTrigger()
		}
	}
}

func (p *Physics) Stop() {}
