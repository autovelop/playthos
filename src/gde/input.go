package gde

type Input interface {
	SystemRoutine

	BindOn(int, func())
	BindAt(int, func(float64, float64))
	BindMove(func(float64, float64))
}
