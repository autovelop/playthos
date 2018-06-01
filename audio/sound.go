// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
)

// Sound is a component used to take any platform specific audio clip and play it
type Sound struct {
	engine.Component

	baseClip Clipable
	SetClip  func(Clipable)

	index int
}

// Set used to define all the require properties of a Sound
func (s *Sound) Set(c Clipable) {
	s.baseClip = c
}

// NewSound creates and sets a new orphan sound
func NewSound() *Sound {
	s := &Sound{}
	s.SetClip = func(c Clipable) {
		s.baseClip = c
	}
	s.index = -1
	return s
}

// BaseClip returns the clip of a sound
func (s *Sound) BaseClip() Clipable {
	return s.baseClip
}

// SetIndex sets/changes the index of a sound
func (s *Sound) SetIndex(i int) {
	s.index = i
}

// Index returns this sounds index
func (s *Sound) Index() int {
	return s.index
}
