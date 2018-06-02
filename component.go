package engine

// Component is added to an entity and used by systems to simulate the scene
type Component struct {
	unit
	id     string
	entity *Entity
}

// Entity returns component's entity
func (c *Component) Entity() *Entity {
	if c.entity == nil {
		return nil
	}
	return c.entity
}

// initComponent used by the engine to assign a entity to each component
func (c *Component) initComponent(ent *Entity) {
	c.entity = ent
}

// ID returns component id
func (c *Component) ID() string {
	return c.id
}

// ComponentRoutine interface allows for generic component types to be handled by the engine
type ComponentRoutine interface {
	initUnit(*Engine)
	initComponent(*Entity)
	SetActive(bool)
}
