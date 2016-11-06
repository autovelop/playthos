package gde

type Component struct {
	ComponentRoutine
	Properties map[string]interface{}
}

type ComponentRoutine interface {
	Init()
	GetProperty(string) interface{}
	SetProperty(string, interface{})
}
