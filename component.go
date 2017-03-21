package main

type Component interface {
	RegisterToSystem(System)
}

// type Component struct {
// 	Properties map[string]interface{}
// }

// // Many better ways of doing properties but this should suffice for now
// type ComponentRoutine interface {
// 	Id() string
// 	Init()
// 	GetProperty(string) interface{}
// 	SetProperty(string, interface{})
// }
