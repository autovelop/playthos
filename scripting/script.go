// +build scripting !play

package scripting

import (
	"github.com/autovelop/playthos"
)

type Script struct {
	engine.Component
	onUpdate func()
}

func NewScript() *Script {
	return &Script{}
}

func (s *Script) Set() {}

func (s *Script) OnUpdate(onUpdate func()) {
	s.onUpdate = onUpdate
}

func (s *Script) Update() {
	if s.onUpdate != nil {
		s.onUpdate()
	}
}

func (s *Script) Get() {}
