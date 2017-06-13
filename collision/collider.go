// +build collision

package collision

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Collider struct {
	engine.Component
	transform *std.Transform
	relative  *std.Rect
	onHit     func()
}

func (c *Collider) Set(transform *std.Transform, relative *std.Rect) {
	c.transform = transform
	c.relative = relative
}

func (c *Collider) OnHit(onHit func()) {
	c.onHit = onHit
}

func (c *Collider) Hit() {
	if c.onHit != nil {
		c.onHit()
	}
}

func (c *Collider) Get() (*std.Transform, *std.Rect) {
	return c.transform, c.relative
}
