// +build physics

package physics

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	// "log"
)

type Acceleration struct {
	engine.Component
	std.Vector3
}

func NewAcceleration() *Acceleration {
	return &Acceleration{}
}

func (a *Acceleration) Set(x float32, y float32, z float32) {
	a.X = x
	a.Y = y
	a.Z = z
}
