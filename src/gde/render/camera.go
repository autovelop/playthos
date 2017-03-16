package render

import (
	"gde/engine"
	// "log"
)

type Camera struct {
	engine.Component
}

func (c *Camera) Init() {
	// log.Printf("Camera > Init")
	c.Properties = make(map[string]interface{})
}

func (c *Camera) GetProperty(key string) interface{} {
	// log.Printf("Camera > Property > Get: %v", key)
	return c.Properties[key]
}

func (c *Camera) SetProperty(key string, val interface{}) {
	// log.Printf("Camera > Property > Set: %v", key)
	c.Properties[key] = val
}
