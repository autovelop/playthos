package std

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func (c *Color) Copy() Animatable {
	n := new(Color)
	n.R += c.R
	n.G += c.G
	n.B += c.B
	n.A += c.A
	return n
}

func (c *Color) Add(a Animatable) {
	o := a.(*Color)
	c.R += o.R
	c.G += o.G
	c.B += o.B
	c.A += o.A
}

func (c *Color) Sub(s Animatable) {
	o := s.(*Color)
	c.R -= o.R
	c.G -= o.G
	c.B -= o.B
	c.A -= o.A
}

func (c *Color) Div(d float32) {
	c.R = c.R / d
	c.G = c.G / d
	c.B = c.B / d
	c.A = c.A / d
}

func (c *Color) Zero() {
	c.R = 0
	c.G = 0
	c.B = 0
	c.A = 0
}

func (c *Color) One() {
	c.R = 0
	c.G = 0
	c.B = 0
	c.A = 0
}

func (c *Color) Set(n float32) {
	c.R = n
	c.G = n
	c.B = n
	c.A = n
}
func (c *Color) Mul(d Animatable) {
	o := d.(*Color)
	c.R = c.R * o.R
	c.G = c.G * o.G
	c.B = c.B * o.B
	c.A = c.A * o.A
}
