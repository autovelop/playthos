// +build render
package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Camera struct {
	engine.Component
	eye    *std.Vector3
	center *std.Vector3
	up     *std.Vector3
}

func (c *Camera) SetEye(eye *std.Vector3) {
	c.eye = eye
}

func (c *Camera) SetCenter(center *std.Vector3) {
	c.center = center
}

func (c *Camera) SetUp(up *std.Vector3) {
	c.up = up
}

func (c *Camera) Eye() *std.Vector3 {
	return c.eye
}

func (c *Camera) Center() *std.Vector3 {
	return c.center
}

func (c *Camera) Up() *std.Vector3 {
	return c.up
}
