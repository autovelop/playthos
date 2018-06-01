package engine

// System used to procedurally execute with engine.
type System struct {
	unit
}

// SystemRoutine plugs into engine to procedurally alter components.
type SystemRoutine interface {
	initUnit(*Engine)
	InitSystem()
	Destroy()
	AddComponent(ComponentRoutine)
	DeleteEntity(*Entity)
	AddIntegrant(IntegrantRoutine)
	ComponentTypes() []ComponentRoutine
	SetActive(bool)
}

// Drawer interface used on components that render on screen.
type Drawer interface {
	SystemRoutine
	Draw()
}

// Updater interface used on components that update every game loop.
type Updater interface {
	SystemRoutine
	Update()
}
