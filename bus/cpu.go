package bus

type Register struct {
	High byte
	Low  byte
}
type Cpu struct {
	R       Register
	bus     *Bus
	PC      uint16
	SP      uint8
	opcodes map[string]Opcode
	stack   [16]uint16
}

func (cpu *Cpu) configCpu() {
	cpu.opcodes = make(map[string]Opcode)

}
