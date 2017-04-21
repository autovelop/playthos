package engine

type Component interface {
	Prepare()
	RegisterToSystem(System)
}
