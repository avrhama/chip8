package main

import (
	"fmt"
	//works
	//"chip8/cpu"
	"./bus"
)

func main() {
	fmt.Println("Chip8")
	bus_ := bus.Bus{}

	//bus_.cpu.R.High = 9
	bus_.ConfigBus()
	bus_.TurnOn()
	//fmt.Println(bus.cpu.R.High)
	//r := cpu.Register{}
	//r.High = 8
	//fmt.Println(r.High)
	//fmt.Println(cpu.GetOperation(0xA015))

}
