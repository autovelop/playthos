// +build deploy audio

package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
	"time"
)

// Clipable interface allows for decoding and locating audio clips
type Clipable interface {
	Path() string
	Decode()
}

// Clip defines the path and binary data of an audio clip
type Clip struct {
	path       string
	channels   byte
	sampleRate uint32
	bitDepth   uint32
	length     uint32
	duration   time.Duration
}

// NewClip creates and sets a new orphan clip
func NewClip() *Clip {
	return &Clip{}
}

// Decode called when necessary for the clip to be decoded
func (c *Clip) Decode() {}

// Set used to define all the require properties of a Clip
func (c *Clip) Set(l uint32, ch byte, s uint32, b uint32) {
	c.length = l
	c.channels = ch
	c.sampleRate = s
	c.bitDepth = b
	c.duration, _ = time.ParseDuration(fmt.Sprintf("%vs", float32(c.length)/float32(uint32(c.sampleRate)*uint32(c.channels)*c.bitDepth/8)))
}

// LoadClip sets path of clip file and loads it into the engine
func (c *Clip) LoadClip(p string) {
	c.path = p
	engine.LoadAsset(p)
}

// Path returns the clip path
func (c *Clip) Path() string {
	return c.path
}

// Length returns the clip length
func (c *Clip) Length() uint32 {
	return c.length
}

// Duration returns the clip duration
func (c *Clip) Duration() time.Duration {
	return c.duration
}

// SampleRate returns the clip SampleRate
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
