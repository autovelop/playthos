package std

import (
	"github.com/autovelop/playthos"
)

type Transform struct {
	engine.Component
	position *Vector3
	rotation *Vector3
	scale    *Vector3
}

func NewTransform() *Transform {
	return &Transform{}
}

func (t *Transform) Set(pos *Vector3, rot *Vector3, scl *Vector3) {
	t.position = pos
	t.rotation = rot
	t.scale = scl
}

func (t *Transform) SetPosition(x float32, y float32, z float32) {
	t.position.X = x
	t.position.Y = y
	t.position.Z = z
}

func (t *Transform) SetRotation(x float32, y float32, z float32) {
	t.rotation.X = x
	t.rotation.Y = y
	t.rotation.Z = z
}

func (t *Transform) SetScale(scale *Vector3) {
	t.scale = scale
}

func (t *Transform) Position() *Vector3 {
	return t.position
}

func (t *Transform) Position2D() *Vector2 {
	return &Vector2{t.position.X, t.position.Y}
}

func (t *Transform) Rotation() *Vector3 {
	return t.rotation
}

func (t *Transform) Scale() *Vector3 {
	return t.scale
}
