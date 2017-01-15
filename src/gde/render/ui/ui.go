package ui

import (
	"gde/engine"
	"gde/render"
)

type UI struct {
	render.RenderRoutine
	ShaderProgram uint32
	Platform      *engine.Platform
	mesh          *render.Mesh
}

type UIRoutine interface {
	Init()
	LoadRenderer(render.RendererRoutine)
	NewShader(string, string) uint32
	AddSubSystem(render.RenderRoutine)
	Stop()
	Update(*map[string]*engine.Entity)
}

func (u *UI) Init() {}

func (u *UI) AddSubSystem(system render.RenderRoutine) {}

func (u *UI) Update(entities *map[string]*engine.Entity) {}

func (u *UI) Stop() {}

func (r *UI) LoadRenderer(renderer render.RendererRoutine) {}
