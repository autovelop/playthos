package engine

import (
	"testing"
)

// Trivial test example
func (e *Engine) TestInit(t *testing.T) {
	e.Entities = make(map[string]*Entity)
	e.Systems = make(map[int]System)
	e.Entities["one"] = &Entity{}
	e.Systems[1] = &SimpleSystem{}
}

type SimpleSystem struct {
}

type SimpleSystemRoutine interface {
	System
}

func (s *SimpleSystem) Init()                               {}
func (s *SimpleSystem) Stop()                               {}
func (s *SimpleSystem) Update(entities *map[string]*Entity) {}
