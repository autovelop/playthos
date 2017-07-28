package wav

import (
	"bytes"
	"encoding/binary"
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	"io"
	"log"
)

type WAVAudioFile struct {
	audio.AudioFile
	IsPCM bool
}

func NewWAVFile() *WAVAudioFile {
	wavFile := &WAVAudioFile{}
	wavFile.SetHeaderSize(44)
	return wavFile
}

func (w *WAVAudioFile) Load(assetFolder string, filename string) {
	// don't load the whole file
	buf, err := engine.LoadAsset(assetFolder, filename)
	if err != nil {
		log.Fatal(err)
	}

	if buf[21] == 1 {
		w.IsPCM = true
	}
	w.Set(string(buf[0:4]), binary.LittleEndian.Uint32(buf[4:8]), string(buf[8:12]), buf[22], uint32(buf[24])|uint32(buf[25])<<8|uint32(buf[26])<<16|uint32(buf[27])<<24, uint32(buf[34]))
	buf = buf[w.HeaderSize():]
	w.SetReader(bytes.NewReader(buf))
}

func (w *WAVAudioFile) GetReader() io.Reader {
	return w.Reader
}
