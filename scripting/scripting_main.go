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

func (s *Scripting) NewIntegrant(integrant engine.IntegrantRoutine) {}

func (s *Scripting) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Script:
		s.scripts = append(s.scripts, component)
		break
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
