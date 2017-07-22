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
	scale  *float32
	clear  *std.Color
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) Set(s *float32, cl *std.Color) {
	// func (c *Camera) Set(s *float32, center *std.Vector3, up *std.Vector3) {
	c.scale = s
	c.clear = cl
	// c.up = up
	// c.center = center
}

// func (c *Camera) SetCenter(x float32, y float32, z float32) {
// 	c.center.X = x
// 	c.center.Y = y
// 	c.center.Z = z
// }

// func (c *Camera) SetUp(x float32, y float32, z float32) {
// 	c.up.X = x
// 	c.up.Y = y
// 	c.up.Z = z
// }

func (c *Camera) Eye() *std.Vector3 {
	e := c.Entity()
	if e == nil {
		return nil
	}
	t := e.Component(&std.Transform{}).(*std.Transform)
	if t == nil {
		return nil
	}
	return t.Position()
}

func (c *Camera) ClearColor() *std.Color {
	return c.clear
}

func (c *Camera) Scale() *float32 {
	return c.scale
}

func (c *Camera) Center() *std.Vector3 {
	e := c.Entity()
	if e == nil {
		return nil
	}
	t := e.Component(&std.Transform{}).(*std.Transform)
	if t == nil {
		return nil
	}
	return t.Rotation()
	// return c.center
}

func (c *Camera) Up() *std.Vector3 {
	e := c.Entity()
	if e == nil {
		return nil
	}
	t := e.Component(&std.Transform{}).(*std.Transform)
	if t == nil {
		return nil
	}
	return t.Scale()
	// return c.up
}
