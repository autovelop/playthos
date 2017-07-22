// +build scripting

package scripting

import (
	"github.com/autovelop/playthos"
	// "log"
)

func init() {
	engine.NewSystem(&Scripting{})
}

type Scripting struct {
	engine.System
	scripts []*Script
}

func (s *Scripting) InitSystem() {}

func (s *Scripting) AddIntegrant(integrant engine.IntegrantRoutine) {}

func (s *Scripting) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Script:
		s.scripts = append(s.scripts, component)
		break
	}
}

func (s *Scripting) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(s.scripts); i++ {
		collider := s.scripts[i]
		if collider.Entity().ID() == entity.ID() {
			copy(s.scripts[i:], s.scripts[i+1:])
			s.scripts[len(s.scripts)-1] = nil
			s.scripts = s.scripts[:len(s.scripts)-1]
		}
	}
}

func (s *Scripting) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Script{}}
}

func (s *Scripting) Update() {
	for _, script := range s.scripts {
		script.Update()
	}
}
