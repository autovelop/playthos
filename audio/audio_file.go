package audio

import (
	"fmt"
	"io"
	"time"
)

type AudioFile struct {
	AudioFileRoutine
	channels      byte
	bitDepth      uint32
	sampleRate    uint32
	HeaderSize    int
	fileSize      uint32
	fileType      string
	audioFileType string
	duration      time.Duration
	data          []byte
	Reader        io.Reader
}

func (a *AudioFile) Set(fileType string, fileSize uint32, audioFileType string, channels byte, sampleRate uint32, bitDepth uint32) {
	a.fileType = fileType
	a.fileSize = fileSize
	a.audioFileType = audioFileType
	a.channels = channels
	a.sampleRate = sampleRate
	a.bitDepth = bitDepth

	a.duration, _ = time.ParseDuration(fmt.Sprintf("%vs", float32(a.fileSize)/float32(uint32(sampleRate)*uint32(channels)*a.bitDepth/8)))
}

func (a *AudioFile) SampleRate() uint32 {
	return a.sampleRate
}

func (a *AudioFile) BitDepth() uint32 {
	return a.bitDepth
}

func (a *AudioFile) FileSize() uint32 {
	return a.fileSize
}

func (a *AudioFile) Duration() time.Duration {
	return a.duration
}

func (a *AudioFile) SetReader(reader io.Reader) {
	a.Reader = reader
}

type AudioFileRoutine interface {
	SampleRate() uint32
	BitDepth() uint32
	FileSize() uint32
	Duration() time.Duration
	GetReader() io.Reader

	Load(string, string)
}
