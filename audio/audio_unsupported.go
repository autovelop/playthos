// +build audio
// +build !linux
// +build windows
// +build darwin

package audio

import (
	"github.com/autovelop/playthos"
)

func init() {
	engine.NewSystem(&Audio{})
}

type Audio struct {
	engine.System
	sounds []*Sound
}

func (a *Audio) InitSystem() {}

func NewSound() *Sound {
	return &Sound{}
}
func (a *Audio) DeleteEntity(entity *engine.Entity) {}

func (a *Audio) StopSound(sound *Sound) {}

func (a *Audio) DeleteSound(sound *Sound) {}

func (a *Audio) PlaySound(sound *Sound) {}

func (a *Audio) NewComponent(sound engine.ComponentRoutine) {}

func (a *Audio) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}

func (a *Audio) NewIntegrant(integrant engine.IntegrantRoutine) {}
