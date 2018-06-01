// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

// Source defines the position, sound, and playback properties of a single hearable sound
type Source struct {
	engine.Component
	position    *std.Vector3
	loop        bool
	playing     bool
	playOnReady bool

	basePlaySound Soundable
	PlaySound     func(Soundable)
}

// Soundable interface allows for finding the index of an source
type Soundable interface {
	Index() int
}

// NewSource creates and sets a new orphan sound
func NewSource() *Source {
	s := &Source{}
	s.PlaySound = func(so Soundable) {
		s.basePlaySound = so
	}
	return s
}

// Set used to define all the require properties of a Source
func (s *Source) Set(pos *std.Vector3, l bool, p bool) {
	s.position = pos
	s.loop = l
	s.playOnReady = p
}

// SetPosition sets/changes the source position
func (s *Source) SetPosition(x float32, y float32, z float32) {
	s.position.X = x
	s.position.Y = y
	s.position.Z = z
}

// Loop return whether the sound loops
func (s *Source) Loop() bool {
	return s.loop
}

// BasePlaySound returns the sound of a source
func (s *Source) BasePlaySound() Soundable {
	return s.basePlaySound
}

// PlayOnReady returns the whether the source plays once it is active/created
func (s *Source) PlayOnReady() bool {
	return s.playOnReady
}
