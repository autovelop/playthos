// +build deploy openal

package openal

import (
	"encoding/binary"
	// "fmt"
	// "github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	"golang.org/x/mobile/exp/audio/al"
)

type OpenALClip struct {
	*audio.Clip
	data []byte
}

func (o *OpenALClip) Data() []byte {
	return o.data
}

func (o *OpenALClip) Decode() {
	o.Clip.Set(
		binary.LittleEndian.Uint32(o.data[4:8]),
		o.data[22],
		uint32(o.data[24])|uint32(o.data[25])<<8|uint32(o.data[26])<<16|uint32(o.data[27])<<24,
		uint32(o.data[34]),
	)
	o.data = o.data[44:]
}

type OpenALSound struct {
	*audio.Sound
	clip   *OpenALClip
	buffer *al.Buffer
}

func NewOpenALSound(s *audio.Sound) *OpenALSound {
	openALSound := &OpenALSound{Sound: s}
	return openALSound
}

func (o *OpenALSound) OverrideClip(fn func(audio.Clipable)) {
	o.SetClip = fn
	o.SetClip(o.BaseClip().(*audio.Clip))
}

func (o *OpenALSound) Buffer() *al.Buffer {
	return o.buffer
}

type OpenALSource struct {
	*audio.Source
	source *al.Source
}

func NewOpenALSource(s *audio.Source) *OpenALSource {
	openALSource := &OpenALSource{Source: s}
	return openALSource
}

func (o *OpenALSource) OverridePlaySound(fn func(audio.Soundable)) {
	o.PlaySound = fn
	s := o.BasePlaySound()
	if s != nil {
		o.PlaySound(s.(*audio.Sound))
	}
}

// func (o *OpenALSource) OverridePlay(fn func(audio.Playable)) {
// 	o.Play = fn
// o.Play(o.BasePlay().(*audio.Source))
// }

// func (o *OpenALSource) PlaySound(s *OpenALSound) {
// 	o.source.QueueBuffers(*s.buffer)
// 	al.PlaySources(*o.source)
// }
