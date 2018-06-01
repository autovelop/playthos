// +build collision !play

package collision

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Collider defines the position, offset, size, and onHit event
type Collider struct {
	engine.Component
	position *std.Vector3
	offset   *std.Vector2
	size     *std.Vector2
	onHit    func(*Collider)
	reverse  bool
}

// NewCollider creates and sets a new orphan collider
func NewCollider() *Collider {
	return &Collider{}
}

// Set used to define all the require properties of a Collider
func (c *Collider) Set(position *std.Vector3, offset *std.Vector2, size *std.Vector2) {
	c.position = position
	// calculate bounds based on shape & orientation
	c.offset = offset
	c.size = size
}

// Reverse returns whether collider is inverted
//
// TODO(F): Rename Reverse() to Invert() and implement
func (c *Collider) Reverse() bool {
	return c.reverse
}

// SetReverse set/changes whether the collider is reversed
func (c *Collider) SetReverse(r bool) {
	c.reverse = r
}

// OnHit set the onHit event for when two colliders collide
func (c *Collider) OnHit(onHit func(*Collider)) {
	c.onHit = onHit
}

// Hit called by collision system on collision
func (c *Collider) Hit(other *Collider) {
	if c.onHit != nil {
		c.onHit(other)
	}
}

// Get returns all properties that make up a collider
func (c *Collider) Get() (*std.Vector3, *std.Vector2, *std.Vector2) {
	return c.position, c.offset, c.size
}
