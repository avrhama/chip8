package bus

import "fmt"

type Bus struct {
	cpu *Cpu
	ram *Ram
}

func (bus *Bus) ConfigBus() {
	bus.cpu = &Cpu{}
	bus.cpu.bus = bus
	bus.cpu.configCpu()
	bus.cpu.configOpcodes()

}
func (bus *Bus) TurnOn() {
	op := bus.cpu.opcodes["op0nnn"]
	fmt.Println(op)
	bus.cpu.opcodes["op0nnn"].operation()
	bus.cpu.dt.tick()
	//bus.cpu.opcodes["0nnn"].operation(bus.cpu)
}
