package gde

// HOW TO USE SYSTEMS:https://play.golang.org/p/0Ab-Ufc5A3

type SystemRoutine interface {
	Init()
	Update(*map[string]EntityRoutine)
	Stop()
}
