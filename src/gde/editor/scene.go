package editor

import (
	"gde/engine"
	// "gde/network"
	// "encoding/json"
	input "gde/input"
	"gde/render"

	// "gde/render/animation"
	"fmt"
	// "gde/render/ui"
	// "github.com/gorilla/websocket"
	"log"
	// "net"
	// "net/http"
	// "strconv"
)

type Scene struct {
	name                string
	Game                *engine.Engine
	RenderSystem        render.RenderRoutine
	KeyboardInputSystem input.InputListener
}

// var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
// 	return true
// }}

type EditorAction struct {
	Action uint
	Data   string
}

type EditorUpdate struct {
	Entity      string
	Component   string
	Property    string
	SubProperty string
	Value       string
}

// func (s *Scene) CreateEditorServer(game *engine.Engine) {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")

// 		ws, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Print("upgrade:", err)
// 			return
// 		}
// 		// defer c.Close()
// 		for {
// 			mt, message, err := ws.ReadMessage()
// 			if err != nil {
// 				log.Println("read:", err)
// 				break
// 			}

// 			var editor_action EditorAction
// 			err_unmarshal := json.Unmarshal(message, &editor_action)
// 			if err_unmarshal != nil {
// 				fmt.Println("error:", err_unmarshal)
// 				fmt.Println("error:", string(message))
// 			}

// 			switch editor_action.Action {
// 			case 0:
// 				entity_json, err := json.Marshal(game.GetEntity("Player"))
// 				if err != nil {
// 					fmt.Println("error:", err)
// 				}

// 				err = ws.WriteMessage(mt, entity_json)
// 				if err != nil {
// 					log.Println("write:", err)
// 					break
// 				}
// 			case 1:
// 				fmt.Printf("%v", editor_action.Data)
// 				var editor_update EditorUpdate
// 				err_update_unmarshal := json.Unmarshal([]byte(editor_action.Data), &editor_update)
// 				if err_update_unmarshal != nil {
// 					fmt.Println("error:", err)
// 				}
// 				// fmt.Printf("%+v", game.GetEntity(editor_update.Entity))
// 				fmt.Printf("%+v", game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(editor_update.Property))
// 				switch editor_update.Component {
// 				case "*render.Transform":
// 					{
// 						vec3 := game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(editor_update.Property)
// 						switch vec3 := vec3.(type) {
// 						case render.Vector3:
// 							switch editor_update.SubProperty {
// 							case "X":
// 								v, err := strconv.ParseFloat(editor_update.Value, 64)
// 								if err != nil {
// 									fmt.Println("error:", err)
// 								}
// 								vec3.X = float32(v)
// 								break
// 							case "Y":
// 								v, err := strconv.ParseFloat(editor_update.Value, 64)
// 								if err != nil {
// 									fmt.Println("error:", err)
// 								}
// 								vec3.Y = float32(v)
// 								break
// 							case "Z":
// 								v, err := strconv.ParseFloat(editor_update.Value, 64)
// 								if err != nil {
// 									fmt.Println("error:", err)
// 								}
// 								vec3.Z = float32(v)
// 								break
// 							}
// 							game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).SetProperty(editor_update.Property, vec3)
// 						}
// 					}
// 				}

// 				// game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(
// 				// var vec3 render.Vector3
// 				// err_vec3_unmarshal := json.Unmarshal([]byte(editor_update.Value), &vec3)
// 				// if err_vec3_unmarshal != nil {
// 				// 	fmt.Println("error:", err_vec3_unmarshal)
// 				// }
// 			}

// 			// log.Printf("recv: %s", message)
// 			// err = ws.WriteMessage(mt, message)
// 			// if err != nil {
// 			// 	log.Println("write:", err)
// 			// 	break
// 			// }
// 		}

// 		// GET ON CONNECT EVENT
// 		// SEND JSON OF ENTITIES TO CLIENT
// 		// ALLOW CLIENTS TO SEND NEW VALUES OF PROPERTIES BY PROPERTY NAME AND ENTITY ID AS JSON VALUE (SHOULD BE CONVERTED ACCORDINGLY)
// 		// BROADCAST UPDATE TO ALL CLIENTS

