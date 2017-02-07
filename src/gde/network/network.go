package network

import (
	"gde/engine"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	// "time"
)

type Network struct {
	engine.System
	// lastUpdate time.Time
}

// var addr = flag.String("addr", "localhost:8080", "http service address")

func (a *Network) Init() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "192.168.1.104:8080", Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// go func() {
	// 	ticker := time.NewTicker(time.Second)
	// 	defer ticker.Stop()
	// 	for {
	// 		select {
	// 		case t := <-ticker.C:
	// 			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	// 			if err != nil {
	// 				log.Println("write:", err)
	// 				return
	// 			}
	// 		case <-interrupt:
	// 			log.Println("interrupt")
	// 			// To cleanly close a connection, a client should send a close
	// 			// frame and wait for the server to close the connection.
	// 			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// 			if err != nil {
	// 				log.Println("write close:", err)
	// 				return
	// 			}
	// 			select {
	// 			case <-done:
	// 			case <-time.After(time.Second):
	// 			}
	// 			c.Close()
	// 			return
	// 		}
	// 	}
	// 	defer c.Close()
	// }()
}

func (a *Network) Update(entities *map[string]*engine.Entity) {
	// if time.Since(a.lastUpdate).Seconds() >= 1.0 {
	// 	a.lastUpdate = time.Now()
	// } else {
	// 	return
	// }
	// res, err := http.Get("http://192.168.1.104:7777")
	// if err != nil {
	// 	log.Printf("ERR: %v", err)
	// }

	// defer res.Body.Close()

	// var body struct {
	// 	// httpbin.org sends back key/value pairs, no map[string][]string
	// 	Value float64 `json:"value"`
	// }
	// json.NewDecoder(res.Body).Decode(&body)

	// for _, v := range *entities {
	// 	if v.Id == "Box" {
	// 		v.GetComponent(&render.Transform{}).SetProperty("Position", render.Vector3{120 * float32(body.Value), 200, 0.5})
	// 	}
	// }
	// log.Printf("%v", body.Value)

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
