package wav

import (
	"bytes"
	"encoding/binary"
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	// "go/build"
	"io"
	// "io/ioutil"
	"log"
	// "os"
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

// func (w *WAVAudioFile) Load(assetFolder string, filename string) {
// 	dir, err := build.Import(assetFolder, "", build.FindOnly)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	if err != nil {
// 		log.Fatal("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
// 		return
// 	}
// 	err = os.Chdir(dir.Dir)
// 	if err != nil {
// 		log.Fatal("os.Chdir:", err)
// 		return
// 	}

// 	file, err := os.Open(filename)
// 	// file, err := os.Open("COBRK3618.wav")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	buf := make([]byte, w.HeaderSize)
// 	if n, _ := io.ReadFull(file, buf); n != w.HeaderSize {
// 		log.Fatal("No header present or read error.")
// 		return
// 	}

// 	// 1 - 4 | "RIFF" | Marks the file as a riff file. Characters are each 1 byte long.
// 	// 5 - 8 | File size (integer) | Size of the overall file - 8 bytes, in bytes (32-bit integer). Typically, you'd fill this in after creation.
// 	// 9 -12 | "WAVE" | File Type Header. For our purposes, it always equals "WAVE".
// 	// 13-16 | "fmt " | Format chunk marker. Includes trailing null
// 	// 17-20 | 16 | Length of format data as listed above
// 	// 21-22 | 1 | Type of format (1 is PCM) - 2 byte integer
// 	// 23-24 | 2 | Number of Channels - 2 byte integer
// 	// 25-28 | 44100 | Sample Rate - 32 byte integer. Common values are 44100 (CD), 48000 (DAT). Sample Rate = Number of Samples per second, or Hertz.
// 	// 29-32 | 176400 | (Sample Rate * BitsPerSample * Channels) / 8.
// 	// 33-34 | 4 | (BitsPerSample * Channels) / 8.1 - 8 bit mono2 - 8 bit stereo/16 bit mono4 - 16 bit stereo
// 	// 35-36 | 16 | Bits per sample
// 	// 37-40 | "data" | "data" chunk header. Marks the beginning of the data section.
// 	// 41-44 | File size | (data)
// 	log.Printf("buf[0:4]: %v\n", buf[0:4])
// 	log.Printf("buf[4:8]: %v\n", buf[4:8])
// 	log.Printf("buf[8:12]: %v\n", buf[8:12])
// 	log.Printf("buf[12:16]: %v\n", buf[12:16])
// 	log.Printf("buf[16:20]: %v\n", buf[16:20])
// 	log.Printf("buf[20:22]: %v\n", buf[20:22])
// 	log.Printf("buf[22:24]: %v\n", buf[22:24])
// 	log.Printf("buf[24:28]: %v\n", buf[24:28])
// 	log.Printf("buf[28:32]: %v\n", buf[28:32])
// 	log.Printf("buf[32:34]: %v\n", buf[32:34])
// 	log.Printf("buf[34:36]: %v\n", buf[34:36])
// 	log.Printf("buf[36:40]: %v\n", buf[36:40])
// 	log.Printf("buf[40:44]: %v\n", buf[40:44])
// 	log.Println("-")

// 	if buf[21] == 1 {
// 		w.IsPCM = true
// 	}
// 	w.Set(string(buf[0:4]), binary.LittleEndian.Uint32(buf[4:8]), string(buf[8:12]), buf[22], uint32(buf[24])|uint32(buf[25])<<8|uint32(buf[26])<<16|uint32(buf[27])<<24, uint32(buf[34]))

// 	buf, err = ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.SetReader(bytes.NewReader(buf))
// }

func (w *WAVAudioFile) Load(assetFolder string, filename string) {
	// don't load the whole file
	buf, err := engine.LoadAsset(assetFolder, filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("buf[0:4]: %v\n", buf[0:4])
	log.Printf("buf[4:8]: %v\n", buf[4:8])
	log.Printf("buf[8:12]: %v\n", buf[8:12])
	log.Printf("buf[12:16]: %v\n", buf[12:16])
	log.Printf("buf[16:20]: %v\n", buf[16:20])
	log.Printf("buf[20:22]: %v\n", buf[20:22])
	log.Printf("buf[22:24]: %v\n", buf[22:24])
	log.Printf("buf[24:28]: %v\n", buf[24:28])
	log.Printf("buf[28:32]: %v\n", buf[28:32])
	log.Printf("buf[32:34]: %v\n", buf[32:34])
	log.Printf("buf[34:36]: %v\n", buf[34:36])
	log.Printf("buf[36:40]: %v\n", buf[36:40])
	log.Printf("buf[40:44]: %v\n", buf[40:44])
	log.Println("-")

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
