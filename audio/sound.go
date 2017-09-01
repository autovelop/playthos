// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
)

type Sound struct {
	engine.Component

	baseClip Clipable
	SetClip  func(Clipable)

	index int
}

func (s *Sound) Set(c Clipable) {
	s.baseClip = c
}

func NewSound() *Sound {
	s := &Sound{}
	s.SetClip = func(c Clipable) {
		s.baseClip = c
	}
	s.index = -1
	return s
}

// func (s *Sound) Play(p Playable) {
// }

func (s *Sound) BaseClip() Clipable {
	return s.baseClip
}

func (s *Sound) SetIndex(i int) {
	s.index = i
}

func (s *Sound) Index() int {
	return s.index
}
