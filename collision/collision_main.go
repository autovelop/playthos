// +build collision

package collision

import (
	"github.com/autovelop/playthos"
	// render "github.com/autovelop/playthos-render"
	"log"
)

var componentTypes []engine.ComponentRoutine = []engine.ComponentRoutine{&Collider{}}

func init() {
	col := &Collision{}
	engine.NewSystem(col)
}

type Collision struct {
	engine.System
	colliders []*Collider
}

func (c *Collision) Prepare() {
}

func (c *Collision) LoadComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Collider:
		c.RegisterCollider(component)
		log.Println("LoadComponent(*Collider)")
		break
	}
}

func (c *Collision) ComponentTypes() []engine.ComponentRoutine {
	return componentTypes
}

func (c *Collision) Update() {
	var prev_collider *Collider
	for _, collider := range c.colliders {
		if prev_collider == nil {
			prev_collider = collider
		} else {
			c1 := collider.GetEntity()
			c2 := prev_collider.GetEntity()
			if c1 != nil && c2 != nil {
				// p1 := c1.GetComponent(&render.Transform{}).(*render.Transform)
				// p2 := c2.GetComponent(&render.Transform{}).(*render.Transform)

				if CheckCollisionAABB(collider, prev_collider) {
					collider.Hit()
					prev_collider.Hit()
				}
				// if Distance3(p1.GetPosition(), p2.GetPosition()) < 80 {
				// 	collider.Hit()
				// 	prev_collider.Hit()
				// }
			}
		}
	}
}

func (c *Collision) RegisterCollider(collider *Collider) {
	c.colliders = append(c.colliders, collider)
}

func (c *Collision) UnRegisterEntity(entity *engine.Entity) {
	for i := 0; i < len(c.colliders); i++ {
		collider := c.colliders[i]
		if collider.GetEntity().ID == entity.ID {
			copy(c.colliders[i:], c.colliders[i+1:])
			c.colliders[len(c.colliders)-1] = nil
			c.colliders = c.colliders[:len(c.colliders)-1]
		}
	}
}
