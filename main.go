package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	//works
	//"chip8/cpu"
	//"github.com/veandco/go-sdl2/sdl"
	"./bus"
)

func main() {
	fmt.Println("Chip8")
	bus_ := bus.Bus{}
	//bus_.ConfigBus()
	//_ = bus
	bus_.TurnOn()
	sdl.Delay(2000)
	bus_.TurnOff()
	/*dat, err := ioutil.ReadFile("C:\\Users\\Epsilon\\Documents\\roms\\glitchGhost.ch8")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error")
		os.Exit(1)
	}
	println(len(dat))*/

}
