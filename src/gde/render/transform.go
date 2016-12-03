package render

import (
	"fmt"
	"gde/engine"
)

type Transform struct {
	engine.Component
}

func (t *Transform) Init() {
	fmt.Println("Transform.Init() executed")
	t.Properties = make(map[string]interface{})
}

func (t *Transform) GetProperty(key string) interface{} {
	return t.Properties[key]
}

func (t *Transform) SetProperty(key string, val interface{}) {
	t.Properties[key] = val
}
