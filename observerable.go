package engine

type Observerable interface {
	Prepare()
	LoadComponent(ComponentRoutine)
}
