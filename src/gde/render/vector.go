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

type Vector2 struct {
	X float32
	Y float32
}
