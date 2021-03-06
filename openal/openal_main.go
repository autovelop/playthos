// +build deploy openal

package openal

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"

	// for now we always pull in wav if openal is used until other audio file decoders exist
	// _ "github.com/autovelop/playthos/glfw"
	"bytes"

	"golang.org/x/mobile/exp/audio/al"

	"fmt"
	"io"
	"log"
)

func init() {
	audio.NewAudioSystem(&OpenAL{})
	fmt.Println("> OpenAL Added")
}

// OpenAL listener used for spacial audio using OpenAL
type OpenAL struct {
	audio.Audio
	platform engine.Platformer

	sources  []*OpenALSource
	sounds   []*OpenALSound
	listener *audio.Listener
	settings *engine.Settings
}

// InitSystem called when the system plugs into the engine
func (o *OpenAL) InitSystem() {
	al.OpenDevice()
}

// Destroy called when engine is gracefully shutting down
func (o *OpenAL) Destroy() {
	al.CloseDevice()
}

// AddIntegration helps the engine determine which integrants this system recognizes (Dependency Injection)
func (o *OpenAL) AddIntegrant(integrant engine.IntegrantRoutine) {
	switch integrant := integrant.(type) {
	// case *OpenALFactory:
	// 	o.factory = integrant
	// 	break
	case engine.Platformer:
		o.platform = integrant
		break
	}
}

// AddComponent unorphans a component by adding it to this system
func (o *OpenAL) AddComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *audio.Source:
		o.RegisterSource(component)
		break
	case *audio.Sound:
		o.RegisterSound(component)
		break
	case *audio.Listener:
		o.RegisterListener(component)
		break
	}
}

// ComponentTypes helps the engine determine which components this system recognizes (Dependency Injection)
func (o *OpenAL) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&audio.Source{}, &audio.Sound{}, &audio.Listener{}}
}

// func (o *OpenAL) Draw() {
// 	if o.Active() {
// 	}
// }

// DeleteEntity removes all entity's compoents from this system
func (o *OpenAL) DeleteEntity(entity *engine.Entity) {
	for i := 0; i < len(o.sounds); i++ {
		clip := o.sounds[i]
		if clip.Entity().ID() == entity.ID() {
			copy(o.sources[i:], o.sources[i+1:])
			o.sources[len(o.sources)-1] = nil
			o.sources = o.sources[:len(o.sources)-1]

			copy(o.sounds[i:], o.sounds[i+1:])
			o.sounds[len(o.sounds)-1] = nil
			o.sounds = o.sounds[:len(o.sounds)-1]
		}
	}
}

// RegisterSource tells open al about a given source
func (o *OpenAL) RegisterSource(s *audio.Source) {
	// play := s.BasePlay()
	openALSource := &OpenALSource{Source: s}
	source := &al.GenSources(1)[0]
	openALSource.OverridePlaySound(func(so audio.Soundable) {
		// fmt.Println(o.sounds[so.Index()].Buffer())
		// openALSource.source =
		// openALSource.PlaySound(p)
		source.QueueBuffers(*o.sounds[so.Index()].Buffer())
		al.PlaySources(*source)

		if s.Loop() {
			source.Seti(paramLooping, 1)
		}
		// log.Fatal(sound.audioFile.Duration())
		// } else {
		// 	go func(sound *Sound) {
		// 		time.Sleep(sound.audioFile.Duration())
		// 		sound.playing = false
		// 	}(sound)
		// }
	})
	openALSource.source = source
	o.sources = append(o.sources, openALSource)
}

// RegisterSource tells open al about a given sound
func (o *OpenAL) RegisterSound(c *audio.Sound) {
	clip := c.BaseClip()
	openALSound := &OpenALSound{Sound: c}
	if clip != nil {
		openALSound.OverrideClip(func(c audio.Clipable) {
			raw := o.platform.Asset(c.Path())
			if len(raw) > 0 {
				openALClip := &OpenALClip{c.(*audio.Clip), raw}
				openALClip.Decode()
				openALSound.clip = openALClip

				buffer := &al.GenBuffers(1)[0]

				r := bytes.NewReader(openALClip.data)

				s := make([]byte, openALClip.Length())
				// Uncomment this variable and try use it to determine the size of the audio clip
				// size := int64(0)
				for {
					n, err := r.Read(s)
					if n > 0 {
						// size += int64(n)
						buffer.BufferData(al.FormatStereo16, s[:n], int32(openALClip.SampleRate()))
					}
					if err == io.EOF {
						break
					}
					if err != nil {
						log.Fatal(err)
					}
				}

				openALSound.buffer = buffer
			}
		})
	}
	o.sounds = append(o.sounds, openALSound)
	openALSound.SetIndex(len(o.sounds) - 1)

	// if openALSound.PlayOnReady() {
	// 	openALSound.Play()
	// }
}

// RegisterSource tells open al about a given listener
func (o *OpenAL) RegisterListener(l *audio.Listener) {
	pos := l.Position()
	al.SetListenerPosition([3]float32{pos.X, pos.Y, pos.Z})
	o.listener = l
}

const (
	paramLooping = 0x1007
)
