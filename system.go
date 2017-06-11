package engine

type System interface {
	Update()
	Prepare()
	ComponentTypes() []ComponentRoutine
	LoadComponent(ComponentRoutine)
}
