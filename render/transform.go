package render

import (
	"./../engine"
	"log"
)

type Transform struct {
	position Vector3
	rotation Vector3
	scale    Vector3
}

func (t *Transform) RegisterToSystem(system engine.System) {
	log.Println("Registering Transform")
	switch system := system.(type) {
	case Render:
		system.RegisterTransform(t)
	}
}

func (t *Transform) Set(pos Vector3, rot Vector3, scl Vector3) {
	t.position = pos
	t.rotation = rot
	t.scale = scl
}

func (t *Transform) GetPosition() Vector3 {
	return t.position
}

func (t *Transform) GetRotation() Vector3 {
	return t.rotation
}

func (t *Transform) GetScale() Vector3 {
	return t.scale
}

// type Transform struct {
// 	engine.Component
// }

// func (t *Transform) Init() {
// 	// fmt.Println("Transform.Init() executed")
// 	t.Properties = make(map[string]interface{})
// }

// func (t *Transform) Id() string {
// 	return "Transform"
// }

// func (t *Transform) GetProperty(key string) interface{} {
// 	return t.Properties[key]
// }

// func (t *Transform) SetProperty(key string, val interface{}) {
// 	t.Properties[key] = val
// }
