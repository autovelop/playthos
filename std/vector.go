package std

// Animatable interface allows math to be applied to variables for animation
//
// TODO(F): Rename this to something more generic because it could be applied to something other than animation
// TODO(F): Split other structs inside vector.go file to their respective files
type Animatable interface {
	Add(Animatable)
	Sub(Animatable)
	Div(float32)
	Set(float32)
	Mul(Animatable)
	Copy() Animatable
	Zero()
	One()
}

// Integer defines an integer
type Integer struct {
	V int
}

// Vector2 defines two vector values (x, y)
type Vector2 struct {
	X float32
	Y float32
}

// Vector3 defines three vector values (x, y, z)
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// Copy clones a vector3
func (v *Vector3) Copy() Animatable {
	n := new(Vector3)
	n.X += v.X
	n.Y += v.Y
	n.Z += v.Z
	return n
}

// Add adds two vector3 values
func (v *Vector3) Add(a Animatable) {
	o := a.(*Vector3)
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
}

// Sub subtracts two vector3 values
func (v *Vector3) Sub(s Animatable) {
	o := s.(*Vector3)
	v.X -= o.X
	v.Y -= o.Y
	v.Z -= o.Z
}

// Div divides a vector3's values
func (v *Vector3) Div(d float32) {
	v.X = v.X / d
	v.Y = v.Y / d
	v.Z = v.Z / d
}

// Zero sets vector3 values to zero
func (v *Vector3) Zero() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

// One sets vector3 values to zero
func (v *Vector3) One() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

// Set used to define all the require properties of a Vector3
func (v *Vector3) Set(n float32) {
	v.X = n
	v.Y = n
	v.Z = n
}

// Div multiplies two vector3 values
func (v *Vector3) Mul(d Animatable) {
	o := d.(*Vector3)
	v.X = v.X * o.X
	v.Y = v.Y * o.Y
	v.Z = v.Z * o.Z
}

// Zero sets vector2 values to zero
func (v *Vector2) Zero() {
	v.X = 0
	v.Y = 0
}

// One sets vector2 values to zero
func (v *Vector2) One() {
	v.X = 0
	v.Y = 0
}

// Set used to define all the require properties of a Vector2
func (v *Vector2) Set(n float32) {
	v.X = n
	v.Y = n
}

// Copy clones a vector2
func (v *Vector2) Copy() Animatable {
	n := new(Vector2)
	n.X += v.X
	n.Y += v.Y
	return n
}

// Div multiplies two vector2 values
func (v *Vector2) Mul(d Animatable) {
	o := d.(*Vector2)
	v.X = v.X * o.X
	v.Y = v.Y * o.Y
}

// Add adds two vector2 values
func (v *Vector2) Add(a Animatable) {
	o := a.(*Vector2)
	v.X += o.X
	v.Y += o.Y
}

// Sub subtracts two vector2 values
func (v *Vector2) Sub(s Animatable) {
	o := s.(*Vector2)
	v.X -= o.X
	v.Y -= o.Y
}

// Div divides a vector2's values
func (v *Vector2) Div(d float32) {
	v.X = v.X / d
	v.Y = v.Y / d
}

// Zero sets integer value to zero
func (v *Integer) Zero() {
	v.V = 0
}

// One sets integer value to zero
func (v *Integer) One() {
	v.V = 0
}

// Set used to define all the require properties of a Integer
func (v *Integer) Set(n float32) {
	v.V = int(n)
}

// Copy clones an integer
func (v *Integer) Copy() Animatable {
	n := new(Integer)
	n.V += v.V
	return n
}

// Div multiplies two integer values
func (v *Integer) Mul(d Animatable) {
	o := d.(*Integer)
	v.V = v.V * o.V
}

// Add adds two integer values
func (v *Integer) Add(a Animatable) {
	o := a.(*Integer)
	v.V += o.V
}

// Sub subtracts two integer values
func (v *Integer) Sub(s Animatable) {
	o := s.(*Integer)
	v.V -= o.V
}

// Div divides a integer's value
func (v *Integer) Div(d float32) {
	v.V = int(float32(v.V) / d)
}
