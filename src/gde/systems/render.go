package systems

import (
	"gde"
	"gde/components"
)

type Render struct {
	OpenGL string
	GLSL   string
}

type RenderRoutine interface {
	gde.SystemRoutine

	LoadRenderer(*components.Renderer)
}

func (r *Render) Init()                                         {}
func (r *Render) End()                                          {}
func (r *Render) Add(engine *gde.Engine)                        {}
func (r *Render) Update(entities *map[string]gde.EntityRoutine) {}
