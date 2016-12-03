package input

import (
	"gde/engine"
)

type InputListener interface {
	engine.System

	BindOn(int, func())
	BindAt(int, func(float64, float64))
	BindMove(func(float64, float64))
}
