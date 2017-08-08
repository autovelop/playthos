package engine

type unit struct {
	engine *Engine
	active bool
	tag    uint
}

func (u *unit) Active() bool {
	return u.active
}

func (u *unit) SetActive(active bool) {
	u.active = active
}

func (u *unit) SetTag(t uint) {
	u.tag = t
}

func (u *unit) Tag() uint {
	return u.tag
}

func (u *unit) Engine() *Engine {
	return u.engine
}

func (u *unit) initUnit(engine *Engine) {
	u.engine = engine
}
