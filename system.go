package engine

type System interface {
	Update()
	Prepare()
	ComponentTypes() []Component
	NewComponentType(Component)
}
