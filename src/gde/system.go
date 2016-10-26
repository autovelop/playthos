package gde

type SystemRoutine interface {
	Init()
	Update()
	End()
	Add(*Engine)
	Property(string) interface{}
}
