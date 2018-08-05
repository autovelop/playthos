// +build deploy render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Camera defines the origin, lookat, and up vectors of the camera. Also contains the viewport data (zoom, clear color)
type Camera struct {
	engine.Component
	eye    *std.Vector3
	center *std.Vector3
	up     *std.Vector3
	window *std.Vector2
	scale  *float32
	clear  *std.Color
}

// NewCamera creates and sets a new orphan camera
func NewCamera() *Camera {
	return &Camera{}
}

// Set used to define all the require properties of a Camera
func (c *Camera) Set(s *float32, cl *std.Color) {
	c.scale = s
	c.clear = cl
	e := c.Entity()
	if e == nil {
		return
	}
	t := std.GetTransform(e)
	if t == nil {
		return
	}
	c.eye = t.Position()
	c.center = t.Rotation()
	c.up = t.Scale()
	// c.up = up
	// c.center = center
}

// SetTransform set/changes all the vector data values
//
// TODO(F): Don't use transform and rather create its own component
func (c *Camera) SetTransform(t *std.Transform) {
	c.eye = t.Position()
	c.center = t.Rotation()
	c.up = t.Scale()
}

// SetWindow set/changes the window scale values
func (c *Camera) SetWindow(w float32, h float32) {
	v := &std.Vector2{w, h}
	c.window = v
	// log.Fatal(v)
}

// PointToRay converts a window position to a world position (Vector2)
func (c *Camera) PointToRay(x float32, y float32) *std.Vector2 {
	v := &std.Vector2{
		(2*x)/c.window.X - 1,
		1 - (2*y)/c.window.Y,
	}
	e := c.Entity()
	if e == nil {
		return nil
	}
	t := std.GetTransform(e)
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

// Eye returns the camera origin vector
func (c *Camera) Eye() *std.Vector3 {
	return c.eye
}

// ClearColor returns the camera clear color
func (c *Camera) ClearColor() *std.Color {
	return c.clear
}

// Scale returns the camera scale value
func (c *Camera) Scale() *float32 {
	return c.scale
}

// Center returns the camera lookat vector
func (c *Camera) Center() *std.Vector3 {
	return c.center
}

// Up returns the camera up vector
func (c *Camera) Up() *std.Vector3 {
	return c.up
}
