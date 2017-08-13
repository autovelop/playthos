// +build autovelop_playthos_collision !play

package collision

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Collider struct {
	engine.Component
	transform *std.Transform
	relative  *std.Rect
	onHit     func(*engine.Entity)
}

func NewCollider() *Collider {
	return &Collider{}
}
func (c *Collider) Set(transform *std.Transform, relative *std.Rect) {
	c.transform = transform
	c.relative = relative
}

func (c *Collider) OnHit(onHit func(*engine.Entity)) {
	c.onHit = onHit
}

func (c *Collider) Hit(other *engine.Entity) {
	if c.onHit != nil {
		c.onHit(other)
	}
}

func (c *Collider) Get() (*std.Transform, *std.Rect) {
	return c.transform, c.relative
}
