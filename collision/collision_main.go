// +build deploy collision

package collision

import (
	"github.com/autovelop/playthos"
	// "log"
)

func init() {
	engine.NewSystem(&Collision{})
}

type Collision struct {
	engine.System
	colliders []*Collider
}

func (c *Collision) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(c.colliders); i++ {
		collider := c.colliders[i]
		if collider.Entity().ID() == entity.ID() {
			copy(c.colliders[i:], c.colliders[i+1:])
			c.colliders[len(c.colliders)-1] = nil
			c.colliders = c.colliders[:len(c.colliders)-1]
		}
	}
}

func (c *Collision) InitSystem() {}

func (c *Collision) Destroy() {}

func (c *Collision) AddIntegrant(integrant engine.IntegrantRoutine) {}

func (c *Collision) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Collider:
		c.colliders = append(c.colliders, component)
		break
	}
}

func (c *Collision) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Collider{}}
}

func (c *Collision) Update() {
	// var prev_collider *Collider
	for a, collider1 := range c.colliders {
		for b, collider2 := range c.colliders {
			if a != b {
				// prev_collider := c.colliders[i-1]
				// if prev_collider == nil {
				// 	prev_collider = collider
				// } else {
				// log.Fatalf("%+v", collider)
				if collider1 == nil {
					continue
				}
				if collider2 == nil {
					continue
				}
				c1 := collider1.Entity()
				c2 := collider2.Entity()
				// continue
				if c1 != nil && c2 != nil {
					// p1 := c1.GetComponent(&render.Transform{}).(*render.Transform)
					// p2 := c2.GetComponent(&render.Transform{}).(*render.Transform)

					if CheckCollisionAABB(collider1, collider2) {
						// collider.Hit(c2)
						// log.Printf("YES %v colliding with %v\n", c1.Tag(), c2.Tag())
						collider2.Hit(c1)
					}
					// } else {
					// 	log.Printf("NO %v colliding with %v\n", c1.Tag(), c2.Tag())
					// }
					// if Distance3(p1.GetPosition(), p2.GetPosition()) < 80 {
					// 	collider.Hit()
					// 	prev_collider.Hit()
					// }
				}
				// }
			}
		}
	}
}
