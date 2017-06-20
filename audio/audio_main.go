// +build desktop,audio

package audio

import (
	"github.com/autovelop/playthos"
	"golang.org/x/mobile/exp/audio/al"
	"io"
	"log"
)

func init() {
	audio := &Audio{}
	engine.NewUnloadedObserverable(audio)
}

type Audio struct {
	// source al.Source
	buffer al.Buffer
}

func (a *Audio) Prepare(settings *engine.Settings) {
	log.Println("Prep")
	al.OpenDevice()
	al.SetListenerPosition([3]float32{0, 0, 0})
}

func (a *Audio) PlaySound(sound *Sound) {
	source := al.GenSources(1)[0]
	if code := al.Error(); code != 0 {
		log.Fatalf("audio: cannot generate an audio source [err=%x]", code)
	}

	buffer := al.GenBuffers(1)[0]
	audioFile := sound.Get()
	buf := make([]byte, audioFile.GetFileSize())
	size := int64(0)
	for {
		n, err := audioFile.GetReader().Read(buf)
		if n > 0 {
			size += int64(n)
			buffer.BufferData(al.FormatStereo16, buf[:n], int32(audioFile.GetSampleRate()))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	source.QueueBuffers(buffer)
	al.PlaySources(source)

	// its fine if its already playing
	if sound.IsLoop() {
		source.Seti(paramLooping, 1)
	} else {
		// wait till end of sound then delete buffer
		// al.DeleteSources(source)
		// al.DeleteBuffers(1, buffer)
	}
}

func (a *Audio) UnRegisterEntity(entity *engine.Entity) {
}

func (a *Audio) LoadComponent(component engine.ComponentRoutine) {
	// switch component := component.(type) {
	// case *glfw.GLFW:
	// 	a.window = component.GetWindow()
	// 	break
	// }
}

/*
The golang experimental openal bindings for have looping as a const yet. So I just guessed it based on the below C code.
https://github.com/kcat/openal-soft/blob/0a361fa9e27b9d9533dffe34663efc3669205b86/Alc/ALc.c
*/
const (
	paramLooping = 0x1007
)
