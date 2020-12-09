package main

import (
	"fmt"

	//works
	//"chip8/cpu"

	"./bus"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("Chip8")
	bus_ := bus.Bus{}

	//bus_.cpu.R.High = 9
	bus_.ConfigBus()
	//bus_.TurnOn()

	sdl.Delay(2000)
	bus_.TurnOff()

}
