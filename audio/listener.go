// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Listener struct {
	engine.Component
	position *std.Vector3
}

func NewListener() *Listener {
	return &Listener{}
}

func (l *Listener) Set(pos *std.Vector3) {
	l.position = pos
}

func (l *Listener) Position() *std.Vector3 {
	return l.position
}

func (l *Listener) SetPosition(x float32, y float32, z float32) {
	l.position.X = x
	l.position.Y = y
	l.position.Z = z
}
