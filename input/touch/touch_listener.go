package touch

import (
	"gde/engine"
	"gde/input"
)

type Touch struct {
	input.InputListener

	touchmap map[int]func(float32, float32)
}

func (i *Touch) Init() {
	i.touchmap = make(map[int]func(float32, float32))
}

func (i *Touch) Update(entities *map[string]*engine.Entity) {
	// TODO
}

func (i *Touch) Stop() {
	// TODO
}

func (i *Touch) BindAt(key int, callback func(float32, float32)) {
	// TODO
	// fmt.Printf("Touch Bind%v\n", key)
	// i.touchmap[key] = callback
}

func (i *Touch) Touch(key int, x float32, y float32) {
	// TODO
	// fmt.Printf("TouchDown - %v (%v, %v)\n", key, x, y)
	// i.touchmap[key](x, y)
}
