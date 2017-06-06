package engine

// type System struct {
// 	SystemRoutine
// }

// func (s *System) RegisterComponent() {
// }

type System interface {
	Update()
	Prepare()
	ComponentTypes() []ComponentRoutine
	LoadComponent(ComponentRoutine)
}
