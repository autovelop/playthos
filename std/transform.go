package std

import (
	"github.com/autovelop/playthos"
)

type Transform struct {
	engine.Component
	position Vector3
	rotation Vector3
	scale    Vector3
}

func (t *Transform) Set(pos Vector3, rot Vector3, scl Vector3) {
	t.position = pos
	t.rotation = rot
	t.scale = scl
}

func (t *Transform) SetPosition(position Vector3) {
	t.position = position
}

func (t *Transform) GetPosition() Vector3 {
	return t.position
}

func (t *Transform) GetPosition2D() Vector2 {
	return Vector2{t.position.X, t.position.Y}
}

func (t *Transform) GetRotation() Vector3 {
	return t.rotation
}

func (t *Transform) GetScale() Vector3 {
	return t.scale
}
