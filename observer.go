package engine

type Observer interface {
	Prepare()
	// Register(uint32, interface{})
	// Fire(uint32)
	// ComponentTypes() []Component
}
