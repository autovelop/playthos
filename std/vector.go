package std

type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (v *Vector3) Add(add Vector3) {
	v.X += add.X
	v.Y += add.Y
	v.Z += add.Z
}

func (v *Vector3) Diff(diff Vector3) Vector3 {
	return Vector3{diff.X - v.X, diff.Y - v.Y, diff.Z - v.Z}
}

func (v *Vector3) Div(d float32) Vector3 {
	return Vector3{v.X / d, v.Y / d, v.Z / d}
}
