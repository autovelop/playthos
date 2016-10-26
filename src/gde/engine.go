package gde

import "fmt"

type Engine struct {
	Entities map[string]EntityRoutine
	Systems  map[string]SystemRoutine
}

func (e *Engine) Init() {
	fmt.Println("Engine.Init() executed")
	e.Entities = make(map[string]EntityRoutine)
	e.Systems = make(map[string]SystemRoutine)
}
func (e *Engine) Update() {
	fmt.Println("Engine.Update() executed")
	for _, v := range e.Systems {
		v.Update()
	}
}
func (e *Engine) GetEntity(id string) EntityRoutine {
	fmt.Println("Engine.GetEntity(string) returned EntityRoutine")
	return e.Entities[id]
}
func (e *Engine) GetSystem(id interface{}) SystemRoutine {
	fmt.Println("Engine.GetSystem(interface{}) returned SystemRoutine")
	return e.Systems[fmt.Sprintf("%T", id)]
}
