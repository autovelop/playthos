package network

import (
	"fmt"
	"gde/engine"
	"net/http"
)

type Network struct {
	engine.System
}

func (a *Network) Init() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7777", nil)
}

func (a *Network) Update(entities *map[string]*engine.Entity) {
	// for _, v := range *entities {
	// animator := v.GetComponent(&Animator{})
	// switch animator := animator.(type) {
	// case AnimatorRoutine:
	// 	animator.StepFrame()
	// }
	// }
}

func (a *Network) Stop() {}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
