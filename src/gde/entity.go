package gde

import "fmt"

type Entity struct {
	EntityRoutine
	Id         string
	Components map[string]ComponentRoutine
}

func (e *Entity) Init() {
	fmt.Println("Engine.Init() executed")
	e.Components = make(map[string]ComponentRoutine)
}

func (e *Entity) Add(engine *Engine) {
	fmt.Println("Entity.Add(Engine) executed")
	engine.Entities[e.Id] = e
}

func (e *Entity) Get() *Entity {
	fmt.Println("Entity.Get() returned Entity")
	return e
}

func (e *Entity) GetId() string {
	fmt.Println("Entity.GetId() returned string")
	return e.Id
}

type EntityRoutine interface {
	GetId() string
	Get() *Entity
	Add(*Engine)
	Init()
}
