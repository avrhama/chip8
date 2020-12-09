package bus

import "fmt"

type Bus struct {
	cpu     *Cpu
	ram     *Ram
	display *Display
}

func (bus *Bus) ConfigBus() {
	bus.cpu = &Cpu{}
	bus.cpu.bus = bus
	bus.display = &Display{width: 200, height: 200}

	bus.cpu.configCpu()
	bus.cpu.configOpcodes()
	bus.display.config()
	bus.display.test()

}
func (bus *Bus) TurnOn() {
	op := bus.cpu.opcodes["op0nnn"]
	fmt.Println(op)
	bus.cpu.opcodes["op0nnn"].operation()
	bus.cpu.dt.tick()
	//bus.cpu.opcodes["0nnn"].operation(bus.cpu)
}
func (bus *Bus) TurnOff() {
	bus.display.turnOff(3)
}
