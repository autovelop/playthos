package engine

// unit is anything that has an active state
type unit struct {
	engine *Engine
	active bool
	tag    uint
}

// initUnity set unit engine pointer
func (u *unit) initUnit(engine *Engine) {
	u.engine = engine
}

// Active returns unit active state
func (u *unit) Active() bool {
	return u.active
}

// SetActive sets unit active state
func (u *unit) SetActive(active bool) {
	u.active = active
}

// SetTag sets unit tag
func (u *unit) SetTag(t uint) {
	u.tag = t
}

// Tag returns unit tag
func (u *unit) Tag() uint {
	return u.tag
}

// Engine returns unit engine pointer
func (u *unit) Engine() *Engine {
	return u.engine
}
