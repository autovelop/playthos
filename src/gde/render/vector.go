package render

type Vector4 [4]float32

func (v Vector4) ToUniformFloat() []float32 {
	return []float32{0, 0, 0, 0}
}

// func (v *Vector4) Add(x float32)

// type Vector4 struct {
// 	X float32
// 	Y float32
// 	Z float32
// 	W float32
// }

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (v1 *Vector3) Add(v2 *Vector3) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

type Vector2 struct {
	X float32
	Y float32
}
