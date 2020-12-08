package bus

/*type Register struct {
	High byte
	Low  byte
}
*/
type Timer struct {
	active bool
	//decremented at a rate of 60HZ(60 times per second)
	value uint8
}

func (timer *Timer) tick() {
	timer.value--
	if timer.value == 0 {
		timer.active = false
	}
}

type Cpu struct {
	registers [16]uint8
	bus       *Bus
	PC        uint16
	I         uint16
	SP        uint8
	opcodes   map[string]Opcode
	stack     [16]uint16
	//current opcode
	opcode uint16
	dt     Timer
	st     Timer
}

func (cpu *Cpu) configCpu() {
	cpu.opcodes = make(map[string]Opcode)

}
