// +build deploy openal

package openal

import (
	"encoding/binary"
	// "fmt"
	// "github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	"golang.org/x/mobile/exp/audio/al"
)

// OpenALClip defines a clip component specifically with OpenAL
type OpenALClip struct {
	*audio.Clip
	data []byte
}

// Data returns binary data of openal clip
func (o *OpenALClip) Data() []byte {
	return o.data
}

// Decode decodes clip for openal playback
func (o *OpenALClip) Decode() {
	o.Clip.Set(
		binary.LittleEndian.Uint32(o.data[4:8]),
		o.data[22],
		uint32(o.data[24])|uint32(o.data[25])<<8|uint32(o.data[26])<<16|uint32(o.data[27])<<24,
		uint32(o.data[34]),
	)
	o.data = o.data[44:]
}

// OpenALSound defines sound component specifically with OpenAL
type OpenALSound struct {
	*audio.Sound
	clip   *OpenALClip
	buffer *al.Buffer
}

// NewOpenALSound creates and sets a new orphan openal sound
func NewOpenALSound(s *audio.Sound) *OpenALSound {
	openALSound := &OpenALSound{Sound: s}
	return openALSound
}

// OverrideClip overrides base clip with openal clip
func (o *OpenALSound) OverrideClip(fn func(audio.Clipable)) {
	o.SetClip = fn
	o.SetClip(o.BaseClip().(*audio.Clip))
}

// Buffer returns pointer to openal buffer
func (o *OpenALSound) Buffer() *al.Buffer {
	return o.buffer
}

// OpenALSource defines source component specifically with OpenAL
type OpenALSource struct {
	*audio.Source
	source *al.Source
}

// NewOpenALSource creates and sets a new orphan openal source
func NewOpenALSource(s *audio.Source) *OpenALSource {
	openALSource := &OpenALSource{Source: s}
	return openALSource
}

// OverridePlaySound overrides base sound with openal sound
func (o *OpenALSource) OverridePlaySound(fn func(audio.Soundable)) {
	o.PlaySound = fn
	s := o.BasePlaySound()
	if s != nil {
		o.PlaySound(s.(*audio.Sound))
	}
}
