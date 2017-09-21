// +build physics !play

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type RigidBody struct {
	engine.Component
	velocity *std.Vector3
	force    *std.Vector3
	mass     float32
	friction float32
}

func NewRigidBody() *RigidBody {
	return &RigidBody{
		velocity: &std.Vector3{0, 0, 0},
		force:    &std.Vector3{0, 0, 0},
		mass:     1,
		friction: 0,
	}
}

func (r *RigidBody) SetVelocity(x float32, y float32, z float32) {
	r.velocity.X = x
	r.velocity.Y = y
	r.velocity.Z = z
}

func (r *RigidBody) Velocity() *std.Vector3 {
	return r.velocity
}

func (r *RigidBody) AddForce(x float32, y float32, z float32) {
	r.force.X += x
	r.force.Y += y
	r.force.Z += z
}

func (r *RigidBody) SetForce(x float32, y float32, z float32) {
	r.force.X = x
	r.force.Y = y
	r.force.Z = z
}

func (r *RigidBody) Force() *std.Vector3 {
	return r.force
}

func (r *RigidBody) Friction() float32 {
	return r.friction
}

func (r *RigidBody) Mass() float32 {
	return r.mass
}
