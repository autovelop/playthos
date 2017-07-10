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

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) Set(eye *std.Vector3, center *std.Vector3, up *std.Vector3) {
	c.eye = eye
	c.up = up
	c.center = center
}

func (c *Camera) SetEye(x float32, y float32, z float32) {
	c.eye.X = x
	c.eye.Y = y
	c.eye.Z = z
}

func (c *Camera) SetCenter(x float32, y float32, z float32) {
	c.center.X = x
	c.center.Y = y
	c.center.Z = z
}

func (c *Camera) SetUp(x float32, y float32, z float32) {
	c.up.X = x
	c.up.Y = y
	c.up.Z = z
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
