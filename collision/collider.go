// +build collision !play

package collision

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Collider struct {
	engine.Component
	position *std.Vector3
	offset   *std.Vector2
	size     *std.Vector2
	onHit    func(*Collider)
	reverse  bool
}

func NewCollider() *Collider {
	return &Collider{}
}
func (c *Collider) Set(position *std.Vector3, offset *std.Vector2, size *std.Vector2) {
	c.position = position
	// calculate bounds based on shape & orientation
	c.offset = offset
	c.size = size
}

func (c *Collider) Reverse() bool {
	return c.reverse
}

func (c *Collider) SetReverse(r bool) {
	c.reverse = r
}

func (c *Collider) OnHit(onHit func(*Collider)) {
	c.onHit = onHit
}

func (c *Collider) Hit(other *Collider) {
	if c.onHit != nil {
		c.onHit(other)
	}
}

func (c *Collider) Get() (*std.Vector3, *std.Vector2, *std.Vector2) {
	return c.position, c.offset, c.size
}
