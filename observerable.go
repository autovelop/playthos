package engine

type Observerable interface {
	Prepare(*Settings)
	LoadComponent(ComponentRoutine)
	UnRegisterEntity(*Entity)
}
