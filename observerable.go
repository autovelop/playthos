package engine

// type Observerable struct {
// 	ObserverableRoutine
// }

type Observerable interface {
	Prepare()
	LoadComponent(ComponentRoutine)
}
