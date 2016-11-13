package components

import (
	"fmt"

	"gde"
)

type Transform struct {
	gde.Component
	// gde.ComponentRoutine
}

func (r *Transform) Init() {
	fmt.Println("Transform.Init() executed")
	r.Properties = make(map[string]interface{})

}

func (r *Transform) GetProperty(key string) interface{} {
	return r.Properties[key]
}

func (r *Transform) SetProperty(key string, val interface{}) {
	r.Properties[key] = val
}
