package engine_test

import (
	"github.com/autovelop/playthos"
	// "testing"
)

// TODO(F): Write System tests

type MockUpdaterSystem struct{ engine.System }

func (m *MockUpdaterSystem) InitSystem()                                    {}
func (m *MockUpdaterSystem) Destroy()                                       {}
func (m *MockUpdaterSystem) DeleteEntity(entity *engine.Entity)             {}
func (m *MockUpdaterSystem) Update()                                        {}
func (m *MockUpdaterSystem) AddComponent(component engine.ComponentRoutine) {}
func (m *MockUpdaterSystem) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}
func (m *MockUpdaterSystem) AddIntegrant(integrant engine.IntegrantRoutine) {}

type MockDrawerSystem struct{ engine.System }

func (m *MockDrawerSystem) InitSystem()                                    {}
func (m *MockDrawerSystem) Destroy()                                       {}
func (m *MockDrawerSystem) DeleteEntity(entity *engine.Entity)             {}
func (m *MockDrawerSystem) Draw()                                          {}
func (m *MockDrawerSystem) AddComponent(component engine.ComponentRoutine) {}
func (m *MockDrawerSystem) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{}
}
func (m *MockDrawerSystem) AddIntegrant(integrant engine.IntegrantRoutine) {}
