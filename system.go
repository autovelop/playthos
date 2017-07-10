package engine

type System struct {
	unit
}

type SystemRoutine interface {
	initUnit(*Engine)
	InitSystem()
	NewComponent(ComponentRoutine)
	NewIntegrant(IntegrantRoutine)
	// *Unit
	// SystemRoutine
	// active bool
	// engine *Engine
	ComponentTypes() []ComponentRoutine
}

// func (s *System) Engine() *Engine {
// 	return s.engine
// }

// func (s *System) Init(engine *Engine) {
// 	s.engine = engine
// }

// func (s *System) Active() bool {
// 	return s.active
// }

// func (s *System) SetActive(active bool) {
// 	s.active = active
// }

// type SystemRoutine interface {
// 	Engine() *Engine
// 	Init(*Engine)
// 	//base
// 	// LoadComponent(ComponentRoutine)
// 	ComponentTypes() []ComponentRoutine

// 	// implemented
// 	NewComponent(ComponentRoutine)
// 	NewIntegrant(IntegrantRoutine)
// 	DeleteEntity(*Entity)

// 	// Update()
// 	// UnloadComponent(ComponentRoutine)
// 	// Prepare(*Settings)
// 	// UnRegisterEntity(*Entity)
// 	Active() bool
// 	SetActive(bool)
// 	Load()
// }

type Updater interface {
	SystemRoutine
	Update()
	// UpdaterRoutine
}

// type UpdaterRoutine interface {
// 	update()
// }

type Listener interface {
	SystemRoutine
	// System
	On(uint, func(...uint))
}
