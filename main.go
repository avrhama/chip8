package main

import (
	"fmt"
	"os"

	"./bus"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("rom path is missing!(main.exe romPath)")
	}
	fmt.Println("Chip8")
	bus_ := bus.Bus{}
	bus_.TurnOn(os.Args[1])
	bus_.TurnOff()
	fmt.Println("goodbye!...")

}
