// +build script

package script

import (
	"github.com/autovelop/playthos"
)

type CustomScript struct {
	engine.Component
	onUpdate func()
}

func (c *CustomScript) Set() {}

func (c *CustomScript) OnUpdate(onUpdate func()) {
	c.onUpdate = onUpdate
}

func (c *CustomScript) Update() {
	if c.onUpdate != nil {
		c.onUpdate()
	}
}

func (c *CustomScript) Get() {}
