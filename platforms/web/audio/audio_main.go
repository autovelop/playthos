// +build deploy webaudio

package audio

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"

	"github.com/autovelop/playthos/platforms/web"
	// for now we always pull in wav if openal is used until other audio file decoders exist
	// _ "github.com/autovelop/playthos/glfw"
	// "bytes"
	// "golang.org/x/mobile/exp/audio/al"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	// "io"
	// "log"
)

func init() {
	audio.NewAudioSystem(&WebAudio{})
	fmt.Println("> Audio (Web) Added")
}

type WebAudio struct {
	audio.Audio
	platform *web.Web
	context  *js.Object

	sources  []*WebAudioSource
	sounds   []*WebAudioSound
	listener *audio.Listener
	settings *engine.Settings
}

func (a *WebAudio) InitSystem() {
	// al.OpenDevice()
	if js.Global.Get("AudioContext") != nil {
		a.context = js.Global.Get("AudioContext").New()
	} else {
		a.context = js.Global.Get("webkitAudioContext").New()
	}
}

func (a *WebAudio) Destroy() {
	// al.CloseDevice()
}

func (a *WebAudio) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	case *web.Web:
		a.platform = integrant
		break
	}
}

func (a *WebAudio) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *audio.Source:
		a.RegisterSource(component)
		break
	case *audio.Sound:
		a.RegisterSound(component)
		break
	case *audio.Listener:
		a.RegisterListener(component)
		break
	}
}

func (a *WebAudio) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&audio.Source{}, &audio.Sound{}, &audio.Listener{}}
}

func (a *WebAudio) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(a.sounds); i++ {
		clip := a.sounds[i]
		if clip.Entity().ID() == entity.ID() {
			copy(a.sources[i:], a.sources[i+1:])
			a.sources[len(a.sources)-1] = nil
			a.sources = a.sources[:len(a.sources)-1]

			copy(a.sounds[i:], a.sounds[i+1:])
			a.sounds[len(a.sounds)-1] = nil
			a.sounds = a.sounds[:len(a.sounds)-1]
		}
	}
}

func (a *WebAudio) RegisterSource(s *audio.Source) {
	webAudioSource := &WebAudioSource{Source: s}
	webAudioSource.OverridePlaySound(func(s audio.Soundable) {
		if a.sounds[s.Index()] != nil {
			soundSource := a.sounds[s.Index()].Buffer()
			soundSource.Set("loop", webAudioSource.Loop)
			soundSource.Get("mediaElement").Call("play")
		}
	})
	a.sources = append(a.sources, webAudioSource)
}

func (a *WebAudio) RegisterSound(c *audio.Sound) {
	clip := c.BaseClip()
	webAudioSound := &WebAudioSound{Sound: c}
	if clip != nil {
		webAudioSound.OverrideClip(func(c audio.Clipable) {
			raw := a.platform.Asset(c.Path())
			// fmt.Println(a.context)
			// soundSource := a.context.Call("createBufferSource")
			soundSource := a.context.Call("createMediaElementSource", raw)
			// a.context.Call("decodeAudioData", , func(buffer *js.Object) {
			webAudioSound.buffer = soundSource
			// soundSource.Set("buffer", buffer)
			// })
			soundSource.Call("connect", a.context.Get("destination"))
			// raw.Call("play")
			// fmt.Println(soundSource)
			// 			openALClip := &OpenALClip{c.(*audio.Clip), raw}
			// 			openALClip.Decode()
			// 			webAudioSound.clip = openALClip

			// 			buffer := &al.GenBuffers(1)[0]

			// 			r := bytes.NewReader(openALClip.data)

			// 			s := make([]byte, openALClip.Length())
			// 			size := int64(0)
			// 			for {
			// 				n, err := r.Read(s)
			// 				if n > 0 {
			// 					size += int64(n)
			// 					buffer.BufferData(al.FormatStereo16, s[:n], int32(openALClip.SampleRate()))
			// 				}
			// 				if err == io.EOF {
			// 					break
			// 				}
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 			}

			// 			webAudioSound.buffer = buffer
		})
	}
	a.sounds = append(a.sounds, webAudioSound)
	webAudioSound.SetIndex(len(a.sounds) - 1)

	// if webAudioSource.PlayOnReady() {
	// 	openALSound.Play()
	// }
}

func (a *WebAudio) RegisterListener(l *audio.Listener) {
	// pos := l.Position()
	// al.SetListenerPosition([3]float32{pos.X, pos.Y, pos.Z})
	// a.listener = l
}

// func (a *WebAudio) AddIntegrant(integrant engine.IntegrantRoutine) {
// 	switch integrant := integrant.(type) {
// 	}
// }
