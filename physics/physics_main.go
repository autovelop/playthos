// +build physics

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	"log"
)

var componentTypes []engine.ComponentRoutine = []engine.ComponentRoutine{&Acceleration{}, &Velocity{}}

func init() {
	phy := &Physics{}
	engine.NewSystem(phy)
}

type Physics struct {
	engine.System
	accelerations []*Acceleration
	velocities    []*Velocity
}

func (p *Physics) Prepare() {
}

func (p *Physics) LoadComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Velocity:
		p.RegisterVelocity(component)
		log.Println("LoadComponent(*Velocity)")
		break
	case *Acceleration:
		p.RegisterAcceleration(component)
		log.Println("LoadComponent(*Acceleration)")
		break
	}
}

func (p *Physics) ComponentTypes() []engine.ComponentRoutine {
	return componentTypes
}

func (p *Physics) Update() {
	for idx, velocity := range p.velocities {
		entity := velocity.GetEntity()
		if entity != nil {
			acceleration := p.accelerations[idx]
			new_velocity := std.Vector3{acceleration.X + velocity.X, acceleration.Y + velocity.Y, acceleration.Z + velocity.Z}
			velocity.Set(new_velocity.X, new_velocity.Y, new_velocity.Z)
			transform := entity.GetComponent(&std.Transform{}).(*std.Transform)
			position := transform.GetPosition()
			transform.SetPosition(std.Vector3{position.X + new_velocity.X, position.Y + new_velocity.Y, position.Z + new_velocity.Z})
			// log.Println(transform)
		}
	}
}

func (p *Physics) RegisterAcceleration(acceleration *Acceleration) {
	p.accelerations = append(p.accelerations, acceleration)
}

func (p *Physics) RegisterVelocity(velocity *Velocity) {
	p.velocities = append(p.velocities, velocity)
}

func (p *Physics) UnRegisterEntity(entity *engine.Entity) {
	for i := 0; i < len(p.accelerations); i++ {
		acceleration := p.accelerations[i]
		if acceleration.GetEntity().ID == entity.ID {
			copy(p.accelerations[i:], p.accelerations[i+1:])
			p.accelerations[len(p.accelerations)-1] = nil
			p.accelerations = p.accelerations[:len(p.accelerations)-1]

			copy(p.velocities[i:], p.velocities[i+1:])
			p.velocities[len(p.velocities)-1] = nil
			p.velocities = p.velocities[:len(p.velocities)-1]
		}
	}
}
