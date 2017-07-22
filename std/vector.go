package std

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

type Integer struct {
	V int
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

func (v *Vector3) Copy() Animatable {
	n := new(Vector3)
	n.X += v.X
	n.Y += v.Y
	n.Z += v.Z
	return n
}

func (v *Vector3) Add(a Animatable) {
	o := a.(*Vector3)
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
}

func (v *Vector3) Sub(s Animatable) {
	o := s.(*Vector3)
	v.X -= o.X
	v.Y -= o.Y
	v.Z -= o.Z
}

func (v *Vector3) Div(d float32) {
	v.X = v.X / d
	v.Y = v.Y / d
	v.Z = v.Z / d
}

func (v *Vector3) Zero() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

func (v *Vector3) One() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

func (v *Vector3) Set(n float32) {
	v.X = n
	v.Y = n
	v.Z = n
}
func (v *Vector3) Mul(d Animatable) {
	o := d.(*Vector3)
	v.X = v.X * o.X
	v.Y = v.Y * o.Y
	v.Z = v.Z * o.Z
}

func (v *Vector2) Zero() {
	v.X = 0
	v.Y = 0
}
func (v *Vector2) One() {
	v.X = 0
	v.Y = 0
}
func (v *Vector2) Set(n float32) {
	v.X = n
	v.Y = n
}

func (v *Vector2) Copy() Animatable {
	n := new(Vector2)
	n.X += v.X
	n.Y += v.Y
	return n
}

func (v *Vector2) Mul(d Animatable) {
	o := d.(*Vector2)
	v.X = v.X * o.X
	v.Y = v.Y * o.Y
}

func (v *Vector2) Add(a Animatable) {
	o := a.(*Vector2)
	v.X += o.X
	v.Y += o.Y
}

func (v *Vector2) Sub(s Animatable) {
	o := s.(*Vector2)
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vector2) Div(d float32) {
	v.X = v.X / d
	v.Y = v.Y / d
}

func (v *Integer) Zero() {
	v.V = 0
}
func (v *Integer) One() {
	v.V = 0
}
func (v *Integer) Set(n float32) {
	v.V = int(n)
}

func (v *Integer) Copy() Animatable {
	n := new(Integer)
	n.V += v.V
	return n
}

func (v *Integer) Mul(d Animatable) {
	o := d.(*Integer)
	v.V = v.V * o.V
}

func (v *Integer) Add(a Animatable) {
	o := a.(*Integer)
	v.V += o.V
}

func (v *Integer) Sub(s Animatable) {
	o := s.(*Integer)
	v.V -= o.V
}

func (v *Integer) Div(d float32) {
	v.V = int(float32(v.V) / d)
}
