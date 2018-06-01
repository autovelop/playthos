// +build deploy audio

package audio

import (
	"github.com/autovelop/playthos"
)

// Audio defines an empty system that will be overwritten by a platform specific system
type Audio struct {
	engine.System
}

// NewAudioSystem used by overriding platform specific system
func NewAudioSystem(audio AudioRoutine) {
	engine.NewSystem(audio)
}

// AudioRoutine interface allows for any platform specific system to be used
type AudioRoutine interface {
	engine.SystemRoutine
	RegisterSource(*Source)
	RegisterSound(*Sound)
	RegisterListener(*Listener)
}
