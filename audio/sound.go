// +build audio

package audio

import (
	"github.com/autovelop/playthos"
)

type Sound struct {
	engine.Component
	audioFile AudioFileRoutine
	loop      bool
}

func (s *Sound) Set(audioFile AudioFileRoutine, loop bool) {
	s.audioFile = audioFile
	s.loop = loop
}

func (s *Sound) IsLoop() bool {
	return s.loop
}

func (s *Sound) Get() AudioFileRoutine {
	return s.audioFile
}
