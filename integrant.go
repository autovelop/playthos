package engine

type Integrant struct {
	unit
}

type IntegrantRoutine interface {
	initUnit(*Engine)
	InitIntegrant()
	DeleteIntegrant()
}
