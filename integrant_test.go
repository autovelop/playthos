package engine_test

import (
	"github.com/autovelop/playthos"
	// "testing"
)

// TODO(F): Write integrant tests

type MockIntegrant struct {
	engine.Integrant
}

func (m *MockIntegrant) InitIntegrant()                       {}
func (m *MockIntegrant) AddIntegrant(engine.IntegrantRoutine) {}
func (m *MockIntegrant) Destroy()                             {}

type MockListenerIntegrant struct {
	engine.Integrant
}

func (m *MockListenerIntegrant) InitIntegrant()                       {}
func (m *MockListenerIntegrant) AddIntegrant(engine.IntegrantRoutine) {}
func (m *MockListenerIntegrant) Destroy()                             {}
func (m *MockListenerIntegrant) On(int, func(...int))                 {}
func (m *MockListenerIntegrant) IsSet(int) bool                       { return false }
func (m *MockListenerIntegrant) Emit(int, int)                        {}

type MockPlatformerIntegrant struct {
	engine.Integrant
}

func (m *MockPlatformerIntegrant) InitIntegrant()                       {}
func (m *MockPlatformerIntegrant) AddIntegrant(engine.IntegrantRoutine) {}
func (m *MockPlatformerIntegrant) Destroy()                             {}
func (m *MockPlatformerIntegrant) LoadAsset(string)                     {}
func (m *MockPlatformerIntegrant) Asset(string) []byte                  { return []byte{} }
