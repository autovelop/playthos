package input

import (
	"fmt"
	"gde"
)

type Pointer struct {
	gde.Input

	buttondown func()
	buttonup   func()
}

func (i *Pointer) Init() {
}

func (i *Pointer) Update(entities *map[string]gde.EntityRoutine) {
	// fmt.Printf("%v\n", len(*entities))
}

func (i *Pointer) Stop() {
}

func (i *Pointer) Bind(key string, callback func()) {
	fmt.Printf("%v\n", key)
}

func (i *Pointer) ButtonDown(key string) {
	fmt.Printf("%v\n", key)
}

func (i *Pointer) ButtonUp(key string) {
	fmt.Printf("%v\n", key)
}
