package bus

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Bus struct {
	cpu     *Cpu
	ram     *Ram
	display *Display
	joypad  *Joypad
}

func (bus *Bus) configBus() {
	bus.cpu = &Cpu{}
	bus.cpu.bus = bus
	bus.display = &Display{width: 64, height: 32, windowWidth: 640, windowHeight: 320}
	bus.joypad = &Joypad{}
	bus.ram = &Ram{}

	bus.cpu.configCpu()
	bus.cpu.configOpcodes()
	bus.display.config()
	bus.joypad.config()

}
func (bus *Bus) TurnOn() {
	bus.configBus()

	bus.ram.loadRom("C:\\Users\\Epsilon\\Documents\\roms\\chip8\\Space Invaders [David Winter].ch8")

	//bus.ram.loadRom("C:\\Users\\Epsilon\\Documents\\roms\\chip8\\glitchGhost.ch8")
	//bus.ram.loadRom("C:\\Users\\Epsilon\\Documents\\roms\\chip8\\pong.rom")
	//the amount of opcodes that the cpu executes per cpu.dt tick(=60HZ).
	oPTT := 60
	oPTTCounter := 0
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch et := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if event.GetType() == sdl.KEYDOWN {
					//fmt.Println(sdl.GetKeyName(et.Keysym.Sym))
					if key, ok := bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)]; ok {
						key.pressed = true
						bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)] = key
						fmt.Println(key.name)
					}

				}

				if event.GetType() == sdl.KEYUP {
					if key, ok := bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)]; ok {
						key.pressed = false
						bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)] = key
					}
				}

				break
			}
		}
		bus.cpu.dt.tick()
		bus.cpu.st.tick()
		oPTTCounter = 0
		for oPTTCounter < oPTT {
			bus.cpu.execute()
			oPTTCounter++
		}

		bus.display.draw()
		sdl.Delay(16)

	}
}
func (bus *Bus) TurnOff() {
	bus.display.turnOff(3)
}
