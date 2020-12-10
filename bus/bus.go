package bus

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Bus struct {
	cpu     *Cpu
	ram     *Ram
	display *Display
	joypad  *Joypad
	apu     *Apu
}

func (bus *Bus) configBus() {
	bus.cpu = &Cpu{}
	bus.cpu.bus = bus
	bus.display = &Display{width: 64, height: 32, windowWidth: 640, windowHeight: 320}
	bus.joypad = &Joypad{}
	bus.ram = &Ram{}
	bus.apu = &Apu{}

	bus.cpu.config()
	bus.display.config()
	bus.joypad.config()
	bus.apu.config()

}
func (bus *Bus) TurnOn(romPath string) {
	bus.configBus()
	bus.ram.loadRom(romPath)
	//the amount of opcodes that the cpu executes per cpu.dt tick(=60HZ).
	//this variable controls the speed of the game  increasing it increases the game speed.
	oPTT := 10
	oPTTCounter := 0
	sdl.Delay(5000)
	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch et := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if event.GetType() == sdl.KEYDOWN {
					if key, ok := bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)]; ok {
						key.pressed = true
						bus.joypad.keys[sdl.GetKeyName(et.Keysym.Sym)] = key
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

		oPTTCounter = 0
		for oPTTCounter < oPTT {
			bus.cpu.execute()
			oPTTCounter++
		}
		bus.cpu.dt.tick()
		bus.cpu.st.tick()
		bus.display.draw()
		if bus.cpu.st.value > 0 {
			bus.apu.play()
		}
		sdl.Delay(16)

	}
}
func (bus *Bus) TurnOff() {
	bus.display.turnOff(3)
	bus.apu.turnOff()
}