// 		// ticker := time.NewTicker(time.Second)
// 		// defer ticker.Stop()
// 		// for {
// 		// 	select {
// 		// 	case t := <-ticker.C:
// 		// 		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
// 		// 		if err != nil {
// 		// 			log.Println("write:", err)
// 		// 			return
// 		// 		}
// 		// 	}
// 		// }
// 	})
// 	var ip net.IP
// 	ifaces, err := net.Interfaces()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	for _, i := range ifaces {
// 		addrs, err := i.Addrs()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		for _, addr := range addrs {
// 			switch v := addr.(type) {
// 			case *net.IPNet:
// 				ip = v.IP
// 			case *net.IPAddr:
// 				ip = v.IP
// 			}
// 		}
// 	}
// 	address := "192.168.43.1:8080"
// 	if len(fmt.Sprintf("%v", ip)) < 25 {
// 		address = fmt.Sprintf("%v:8080", ip)
// 	}
// 	go func(address string) {
// 		fmt.Println(address)
// 		log.Fatal(http.ListenAndServe(address, nil))
// 	}(address)

// 	// THIS IS NOT REALLY REQUIRED BUT JUST USED TO DEBUGGING ON LOCAL MACHINE (testing broadcast etc.)
// 	// network := &network.Network{ServerIP: address}
// 	// game.AddSystem(engine.SystemNetwork, network)
// 	// network.Init()
// }

func (s *Scene) NewGameObject(id string, mesh *render.Mesh, color render.Color, position render.Vector3, rotation render.Vector3, scale render.Vector3) {
	// Create player entity
	entity := &engine.Entity{Id: id}
	entity.Init()
	entity.Add(s.Game)

	renderer := &render.MeshRenderer{}
	renderer.Init()
	renderer.LoadMesh(mesh)
	renderer.SetColor(&color)

	s.RenderSystem.LoadRenderer(renderer)

	entity.AddComponent(renderer)

	transform := &render.Transform{}
	transform.Init()
	transform.SetProperty("Position", position)
	transform.SetProperty("Rotation", rotation)
	transform.SetProperty("Scale", scale)
	entity.AddComponent(transform)
}

func (s *Scene) NewTextureGameObject(id string, mesh *render.Mesh, texture *render.Texture, color render.Color, position render.Vector3, rotation render.Vector3, scale render.Vector3) {
	// Create player entity
	entity := &engine.Entity{Id: id}
	entity.Init()
	entity.Add(s.Game)

	renderer := &render.MeshRenderer{}
	renderer.Init()
	renderer.LoadMesh(mesh)
	renderer.SetColor(&color)

	renderer.LoadTexture(texture)

	s.RenderSystem.LoadRenderer(renderer)

	entity.AddComponent(renderer)

	transform := &render.Transform{}
	transform.Init()
	transform.SetProperty("Position", position)
	transform.SetProperty("Rotation", rotation)
	transform.SetProperty("Scale", scale)
	entity.AddComponent(transform)

}

func (s *Scene) MoveEntity(id string, direction *render.Vector3) {
	ent := s.Game.GetEntity(id)
	trans := ent.GetComponent(&render.Transform{})
	pos := trans.GetProperty("Position")
	switch pos := pos.(type) {
	case render.Vector3:
		pos.Add(direction)
		trans.SetProperty("Position", pos)
		log.Printf("Scene > Player > Position: %v", pos)
	}
}

func (s *Scene) LoadEngine(game *engine.Engine) {
	s.Game = game
}

