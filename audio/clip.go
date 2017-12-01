// +build deploy audio

package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
	"time"
)

type Clipable interface {
	// Set(Clipable, bool, bool)
	Path() string
	Decode()
}

type Clip struct {
	path       string
	channels   byte
	sampleRate uint32
	bitDepth   uint32
	length     uint32
	duration   time.Duration
}

func NewClip() *Clip {
	return &Clip{}
}

// sucks that this has to exist
func (c *Clip) Decode() {}

func (c *Clip) Set(l uint32, ch byte, s uint32, b uint32) {
	c.length = l
	c.channels = ch
	c.sampleRate = s
	c.bitDepth = b
	c.duration, _ = time.ParseDuration(fmt.Sprintf("%vs", float32(c.length)/float32(uint32(c.sampleRate)*uint32(c.channels)*c.bitDepth/8)))
}

func (c *Clip) LoadClip(p string) {
	c.path = p
	engine.LoadAsset(p)
}

func (c *Clip) Path() string {
	return c.path
}

func (c *Clip) Length() uint32 {
	return c.length
}

func (c *Clip) Duration() time.Duration {
	return c.duration
}

func (c *Clip) SampleRate() uint32 {
	return c.sampleRate
}

// func (s *Sound) Play() {
// 	s.audioSystem.PlaySound(s)
// }

// func (s *Sound) Stop() {
// 	s.audioSystem.StopSound(s)
// }

// func (s *Sound) Get() AudioFileRoutine {
// 	return s.audioFile
// }
