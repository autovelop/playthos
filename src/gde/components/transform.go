package components

import (
	"fmt"

	"gde"
)

type Transform struct {
	gde.ComponentRoutine
}

func (t *Transform) Add(entity *gde.Entity) {
	fmt.Println("Transform.Add(Entity) executed")
	entity.Components[fmt.Sprintf("%T", Transform{})] = t
}
