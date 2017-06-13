package engine

type System interface {
	Update()
	Prepare(*Settings)
	ComponentTypes() []ComponentRoutine
	LoadComponent(ComponentRoutine)
	// UnloadComponent(ComponentRoutine)
	UnRegisterEntity(*Entity)
}
