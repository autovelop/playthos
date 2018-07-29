package engine

// Integrant used to sporadically execute with engine.
type Integrant struct {
	unit
}

// IntegrantRoutine plugs into engine to sporadically alter components.
type IntegrantRoutine interface {
	initUnit(*Engine)
	InitIntegrant()
	AddIntegrant(IntegrantRoutine)
	SetActive(bool)
	Destroy()
}

// Listener interface used to listen and emit engine events.
type Listener interface {
	IntegrantRoutine
	On(int, func(...int))
	IsSet(int) bool
	Emit(int, int)
}

// Platformer interface used to load and retrieve platform specific assets.
type Platformer interface {
	IntegrantRoutine
	IsDeploy()
	LoadAsset(string)
}

// Desktoper interface used to load Assets for desktops and return []byte
type Desktoper interface {
	Platformer
	Asset(string) []byte
}
