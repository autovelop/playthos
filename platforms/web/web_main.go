// +build deploy web

package web

import (
	"fmt"
	"github.com/autovelop/playthos"
	"github.com/gopherjs/gopherjs/js"
	"time"
	// "go/build"
	// "io/ioutil"
	// "os"
	"strings"
)

func init() {
	engine.NewIntegrant(&Web{})
	fmt.Println("> Web: Ready")
}

type Web struct {
	engine.Integrant
	assets  map[string]*js.Object
	Loaded  func()
	waiting int
}

func (w *Web) AddIntegrant(engine.IntegrantRoutine) {}

func (w *Web) InitIntegrant() {
	w.assets = make(map[string]*js.Object, 0)
	js.Global.Call("addEventListener", "load", func() {
		go func() {
			ready := make(chan bool)
			go func() {
				for {
					if w.waiting <= 0 {
						ready <- true
					}
					time.Sleep(time.Millisecond * 100)
				}
			}()
			<-ready
			w.Loaded()
		}()
	}, false)
}

func (w *Web) Destroy() {}

func (w *Web) Asset(p string) *js.Object {
	return w.assets[p]
}

func (w *Web) LoadAsset(p string) {
	ready := make(chan bool)
	dotSplit := strings.Split(p, ".")
	ext := dotSplit[len(dotSplit)-1]
	switch ext {
	case "png":
		imageFile := js.Global.Get("Image").New()
		imageFile.Set("onload", func(s *js.Object) {
			w.assets[p] = imageFile
			ready <- true
		})
		imageFile.Set("src", p)
		break
	case "wav":
		audioFile := js.Global.Get("Audio").New()
		audioFile.Set("oncanplaythrough", func(s *js.Object) {
			w.assets[p] = audioFile
			ready <- true
		})
		audioFile.Set("src", p)
		// audioFile.Set("autoplay", true)
		break
	}
	w.waiting++
	<-ready
	w.waiting--
}
