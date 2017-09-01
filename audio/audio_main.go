// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
	// "io"
	// "log"
	// "time"
)

type Audio struct {
	engine.System
}

func NewAudioSystem(audio AudioRoutine) {
	engine.NewSystem(audio)
}

type AudioRoutine interface {
	engine.SystemRoutine
	RegisterSource(*Source)
	RegisterSound(*Sound)
	RegisterListener(*Listener)
}
