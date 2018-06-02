package std

// Color defines the RGBA values
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// Copy clones a color
func (c *Color) Copy() Animatable {
	n := new(Color)
	n.R += c.R
	n.G += c.G
	n.B += c.B
	n.A += c.A
	return n
}

// Add adds two color values
func (c *Color) Add(a Animatable) {
	o := a.(*Color)
	c.R += o.R
	c.G += o.G
	c.B += o.B
	c.A += o.A
}

// Sub subtracts two color values
func (c *Color) Sub(s Animatable) {
	o := s.(*Color)
	c.R -= o.R
	c.G -= o.G
	c.B -= o.B
	c.A -= o.A
}

// Div divides a color's values
func (c *Color) Div(d float32) {
	c.R = c.R / d
	c.G = c.G / d
	c.B = c.B / d
	c.A = c.A / d
}

// Zero sets all color values to zero
func (c *Color) Zero() {
	c.R = 0
	c.G = 0
	c.B = 0
	c.A = 0
}

// One sets all color values to zero
func (c *Color) One() {
	c.R = 0
	c.G = 0
	c.B = 0
	c.A = 0
}

// Set used to define all the require properties of a Clip
func (c *Color) Set(n float32) {
	c.R = n
	c.G = n
	c.B = n
	c.A = n
}

// Div multiplies two color values
func (c *Color) Mul(d Animatable) {
	o := d.(*Color)
	c.R = c.R * o.R
	c.G = c.G * o.G
	c.B = c.B * o.B
	c.A = c.A * o.A
}
