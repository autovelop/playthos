// +build physics

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Velocity struct {
	engine.Component
	std.Vector3
}

func (v *Velocity) Set(x float32, y float32, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v *Velocity) SetX(x float32) {
	v.X = x
}

func (v *Velocity) SetY(y float32) {
	v.Y = y
}

func (v *Velocity) Add(x float32, y float32, z float32) {
	v.X += x
	v.Y += y
	v.Z += z
}