func (s *Scene) LoadScene() {
	sys_render, err := s.Game.GetSystem(engine.SystemRender).(render.RenderRoutine)
	if !err {
		log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
		return
	}
	s.RenderSystem = sys_render

	quad := &render.Mesh{
		Vertices: []float32{
			0.2, 0.2, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0,
			0.2, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 0.0, 0.0,
			0.0, 0.2, 0.0, 1.0, 1.0, 1.0, 0.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	}

	sys_render.GetCamera().SetProperty("LookAt", render.Vector3{0, 0, 1})
	sys_render.GetCamera().SetProperty("LookFrom", render.Vector3{0, 0, 0})

	sys_keyboard_input, err := s.Game.GetSystem(engine.SystemInputKeyboard).(input.InputListener)
	if !err {
		log.Println(err)
		return
	}
	s.KeyboardInputSystem = sys_keyboard_input

	const (
		KEY_DOWN  = 264
		KEY_UP    = 265
		KEY_LEFT  = 263
		KEY_RIGHT = 262
		KEY_SPACE = 32
	)

	// Game starts here

	s.NewGameObject("Player",
		quad,
		render.Color{1, 0, 0, 0.2},
		render.Vector3{300, 300, 0.4},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	sys_keyboard_input.BindOn(
		KEY_DOWN,
		func() {
			s.MoveEntity("Player", &render.Vector3{0, 60, 0})
		},
	)
	sys_keyboard_input.BindOn(
		KEY_UP,
		func() {
			s.MoveEntity("Player", &render.Vector3{0, -60, 0})
		},
	)
	sys_keyboard_input.BindOn(
		KEY_LEFT,
		func() {
			s.MoveEntity("Player", &render.Vector3{-60, 0, 0})
		},
	)
	sys_keyboard_input.BindOn(
		KEY_RIGHT,
		func() {
			s.MoveEntity("Player", &render.Vector3{60, 0, 0})
		},
	)

	sys_keyboard_input.BindOn(
		KEY_SPACE,
		func() {
			ent := s.Game.GetEntity("Player")
			trans := ent.GetComponent(&render.Transform{})
			pos := trans.GetProperty("Position")
			switch pos := pos.(type) {
			case render.Vector3:
				ent_selector := s.Game.GetEntity("Selector")
				if ent_selector == nil {
					s.NewGameObject("Selector",
						quad,
						render.Color{0, 1, 0, 0.2},
						pos,
						render.Vector3{0, 0, 0},
						render.Vector3{3, 3, 3},
					)
				} else {
					s.Game.DeleteEntity("Selector")
				}
			}
		},
	)

	for letter_idx, letter_val := range [8]string{"A", "B", "C", "D", "E", "F", "G", "H"} {
		for number_idx, number_val := range [8]uint{1, 2, 3, 4, 5, 6, 7, 8} {
			col := render.Color{0.15, 0.15, 0.15, 1}
			if letter_idx%2 == 0 {
				if number_idx%2 == 0 {
					col = render.Color{0.85, 0.85, 0.85, 1}
				}
			} else {
				col = render.Color{0.85, 0.85, 0.85, 1}
				if number_idx%2 == 0 {
					col = render.Color{0.15, 0.15, 0.15, 1}
				}
			}
			// if number_idx%2 > 0 {
			// 	col = render.Color{0.85, 0.85, 0.85, 1}
			// }
			s.NewGameObject(fmt.Sprintf("%v%v", letter_val, number_val),
				quad,
				col,
				render.Vector3{float32(letter_idx * 60), float32(number_idx * 60), 0},
				render.Vector3{0, 0, 0},
				render.Vector3{3, 3, 3},
			)
		}
	}

	pawn_texture := &render.Texture{}
	pawn_texture.NewTexture("assets", "pawn.png")

	s.NewTextureGameObject("Pawn_W_A",
		quad,
		pawn_texture,
		render.Color{0, 0, 0, 1.0},
		render.Vector3{0, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_B",
		quad,
		pawn_texture,
		render.Color{1, 1, 1, 1.0},
		render.Vector3{60, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_C",
		quad,
		pawn_texture,
		render.Color{0, 0, 0, 1.0},
		render.Vector3{120, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_D",
		quad,
		pawn_texture,
		render.Color{1, 1, 1, 1.0},
		render.Vector3{180, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_E",
		quad,
		pawn_texture,
		render.Color{0, 0, 0, 1.0},
		render.Vector3{240, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_F",
		quad,
		pawn_texture,
		render.Color{1, 1, 1, 1.0},
		render.Vector3{300, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_G",
		quad,
		pawn_texture,
		render.Color{0, 0, 0, 1.0},
		render.Vector3{360, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	s.NewTextureGameObject("Pawn_W_H",
		quad,
		pawn_texture,
		render.Color{1, 1, 1, 1.0},
		render.Vector3{420, 360, 0.2},
		render.Vector3{0, 0, 0},
		render.Vector3{3, 3, 3},
	)

	// Create UI entity

	// // First create a UI system
	// sys_render.AddUISystem(game)
	// sys_ui, err := game.GetSystem(engine.SystemUI).(ui.UIRoutine)
	// if !err {
	// 	log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
	// 	return
	// }
	// log.Printf("UI SYSTEM: %+v", sys_ui)

	// // Load a simple font
	// font := &ui.Font{}
	// font.NewFont()

	// ent_box := &engine.Entity{Id: "Box"}
	// ent_box.Init()
	// ent_box.Add(game)

	// ent_box_comp_transform := &render.Transform{}
	// ent_box_comp_transform.Init()
	// ent_box_comp_transform.SetProperty("Position", render.Vector3{120, 200, 0.5})
	// ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{240, 400})
	// ent_box.AddComponent(ent_box_comp_transform)

	// comp_ui_renderer := &ui.UIRenderer{}
	// comp_ui_renderer.Init()

	// text := &ui.Text{}
	// text.SetFont(font)
	// text.SetText(`Common Sword`)
	// comp_ui_renderer.SetProperty("Text", text.TextToVec4())
	// comp_ui_renderer.SetProperty("Scale", 2.0)
	// comp_ui_renderer.SetProperty("Padding", render.Vector4{0 /*top*/, 0 /*right*/, 0 /*bottom*/, 0 /*left*/})

	// ent_box.AddComponent(comp_ui_renderer)

	// sys_ui.LoadRenderer(comp_ui_renderer)

	// s.CreateEditorServer(game)

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

	// // Right arrow
	// keyInput.BindOn(262, func() {
	// 	pos := ent_player_comp_transform.GetProperty("Position")
	// 	switch pos := pos.(type) {
	// 	case render.Vector3:
	// 		pos.X += 60.0
	// 		ent_player_comp_transform.SetProperty("Position", pos)
	// 		log.Printf("Scene > Player > Position: %v", pos)
	// 	}
	// })

	// // Left arrow
	// keyInput.BindOn(263, func() {
	// 	pos := ent_player_comp_transform.GetProperty("Position")
	// 	switch pos := pos.(type) {
	// 	case render.Vector3:
	// 		pos.X -= 60.0
	// 		ent_player_comp_transform.SetProperty("Position", pos)
	// 		log.Printf("Scene > Player > Position: %v", pos)
	// 	}
	// })

	// // Up arrow
	// keyInput.BindOn(265, func() {
	// 	pos := ent_player_comp_transform.GetProperty("Position")
	// 	switch pos := pos.(type) {
	// 	case render.Vector3:
	// 		pos.Y -= 60.0
	// 		ent_player_comp_transform.SetProperty("Position", pos)
	// 		log.Printf("Scene > Player > Position: %v", pos)
	// 	}
	// })

	// ctrl_down := false
	// // Down arrow
	// keyInput.BindOn(264, func() {
	// 	if ctrl_down == false {
	// 		pos := ent_player_comp_transform.GetProperty("Position")
	// 		switch pos := pos.(type) {
	// 		case render.Vector3:
	// 			pos.Y += 60.0
	// 			ent_player_comp_transform.SetProperty("Position", pos)
	// 			log.Printf("Scene > Player > Position: %v", pos)
	// 		}
	// 	} else {
	// 		lookat := sys_render.GetCamera().GetProperty("LookAt")
	// 		switch lookat := lookat.(type) {
	// 		case render.Vector3:
	// 			lookat.Y += 0.05
	// 			sys_render.GetCamera().SetProperty("LookAt", lookat)
	// 			// ent_player_comp_transform.SetProperty("Position", pos)
	// 			// log.Printf("Scene > Player > Position: %v", pos)
	// 		}
	// 	}
	// })

	// // Ctrl down
	// keyInput.BindOnHold(341, func() {
	// 	ctrl_down = true
	// }, func() {
	// 	ctrl_down = false
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
