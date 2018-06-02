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

// Physics system used to simulate rigidbodies physics (force, gravity, friction, etc.)
type Physics struct {
	// Need to be able to uncomment the below at some point
	// engine.Updater
	engine.System
	rigidbodies []*RigidBody
	gravity     *std.Vector3
}

// InitSystem called when the system plugs into the engine
func (p *Physics) InitSystem() {
	p.gravity = &std.Vector3{0, -9.8, 0}
}

// Destroy called when engine is gracefully shutting down
func (p *Physics) Destroy() {}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (p *Physics) AddIntegrant(integrant engine.IntegrantRoutine) {}

// AddComponent unorphans a component by adding it to this system
func (p *Physics) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *RigidBody:
		p.rigidbodies = append(p.rigidbodies, component)
		component.SetForce(p.gravity.X, p.gravity.Y, p.gravity.Z)
		break
	}
}

// DeleteEntity removes all entity's compoents from this system
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

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (p *Physics) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&RigidBody{}}
}

// Update called by engine to progress this system to the next engine loop
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
