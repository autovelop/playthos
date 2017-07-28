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

func (p *Physics) Destroy() {}

func (p *Physics) AddIntegrant(integrant engine.IntegrantRoutine) {}

func (p *Physics) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Velocity:
		p.velocities = append(p.velocities, component)
		break
	case *Acceleration:
		p.accelerations = append(p.accelerations, component)
		break
	}
}
func (p *Physics) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(p.velocities); i++ {
		velocity := p.velocities[i]
		if velocity.Entity().ID() == entity.ID() {
			copy(p.accelerations[i:], p.accelerations[i+1:])
			p.accelerations[len(p.accelerations)-1] = nil
			p.accelerations = p.accelerations[:len(p.accelerations)-1]

			copy(p.velocities[i:], p.velocities[i+1:])
			p.velocities[len(p.velocities)-1] = nil
			p.velocities = p.velocities[:len(p.velocities)-1]
		}
	}
}

func (p *Physics) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Acceleration{}, &Velocity{}}
}

func (p *Physics) Update() {
	if len(p.velocities) != len(p.accelerations) {
		log.Fatalf("Each acceleration component must be paired with a velocity component. And vice versa.")
	}
	for idx, velocity := range p.velocities {
		if velocity.Active() {
			entity := velocity.Entity()
			if entity != nil {
				acceleration := p.accelerations[idx]
				if acceleration.Active() {
					new_velocity := std.Vector3{acceleration.X + velocity.X, acceleration.Y + velocity.Y, acceleration.Z + velocity.Z}
					velocity.Set(new_velocity.X, new_velocity.Y, new_velocity.Z)
					transform := entity.Component(&std.Transform{}).(*std.Transform)
					position := transform.Position()
					transform.SetPosition(position.X+new_velocity.X, position.Y+new_velocity.Y, position.Z+new_velocity.Z)
				} else {
					// new_velocity := std.Vector3{acceleration.X + velocity.X, acceleration.Y + velocity.Y, acceleration.Z + velocity.Z}
					// velocity.Set(new_velocity.X, new_velocity.Y, new_velocity.Z)
					transform := entity.Component(&std.Transform{}).(*std.Transform)
					position := transform.Position()
					transform.SetPosition(position.X+velocity.X, position.Y+velocity.Y, position.Z+velocity.Z)
				}
			}
		}
	}
}
