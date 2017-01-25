package editor

import (
	"gde/engine"
	// "gde/network"
	"gde/render"

	// "gde/render/animation"
	"gde/render/ui"
	"log"
)

type Scene struct {
	name string
}

func (s *Scene) LoadScene(game *engine.Engine) {
	sys_render, err := game.GetSystem(engine.SystemRender).(render.RenderRoutine)
	if !err {
		log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
		return
	}

	// Simple Quad mesh renderer
	comp_renderer := &render.MeshRenderer{}
	comp_renderer.Init()
	comp_renderer.LoadMesh(&render.Mesh{
		Vertices: []float32{
			0.2, 0.2, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0,
			0.2, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
			0.0, 0.2, 0.0, 0.0, 1.0, 1.0, 0.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	texture := &render.Texture{}
	texture.NewTextureMobile("weapon.png")
	comp_renderer.LoadTexture(texture)

	sys_render.LoadRenderer(comp_renderer)

	// Create player entity
	ent_player := &engine.Entity{Id: "Player"}
	ent_player.Init()
	ent_player.Add(game)

	ent_player_comp_transform := &render.Transform{}
	ent_player_comp_transform.Init()
	ent_player_comp_transform.SetProperty("Position", render.Vector3{0.5, 1.0, 0})
	ent_player_comp_transform.SetProperty("Rotation", render.Vector3{0, 0, 45})

	ent_player.AddComponent(ent_player_comp_transform)
	ent_player.AddComponent(comp_renderer)

	// Create UI entity

	// network := &network.Network{}
	// game.AddSystem(engine.SystemNetwork, network)
	// network.Init()

	// // First create a UI system
	sys_render.AddUISystem(game)
	sys_ui, err := game.GetSystem(engine.SystemUI).(ui.UIRoutine)
	if !err {
		log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
		return
	}
	log.Printf("UI SYSTEM: %+v", sys_ui)

	// Load a simple font
	font := &ui.Font{}
	font.NewFont()

	ent_box := &engine.Entity{Id: "Box"}
	ent_box.Init()
	ent_box.Add(game)

	ent_box_comp_transform := &render.Transform{}
	ent_box_comp_transform.Init()
	ent_box_comp_transform.SetProperty("Position", render.Vector3{0.35, 0.8, 0})
	ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{100, 100})
	ent_box.AddComponent(ent_box_comp_transform)

	comp_ui_renderer := &ui.UIRenderer{}
	comp_ui_renderer.Init()

	text := &ui.Text{}
	text.SetFont(font)
	text.SetText(`Common Sword`)
	comp_ui_renderer.SetProperty("Text", text.TextToVec4())
	// comp_ui_renderer.SetProperty("Scale", 2.0)
	// comp_ui_renderer.SetProperty("Padding", render.Vector4{10 /*top*/, 10 /*right*/, 10 /*bottom*/, 10 /*left*/})

	ent_box.AddComponent(comp_ui_renderer)

	sys_ui.LoadRenderer(comp_ui_renderer)

	// Officially the worst Animation system since forever!
	// sys_anim := &animation.Animation{}
	// sys_anim.Init()
	// game.AddSystem(engine.SystemAnimation, sys_anim)

	// ent_box_comp_animator := &animation.Animator{EndFrame: 240, Start: func(frame int) {
	// 	ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{80, 600})
	// }, Step: func(frame int) {
	// 	dimensions := ent_box_comp_transform.GetProperty("Dimensions")
	// 	switch dimensions := dimensions.(type) {
	// 	case render.Vector2:
	// 		ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{dimensions.X + 1, dimensions.Y})
	// 	}
	// }}
	// ent_box_comp_animator.Init()
	// ent_box.AddComponent(ent_box_comp_animator)

	// // // Lets test keyboard support
	// keyInput, err := engine.GetSystem(SystemInputKeyboard).(Input)
	// if !err {
	// 	log.Println(err)
	// 	return
	// }

	// // Right arrow
	// keyInput.BindOn(262, func() {
	// 	box2_position.X += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// 	textTmp := &Text{text.GetText() + "Right", font}
	// 	text_renderer.SetProperty("Text", textTmp.TextToVec2())
	// })

	// // Left arrow
	// keyInput.BindOn(263, func() {
	// 	box2_position.X -= 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Up arrow
	// keyInput.BindOn(265, func() {
	// 	box2_position.Y -= 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Down arrow
	// keyInput.BindOn(264, func() {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Lets test pointer support
	// pointerInput, err := engine.GetSystem(SystemInputPointer).(Input)
	// if !err {
	// 	log.Printf("Pointer Input system not started/found\nERROR:\n%v\n\n", err)
	// 	return
	// }

	// // pointer Move
	// pointerInput.BindMove(func(x float64, y float64) {
	// 	box2_position.X = float32(x/360) * 1
	// 	box2_position.Y = float32(y/640) * 2
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// Left click
	// pointerInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Right click
	// pointerInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Lets test touch support
	// touchInput, err := engine.GetSystem(SystemInputTouch).(InputRoutine)
	// if !err {
	// 	log.Printf("Touch Input system not started/found\nERROR:\n%v\n\n", err)
	// 	return
	// }

	// // Touch down
	// touchInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Touch up
	// touchInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })
}
