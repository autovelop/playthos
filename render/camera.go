// +build deploy render

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
	window *std.Vector2
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
	e := c.Entity()
	if e == nil {
		return
	}
	t := e.Component(&std.Transform{}).(*std.Transform)
	if t == nil {
		return
	}
	c.eye = t.Position()
	c.center = t.Rotation()
	c.up = t.Scale()
	// c.up = up
	// c.center = center
}

func (c *Camera) SetTransform(t *std.Transform) {
	c.eye = t.Position()
	c.center = t.Rotation()
	c.up = t.Scale()
}

// func (c *Camera) SetCenter(x float32, y float32, z float32) {
// 	c.center.X = x
// 	c.center.Y = y
// 	c.center.Z = z
// }

func (c *Camera) SetWindow(w float32, h float32) {
	v := &std.Vector2{w, h}
	c.window = v
	// log.Fatal(v)
}

func (c *Camera) PointToRay(x float32, y float32) *std.Vector2 {
	v := &std.Vector2{
		(2*x)/c.window.X - 1,
		1 - (2*y)/c.window.Y,
	}
	e := c.Entity()
	if e == nil {
		return nil
	}
	t := e.Component(&std.Transform{}).(*std.Transform)
	if t == nil {
		return nil
	}
	v.X = t.Position().X + v.X*(c.window.X/2)
	v.Y = t.Position().Y + v.Y*(c.window.Y/2)
	// transform.SetPosition(camera_transform.Position().X+ray.X*400, camera_transform.Position().Y+ray.Y*300, 1.0)
	// float x = (2.0f * mouse_x) / width - 1.0f;
	// float y = 1.0f - (2.0f * mouse_y) / height;
	return v
}

func (c *Camera) Eye() *std.Vector3 {
	return c.eye
}

func (c *Camera) ClearColor() *std.Color {
	return c.clear
}

func (c *Camera) Scale() *float32 {
	return c.scale
}

func (c *Camera) Center() *std.Vector3 {
	return c.center
}

func (c *Camera) Up() *std.Vector3 {
	return c.up
}
