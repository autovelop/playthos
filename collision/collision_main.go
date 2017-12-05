// +build deploy collision

package collision

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/physics"
	"github.com/autovelop/playthos/std"
)

func init() {
	engine.NewSystem(&Collision{})
}

type Collision struct {
	engine.System
	colliders   []*Collider
	transforms  []*std.Transform
	rigidbodies []*physics.RigidBody
}

func (c *Collision) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(c.colliders); i++ {

		collider := c.colliders[i]
		if collider.Entity() != nil {
			if collider.Entity().ID() == entity.ID() {
				copy(c.colliders[i:], c.colliders[i+1:])
				c.colliders[len(c.colliders)-1] = nil
				c.colliders = c.colliders[:len(c.colliders)-1]
			}
		}
	}
}

func (c *Collision) InitSystem() {}

func (c *Collision) Destroy() {}

func (c *Collision) AddIntegrant(integrant engine.IntegrantRoutine) {}

func (c *Collision) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Collider:
		ent := component.Entity()
		if ent == nil {
			return
		}
		trns := ent.Component(&std.Transform{}).(*std.Transform)
		if trns == nil {
			return
		}
		c.colliders = append(c.colliders, component)
		c.transforms = append(c.transforms, trns)
		rbe := ent.Component(&physics.RigidBody{})
		if rbe != nil {
			rb := rbe.(*physics.RigidBody)
			if rb != nil {
				c.rigidbodies = append(c.rigidbodies, rb)
				break
			}
		}
		c.rigidbodies = append(c.rigidbodies, nil)
		break
	}
}

func (c *Collision) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Collider{}}
}

func (c *Collision) Update() {
	length := len(c.colliders) - 1
	for a := 0; a < length; a++ {
		for b := length; b > a; b-- {
			if len(c.colliders)-1 != length {
				break
			}
			c1 := c.colliders[a]
			if c1 == nil {
				continue
			}
			c2 := c.colliders[b]
			if c2 == nil {
				continue
			}
			isCollision := CheckCollisionAABB(c1, c2)
			if isCollision {
				pos := c.transforms[a].Position()
				// if c.rigidbodies[a] != nil {
				// 	rb := c.rigidbodies[a]
				// 	vel := rb.Velocity()
				// 	rb.SetVelocity(vel.X*(1-(rb.Friction()*rb.Mass())), 0, 0)
				// }
				c.transforms[a].SetPosition(pos.X, pos.Y, pos.Z)
				c2.Hit(c1)
			}
		}
		if len(c.colliders)-1 != length {
			break
		}
	}
}
