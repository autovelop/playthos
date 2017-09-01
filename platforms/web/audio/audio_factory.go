// +build deploy webaudio

package audio

import (
	// "encoding/binary"
	// "fmt"
	// "github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	"github.com/gopherjs/gopherjs/js"
)

// func init() {
// 	engine.NewIntegrant(&WebAudioFactory{})
// 	fmt.Println("> Audio Factory (Web): Ready")
// }

// type WebAudioFactory struct {
// 	engine.Integrant
// }

// func (o *WebAudioFactory) InitIntegrant() {}

// func (o *WebAudioFactory) AddIntegrant(integrant engine.IntegrantRoutine) {}

// func (o *WebAudioFactory) Destroy() {}

// type OpenALClip struct {
// 	*audio.Clip
// 	data []byte
// }

// func (o *OpenALClip) Data() []byte {
// 	return o.data
// }

// func (o *OpenALClip) Decode() {
// 	o.Clip.Set(
// 		binary.LittleEndian.Uint32(o.data[4:8]),
// 		o.data[22],
// 		uint32(o.data[24])|uint32(o.data[25])<<8|uint32(o.data[26])<<16|uint32(o.data[27])<<24,
// 		uint32(o.data[34]),
// 	)
// 	o.data = o.data[44:]
// }

type WebAudioSound struct {
	*audio.Sound
	buffer *js.Object
}

func NewOpenALSound(s *audio.Sound) *WebAudioSound {
	webAudioSound := &WebAudioSound{Sound: s}
	return webAudioSound
}

func (a *WebAudioSound) OverrideClip(fn func(audio.Clipable)) {
	a.SetClip = fn
	a.SetClip(a.BaseClip().(*audio.Clip))
}

func (a *WebAudioSound) Buffer() *js.Object {
	return a.buffer
}

type WebAudioSource struct {
	*audio.Source
	// source position
}

func NewWebAudioSource(s *audio.Source) *WebAudioSource {
	webAudioSource := &WebAudioSource{Source: s}
	return webAudioSource
}

func (a *WebAudioSource) OverridePlaySound(fn func(audio.Soundable)) {
	a.PlaySound = fn
	s := a.BasePlaySound()
	if s != nil {
		a.PlaySound(s.(*audio.Sound))
	}
}
