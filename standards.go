package engine

// type Color [4]float32
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Rect struct {
	Vector2
	W float32
	H float32
}

type Circle struct {
	R float32
}
