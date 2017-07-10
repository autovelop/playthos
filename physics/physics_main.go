// +build physics

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	"log"
)

func init() {
	engine.NewSystem(&Physics{})
}

type Physics struct {
	// Need to be able to uncomment the below at some point
	// engine.Updater
	engine.System
	accelerations []*Acceleration
	velocities    []*Velocity
}

func (p *Physics) InitSystem() {}

func (p *Physics) NewIntegrant(integrant engine.IntegrantRoutine) {}

func (p *Physics) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Velocity:
		p.velocities = append(p.velocities, component)
		break
	case *Acceleration:
		p.accelerations = append(p.accelerations, component)
		break
	}
}

func (p *Physics) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Acceleration{}, &Velocity{}}
}

func (p *Physics) Update() {
	if len(p.velocities) != len(p.accelerations) {
		log.Println("Skew components")
		log.Fatalf("velocities: %v | accelerations: %v", len(p.velocities), len(p.accelerations))
	}
	for idx, velocity := range p.velocities {
		entity := velocity.Entity()
		if entity != nil {
			acceleration := p.accelerations[idx]
			new_velocity := std.Vector3{acceleration.X + velocity.X, acceleration.Y + velocity.Y, acceleration.Z + velocity.Z}
			velocity.Set(new_velocity.X, new_velocity.Y, new_velocity.Z)
			transform := entity.Component(&std.Transform{}).(*std.Transform)
			position := transform.Position()
			transform.SetPosition(&std.Vector3{position.X + new_velocity.X, position.Y + new_velocity.Y, position.Z + new_velocity.Z})
		}
	}
}
