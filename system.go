package main

type System interface {
	Update()
	Prepare()
	ComponentTypes() []Component
}

// type System interface {
// 	Init()
// 	Update(*map[string]*Entity)
// 	Stop()
// }
