// +build scripting !play

package scripting

import (
	"github.com/autovelop/playthos"
)

// Script defines the update function
type Script struct {
	engine.Component
	onUpdate func()
}

// NewScript creates and sets a new orphan script
func NewScript() *Script {
	return &Script{}
}

// Set used to define all the required properties
func (s *Script) Set() {}

// OnUpdate sets/changes the update function
func (s *Script) OnUpdate(onUpdate func()) {
	s.onUpdate = onUpdate
}

// Update executes the update function of script (if set)
func (s *Script) Update() {
	if s.onUpdate != nil {
		s.onUpdate()
	}
}
