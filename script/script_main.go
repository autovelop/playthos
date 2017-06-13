// +build script

package script

import (
	"github.com/autovelop/playthos"
	"log"
)

var componentTypes []engine.ComponentRoutine = []engine.ComponentRoutine{&CustomScript{}}

func init() {
	script := &Script{}
	engine.NewSystem(script)
}

type Script struct {
	engine.System
	scripts []*CustomScript
}

func (s *Script) Prepare() {
}

func (s *Script) LoadComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *CustomScript:
		s.RegisterCustomScript(component)
		log.Println("LoadComponent(*CustomScript)")
		break
	}
}

func (s *Script) ComponentTypes() []engine.ComponentRoutine {
	return componentTypes
}

func (s *Script) Update() {
	for _, script := range s.scripts {
		script.Update()
	}
}

func (s *Script) RegisterCustomScript(script *CustomScript) {
	s.scripts = append(s.scripts, script)
}

func (s *Script) UnRegisterEntity(entity *engine.Entity) {
	for i := 0; i < len(s.scripts); i++ {
		script := s.scripts[i]
		if script.GetEntity().ID == entity.ID {
			copy(s.scripts[i:], s.scripts[i+1:])
			s.scripts[len(s.scripts)-1] = nil
			s.scripts = s.scripts[:len(s.scripts)-1]
		}
	}
}
