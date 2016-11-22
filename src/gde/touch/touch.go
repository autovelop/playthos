package touch

import (
	"fmt"
	"gde"
)

type Touch struct {
	gde.Input

	touchmap map[int]func(float32, float32)
}

func (i *Touch) Init() {
	i.touchmap = make(map[int]func(float32, float32))
}

func (i *Touch) Update(entities *map[string]gde.EntityRoutine) {
	// fmt.Printf("%v\n", len(*entities))
}

func (i *Touch) Stop() {
}

func (i *Touch) BindAt(key int, callback func(float32, float32)) {
	fmt.Printf("Touch Bind%v\n", key)
	i.touchmap[key] = callback
}

func (i *Touch) Touch(key int, x float32, y float32) {
	fmt.Printf("TouchDown - %v (%v, %v)\n", key, x, y)
	i.touchmap[key](x, y)
}
