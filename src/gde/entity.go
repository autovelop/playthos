package gde

import (
	"fmt"
	"log"
)

type Entity struct {
	EntityRoutine
	Id         string
	components map[string]ComponentRoutine
}

func (e *Entity) Init() {
	log.Println("Engine.Init() executed")
	e.components = make(map[string]ComponentRoutine)
}

func (e *Entity) Add(engine *Engine) {
	log.Println("Entity.Add(Engine) executed")
	engine.Entities[e.Id] = e
}

func (e *Entity) Get() *Entity {
	log.Println("Entity.Get() returned Entity")
	return e
}

func (e *Entity) AddComponent(component ComponentRoutine) {
	log.Println("Entity.Component(componentType) returned ComponentRoutine{}")
	e.components[fmt.Sprintf("%T", component)] = component
}

func (e *Entity) Component(componentType string) ComponentRoutine {
	// log.Println("Entity.Component(componentType) returned ComponentRoutine{}")
	return e.components[componentType]
}
func (e *Entity) Components() map[string]ComponentRoutine {
	// log.Println("Entity.Component(componentType) returned ComponentRoutine{}")
	return e.components
}

func (e *Entity) GetId() string {
	log.Println("Entity.GetId() returned string")
	return e.Id
}

type EntityRoutine interface {
	GetId() string
	Get() *Entity
	Add(*Engine)
	AddComponent(ComponentRoutine)
	Component(string) ComponentRoutine
	Components() map[string]ComponentRoutine
	Init()
}
