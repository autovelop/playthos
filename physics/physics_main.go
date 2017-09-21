// +build deploy physics

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	// "log"
)

func init() {
	engine.NewSystem(&Physics{})
}

type Physics struct {
	// Need to be able to uncomment the below at some point
	// engine.Updater
	engine.System
	rigidbodies []*RigidBody
	gravity     *std.Vector3
}

func (p *Physics) InitSystem() {
	p.gravity = &std.Vector3{0, -9.8, 0}
}

func (p *Physics) Destroy() {}

func (p *Physics) AddIntegrant(integrant engine.IntegrantRoutine) {}

func (p *Physics) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *RigidBody:
		p.rigidbodies = append(p.rigidbodies, component)
		component.SetForce(p.gravity.X, p.gravity.Y, p.gravity.Z)
		break
	}
}
func (p *Physics) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(p.rigidbodies); i++ {
		rigidbody := p.rigidbodies[i]
		if rigidbody.Entity().ID() == entity.ID() {
			copy(p.rigidbodies[i:], p.rigidbodies[i+1:])
			p.rigidbodies[len(p.rigidbodies)-1] = nil
			p.rigidbodies = p.rigidbodies[:len(p.rigidbodies)-1]
		}
	}
}

func (p *Physics) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&RigidBody{}}
}

func (p *Physics) Update() {
	for _, rigidbody := range p.rigidbodies {
		if rigidbody.Active() {
			entity := rigidbody.Entity()
			if entity != nil {
				velocity := rigidbody.Velocity()
				force := rigidbody.Force()

				// why do this additional multiplication to get a realistic result?
				mass := rigidbody.Mass() / (20 * 100)

				new_velocity := std.Vector3{velocity.X + force.X, velocity.Y + (force.Y * mass), velocity.Z + force.Z}
				rigidbody.SetVelocity(new_velocity.X, new_velocity.Y, new_velocity.Z)
				transform := entity.Component(&std.Transform{}).(*std.Transform)
				position := transform.Position()
				transform.SetPosition(position.X+new_velocity.X, position.Y+new_velocity.Y, position.Z+new_velocity.Z)
			}
		}
	}
}
