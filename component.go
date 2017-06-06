package engine

type Component struct {
	active bool
	entity *Entity
	ComponentRoutine
}

func (c *Component) SetActive(active bool) {
	c.active = active
}

func (c *Component) IsActive() bool {
	return c.active
}

func (c *Component) GetEntity() *Entity {
	return c.entity
}

func (c *Component) SetEntity(entity *Entity) {
	c.active = true
	c.entity = entity
}

type ComponentRoutine interface {
	// Base
	GetEntity() *Entity
	SetEntity(*Entity)
	IsActive() bool
	SetActive(bool)

	// To implement
	Prepare()
	RegisterToSystem(System)
	UnRegisterFromSystem(System)
	RegisterToObserverable(Observerable)
	UnRegisterFromObserver(Observerable)
}
