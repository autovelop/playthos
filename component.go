package engine

type Component struct {
	unit
	id     string
	entity *Entity
}

func (c *Component) Entity() *Entity {
	if c.entity == nil {
		return nil
	}
	return c.entity
}

func (c *Component) initComponent(ent *Entity) {
	c.entity = ent
}

func (c *Component) ID() string {
	return c.id
}

type ComponentRoutine interface {
	initUnit(*Engine)
	initComponent(*Entity)
	SetActive(bool)
}
