package gde

type SystemRoutine interface {
	Init()
	Update(*map[string]EntityRoutine)
	End()
	Add(*Engine)
	Property(string) interface{}
}
