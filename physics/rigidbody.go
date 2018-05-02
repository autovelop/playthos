// +build physics !play

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// RigidBody to apply force, velocity, mass, and friction to entity
type RigidBody struct {
	engine.Component
	velocity *std.Vector3
	force    *std.Vector3
	mass     float32
	friction float32
}

// NewRigidBody creates empty RigidBody with 1 mass and zero friction, velocity, and force
func NewRigidBody() *RigidBody {
	return &RigidBody{
		velocity: &std.Vector3{0, 0, 0},
		force:    &std.Vector3{0, 0, 0},
		mass:     1,
		friction: 0,
	}
}

// SetVelocity sets velocity of rigidbody
func (r *RigidBody) SetVelocity(x float32, y float32, z float32) {
	r.velocity.X = x
	r.velocity.Y = y
	r.velocity.Z = z
}

// Velocity returns current RigidBody velocity (Vector3)
func (r *RigidBody) Velocity() *std.Vector3 {
	return r.velocity
}

// AddForce adds force to current RigidBody force (3 x float32)
func (r *RigidBody) AddForce(x float32, y float32, z float32) {
	r.force.X += x
	r.force.Y += y
	r.force.Z += z
}

// SetForce sets new force values for RigidBody
func (r *RigidBody) SetForce(x float32, y float32, z float32) {
	r.force.X = x
	r.force.Y = y
	r.force.Z = z
}

// Force returns current RigidBody force (Vector3)
func (r *RigidBody) Force() *std.Vector3 {
	return r.force
}

// Friction returns current RigidBody friction (float32)
func (r *RigidBody) Friction() float32 {
	return r.friction
}

// SetMass sets new mass value for RigidBody
func (r *RigidBody) SetMass(m float32) {
	r.mass = m
}

// Mass returns current RigidBody mass (float32)
func (r *RigidBody) Mass() float32 {
	return r.mass
}
