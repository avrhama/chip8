package bus

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Apu struct {
	beepBytes *[]byte
	audioID   sdl.AudioDeviceID
	audioSpec *sdl.AudioSpec
}

func (apu *Apu) config() {
	var audioSpec sdl.AudioSpec
	beepBytes, _ := sdl.LoadWAV("./bus/beep.wav")
	audioID, err := sdl.OpenAudioDevice("", false, &audioSpec, nil, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	apu.audioID = audioID
	apu.audioSpec = &audioSpec
	apu.beepBytes = &beepBytes
}

func (apu *Apu) play() {
	sdl.QueueAudio(apu.audioID, *apu.beepBytes)
	sdl.PauseAudioDevice(apu.audioID, false)
}
func (apu *Apu) turnOff() {
	sdl.FreeWAV(*apu.beepBytes)
	sdl.PauseAudioDevice(apu.audioID, true)
	sdl.CloseAudioDevice(apu.audioID)
}
