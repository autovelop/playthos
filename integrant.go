package engine

type Integrant struct {
	unit
}

type IntegrantRoutine interface {
	initUnit(*Engine)
	InitIntegrant()
	// 	Engine() *Engine
	// 	Init(*Engine)
	// 	Active() bool
	// 	SetActive(bool)
	// 	Load()
}
