package audio

import (
	"io"
)

type AudioFile struct {
	AudioFileRoutine
	Channels      byte
	BitDepth      uint32
	SampleRate    uint32
	HeaderSize    int
	FileSize      uint32
	FileType      string
	AudioFileType string
	data          []byte
	Reader        io.Reader
}

func (a *AudioFile) GetSampleRate() uint32 {
	return a.SampleRate
}

func (a *AudioFile) GetBitDepth() uint32 {
	return a.BitDepth
}

func (a *AudioFile) GetFileSize() uint32 {
	return a.FileSize
}

type AudioFileRoutine interface {
	// base
	GetSampleRate() uint32
	GetBitDepth() uint32
	GetFileSize() uint32
	GetReader() io.Reader

	Load(string, string)
}
