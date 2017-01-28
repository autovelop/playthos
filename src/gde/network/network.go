package network

import (
	"encoding/json"
	"gde/engine"
	"gde/render"
	"log"
	"net/http"
	"time"
)

type Network struct {
	engine.System
	lastUpdate time.Time
}

func (a *Network) Init() {
	a.lastUpdate = time.Now()
}

func (a *Network) Update(entities *map[string]*engine.Entity) {
	if time.Since(a.lastUpdate).Seconds() >= 1.0 {
		a.lastUpdate = time.Now()
	} else {
		return
	}
	res, err := http.Get("http://192.168.1.104:7777")
	if err != nil {
		log.Printf("ERR: %v", err)
	}

	defer res.Body.Close()

	var body struct {
		// httpbin.org sends back key/value pairs, no map[string][]string
		Value float64 `json:"value"`
	}
	json.NewDecoder(res.Body).Decode(&body)

	for _, v := range *entities {
		if v.Id == "Box" {
			v.GetComponent(&render.Transform{}).SetProperty("Position", render.Vector3{120 * float32(body.Value), 200, 0.5})
		}
	}
	log.Printf("%v", body.Value)

	// log.Println(res.Body)
	// decoder := json.NewDecoder(res.Body)
	// err = decoder.Decode(&data)
	// for _, v := range *entities {
	// animator := v.GetComponent(&Animator{})
	// switch animator := animator.(type) {
	// case AnimatorRoutine:
	// 	animator.StepFrame()
	// }
	// }
}

func (a *Network) Stop() {}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there!")
// 	log.Printf("Hi there!")
// }
