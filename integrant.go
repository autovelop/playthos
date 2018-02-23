package engine

type Integrant struct {
	unit
}

type IntegrantRoutine interface {
	initUnit(*Engine)
	InitIntegrant()
	AddIntegrant(IntegrantRoutine)
	SetActive(bool)
	Destroy()
}

type Listener interface {
	IntegrantRoutine
	On(int, func(...int))
	IsSet(int) bool
	Emit(int, int)
}

type Platformer interface {
	IntegrantRoutine
	LoadAsset(string)
	Asset(string) []byte
}
