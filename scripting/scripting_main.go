// +build deploy scripting

package scripting

import (
	"github.com/autovelop/playthos"
)

func init() {
	engine.NewSystem(&Scripting{})
}

// Scripting system used to update script components
type Scripting struct {
	engine.System
	scripts []*Script
}

// InitSystem called when the system plugs into the engine
func (s *Scripting) InitSystem() {}

// Destroy called when engine is gracefully shutting down
func (s *Scripting) Destroy() {}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (s *Scripting) AddIntegrant(integrant engine.IntegrantRoutine) {}

// AddComponent unorphans a component by adding it to this system
func (s *Scripting) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Script:
		s.scripts = append(s.scripts, component)
		break
	}
}

// DeleteEntity removes all entity's compoents from this system
func (s *Scripting) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(s.scripts); i++ {
		script := s.scripts[i]
		if script.Entity().ID() == entity.ID() {
			copy(s.scripts[i:], s.scripts[i+1:])
			s.scripts[len(s.scripts)-1] = nil
			s.scripts = s.scripts[:len(s.scripts)-1]
		}
	}
}

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (s *Scripting) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Script{}}
}

// Update called by engine to progress this system to the next engine loop
func (s *Scripting) Update() {
	for _, script := range s.scripts {
		script.Update()
	}
}
