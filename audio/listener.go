// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Listener defines the position of which the player hears sounds
type Listener struct {
	engine.Component
	position *std.Vector3
}

// NewListener creates and sets a new orphan listener
func NewListener() *Listener {
	return &Listener{}
}

// Set used to define all the require properties of a Listener
func (l *Listener) Set(pos *std.Vector3) {
	l.position = pos
}

// Position returns the Vector3 position of the listener
func (l *Listener) Position() *std.Vector3 {
	return l.position
}

// SetPosition sets/changes the listener position
func (l *Listener) SetPosition(x float32, y float32, z float32) {
	l.position.X = x
	l.position.Y = y
	l.position.Z = z
}
