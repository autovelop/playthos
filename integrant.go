package engine

type Integrant struct {
	unit
}

type IntegrantRoutine interface {
	initUnit(*Engine)
	InitIntegrant()
	SetActive(bool)
	Destroy()
}

type Platformer interface {
	IntegrantRoutine
	LoadAsset(string)
}
