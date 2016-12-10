package animation

import (
	"fmt"
	"gde/engine"
)

type Animator struct {
	engine.Component
	frame    int
	EndFrame int
	Start    func(int)
	Step     func(int)
}
type AnimatorRoutine interface {
	engine.ComponentRoutine
	StepFrame()
	StartFrame()
}

func (a *Animator) Init() {
	fmt.Println("Animator.Init() executed")
	a.Properties = make(map[string]interface{})
}

func (a *Animator) GetProperty(key string) interface{} {
	return a.Properties[key]
}

func (a *Animator) SetProperty(key string, val interface{}) {
	a.Properties[key] = val
}

func (a *Animator) StartFrame() {
	a.frame = 0
	a.Start(a.frame)
}
func (a *Animator) StepFrame() {
	a.frame++
	if a.frame > a.EndFrame {
		a.StartFrame()
	} else {
		a.Step(a.frame)
	}
}
