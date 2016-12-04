package editor

import (
	"gde/engine"
	"gde/render"
	"gde/render/ui"
	"log"
)

// Loads data from file/db
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
	texture.NewTexture("weapon.png")
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

	// First create a UI system
	sys_ui := &ui.UI{Platform: game.GetPlatform()}
	sys_render.AddSubSystem(sys_ui)
	sys_ui.Init()

	// Load a simple font
	font := &ui.Font{}
	font.NewFont()

	ent_box := &engine.Entity{Id: "Box"}
	ent_box.Init()
	ent_box.Add(game)

	ent_box_comp_transform := &render.Transform{}
	ent_box_comp_transform.Init()
	ent_box_comp_transform.SetProperty("Position", render.Vector3{120, 50, 0})
	ent_box_comp_transform.SetProperty("Rotation", render.Vector3{0, 0, 0})
	ent_box_comp_transform.SetProperty("Dimensions", render.Vector3{100, 170, 1})
	ent_box.AddComponent(ent_box_comp_transform)

	comp_ui_renderer := &ui.UIRenderer{}
	comp_ui_renderer.Init()

	text := &ui.Text{}
	text.SetFont(font)
	text.SetText("Hello! Last key pressed: ")
	comp_ui_renderer.SetProperty("Text", text.TextToVec2())
	comp_ui_renderer.SetProperty("TextStart", render.Vector2{10, 10})

	ent_box.AddComponent(comp_ui_renderer)

	sys_ui.LoadRenderer(comp_ui_renderer)

	// Simple UI comp_renderer
	// ui_renderer := &Renderer{}
	// ui_renderer.Init()
	// ui_renderer.LoadMesh(&Mesh{
	// 	Vertices: []float32{
	// 		1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0,
	// 		1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
	// 		0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	// 		0.0, 1.0, 0.0, 0.0, 1.0, 1.0, 0.0, 1.0,
	// 	},
	// 	Indicies: []uint8{
	// 		0, 1, 3,
	// 		1, 2, 3,
	// 	},
	// })
	// render.LoadRenderer(ui_renderer)

	// // Creatae a text box
	// heading := &Entity{Id: "Heading"}
	// heading.Init()
	// heading.Add(engine)

	// text_transform := &Transform{}
	// text_transform.Init()
	// text_transform.SetProperty("Position", Vector3{0.0, 0.0, 0})

	// // Simple Quad mesh text renderer
	// text_renderer := &TextRenderer{}
	// text_renderer.Init()
	// text_renderer.LoadMesh(&Mesh{
	// 	Vertices: []float32{
	// 		1.0, 2.0, -0.1, 1.0, 0.0, 0.0,
	// 		1.0, 0.0, -0.1, 0.0, 1.0, 0.0,
	// 		0.0, 0.0, -0.1, 0.0, 0.0, 1.0,
	// 		0.0, 2.0, -0.1, 0.0, 1.0, 1.0,
	// 	},
	// 	Indicies: []uint8{
	// 		0, 1, 3,
	// 		1, 2, 3,
	// 	},
	// })

	// font := &Font{}
	// font.NewFont()
	// text := &Text{"Hello! Last key pressed: ", font}
	// text_renderer.SetProperty("Text", text.TextToVec2())

	// render.LoadTextRenderer(text_renderer)

	// heading.AddComponent(text_transform)
	// heading.AddComponent(text_renderer)

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
