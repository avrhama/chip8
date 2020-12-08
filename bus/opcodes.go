package bus

import "fmt"

//I decide to spearte opcodes into 2 types.
//those which have prefix with one(single) operation relates to its(e.g 1nnn only one operation starts with 1 at this example is JP addr)
//and those which have prefix with multiple operations relate to its(e.g 0nnn, 00e0 and 00ee)
func GetOperation(opcode uint16) uint16 {
	prefix := opcode >> 12
	//if prefix!=0&&prefix!=8&&prefix!=0xe&&prefix!=0xf{
	//	return getSingle(opcode);
	//}
	return prefix
}

type Opcode struct {
	name string
	//operation func(cpu *Cpu)
	operation func()
}

func (cpu *Cpu) configOpcodes() {

	cpu.opcodes["op0nnn"] = Opcode{name: "op0nnn", operation: cpu.op0nnn}
	cpu.opcodes["op00E0"] = Opcode{name: "op00E0", operation: cpu.op00E0}
	cpu.opcodes["op00EE"] = Opcode{name: "op00EE", operation: cpu.op00EE}

}

/*
nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
n or nibble - A 4-bit value, the lowest 4 bits of the instruction
x - A 4-bit value, the lower 4 bits of the high byte of the instruction
y - A 4-bit value, the upper 4 bits of the low byte of the instruction
kk or byte - An 8-bit value, the lowest 8 bits of the instruction
*/
func getnnn(v uint16) uint16 {
	return v & 0xfff
}
func getn(v uint16) uint8 {
	return uint8(v & 0xf)
}
func getx(v uint16) uint8 {
	return uint8((v >> 8) & 0xf)
}
func gety(v uint16) uint8 {
	return uint8((v >> 12) & 0xf)
}
func getkk(v uint16) uint8 {
	return uint8(v & 0xff)
}

func (cpu *Cpu) operate() {
	opcodePrefix := uint16(cpu.bus.ram.Read(cpu.PC))
	cpu.PC += 1
	opcodeSuffix := uint16(cpu.bus.ram.Read(cpu.PC))
	opcode := (opcodePrefix << 8) | opcodeSuffix
	_ = opcode
}
func (cpu *Cpu) op0nnn() {
	fmt.Println("wellcome op0nnn")
}

/*
00E0 - CLS
Clear the display.
*/
func (cpu *Cpu) op00E0() {

}

/*
00EE - RET
Return from a subroutine.
The interpreter sets the program counter to the address at the top of the stack, then subtracts 1 from the stack pointer.
*/
func (cpu *Cpu) op00EE() {
	cpu.PC = cpu.stack[cpu.SP]
	cpu.SP--
}
