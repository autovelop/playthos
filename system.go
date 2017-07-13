package engine

type System struct {
	unit
}

type SystemRoutine interface {
	initUnit(*Engine)
	InitSystem()
	NewComponent(ComponentRoutine)
	DeleteEntity(*Entity)
	NewIntegrant(IntegrantRoutine)
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
