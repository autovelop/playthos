package std

import (
	"github.com/autovelop/playthos"
)

// Transform defines the position, orientiation, and scale
type Transform struct {
	engine.Component
	position *Vector3
	rotation *Vector3
	scale    *Vector3
}

// NewTransform creates and sets a new orphan transform
func NewTransform() *Transform {
	return &Transform{}
}

// Set used to define all the require properties of a Clip
func (t *Transform) Set(pos *Vector3, rot *Vector3, scl *Vector3) {
	t.position = pos
	t.rotation = rot
	t.scale = scl
}

// SetPosition sets/changes position
func (t *Transform) SetPosition(x float32, y float32, z float32) {
	t.position.X = x
	t.position.Y = y
	t.position.Z = z
}

// SetRotation sets/changes rotation/orientation
func (t *Transform) SetRotation(x float32, y float32, z float32) {
	t.rotation.X = x
	t.rotation.Y = y
	t.rotation.Z = z
}

// SetScale sets/changes scale
func (t *Transform) SetScale(scale *Vector3) {
	t.scale = scale
}

// Position returns position vector
func (t *Transform) Position() *Vector3 {
	return t.position
}

// Position2D returns two-dimensional position vector
func (t *Transform) Position2D() *Vector2 {
	return &Vector2{t.position.X, t.position.Y}
}

// Rotation return rotation/orientation vector
func (t *Transform) Rotation() *Vector3 {
	return t.rotation
}

// Scale returns scale vector
func (t *Transform) Scale() *Vector3 {
	return t.scale
}
