package gde

type Input struct {
}

type InputRoutine interface {
	SystemRoutine

	Bind(int, func())
}
