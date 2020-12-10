package bus

import "fmt"

type Timer struct {
	active bool
	//decremented at a rate of 60HZ(60 times per second)
	value uint8
}

func (timer *Timer) tick() {
	if timer.value == 0 {
		return
	}
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

func (cpu *Cpu) config() {
	cpu.opcodes = make(map[string]Opcode)
	cpu.PC = 0x200
	cpu.configOpcodes()

}
func (cpu *Cpu) execute() {
	opcodePrefix := uint16(cpu.bus.ram.read(cpu.PC))
	opcodeSuffix := uint16(cpu.bus.ram.read(cpu.PC + 1))
	cpu.PC += 2
	opcode := (opcodePrefix << 8) | opcodeSuffix
	cpu.opcode = opcode
	operationKey := getOperationKey(cpu.opcode)

	if op, ok := cpu.opcodes[operationKey]; ok {
		op.operation()
	} else { //debugging

		operationKey = getOperationKey(cpu.opcode)
		fmt.Printf("Wrong Key:%X\n", operationKey)

	}
}
