package engine

// NOTE
// https://play.golang.org/p/JUpq7Tsf8P

type Component struct {
	unit
	id     string
	entity *Entity
}

func (c *Component) Entity() *Entity {
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
