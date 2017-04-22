package engine

type Subject interface {
	Prepare()
	LoadComponent(Component)
}
