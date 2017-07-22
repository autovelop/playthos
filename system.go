package engine

type System struct {
	unit
}

type SystemRoutine interface {
	initUnit(*Engine)
	InitSystem()
	AddComponent(ComponentRoutine)
	DeleteEntity(*Entity)
	AddIntegrant(IntegrantRoutine)
	ComponentTypes() []ComponentRoutine
	SetActive(bool)
	// Stop()
}

type Updater interface {
	SystemRoutine
	Update()
}

type Listener interface {
	SystemRoutine
	On(uint, func(...uint))
}
