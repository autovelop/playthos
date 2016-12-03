package engine

type System interface {
	Init()
	Update(*map[string]*Entity)
	Stop()
}
