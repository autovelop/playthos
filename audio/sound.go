// +build autovelop_playthos_audio !play

package audio

import (
	"github.com/autovelop/playthos"
)

type Sound struct {
	engine.Component
	audioFile   AudioFileRoutine
	loop        bool
	audioSystem *Audio
	idx         int
	ready       bool
	playing     bool
	playOnReady bool
}

func (s *Sound) Set(audioFile AudioFileRoutine, loop bool, playOnReady bool) {
	s.audioFile = audioFile
	s.loop = loop
	s.playOnReady = playOnReady
}

func (s *Sound) Play() {
	s.audioSystem.PlaySound(s)
}

func (s *Sound) Stop() {
	s.audioSystem.StopSound(s)
}

func (s *Sound) Loops() bool {
	return s.loop
}

func (s *Sound) Get() AudioFileRoutine {
	return s.audioFile
}
