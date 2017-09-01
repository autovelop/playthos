// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Source struct {
	engine.Component
	position    *std.Vector3
	loop        bool
	playing     bool
	playOnReady bool

	basePlaySound Soundable
	PlaySound     func(Soundable)
}

type Soundable interface {
	Index() int
}

func NewSource() *Source {
	s := &Source{}
	s.PlaySound = func(so Soundable) {
		s.basePlaySound = so
	}
	return s
}

func (s *Source) Set(pos *std.Vector3, l bool, p bool) {
	s.position = pos
	s.loop = l
	s.playOnReady = p
	// won't actually play it
	// s.PlaySound(s)
}

func (s *Source) SetPosition(x float32, y float32, z float32) {
	s.position.X = x
	s.position.Y = y
	s.position.Z = z
}

func (s *Source) Loop() bool {
	return s.loop
}

func (s *Source) BasePlaySound() Soundable {
	return s.basePlaySound
}

func (s *Source) PlayOnReady() bool {
	return s.playOnReady
}

// func (s *Source) PlaySound(so *Sound) {
// }

// type Playable interface {
// 	Play(Playable)
// }
