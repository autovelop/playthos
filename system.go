package engine

type System struct {
	unit
}

type SystemRoutine interface {
	initUnit(*Engine)
	InitSystem()
	Destroy()
	AddComponent(ComponentRoutine)
	DeleteEntity(*Entity)
	AddIntegrant(IntegrantRoutine)
	ComponentTypes() []ComponentRoutine
	SetActive(bool)
	// Stop()
}

type Drawer interface {
	SystemRoutine
	Draw()
}

type Updater interface {
	SystemRoutine
	Update()
}
