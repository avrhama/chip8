package bus

import (
	"fmt"
	"math/rand"
	"time"
)

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
	cpu.opcodes["op1nnn"] = Opcode{name: "op1nnn", operation: cpu.op1nnn}
	cpu.opcodes["op2nnn"] = Opcode{name: "op2nnn", operation: cpu.op2nnn}
	//op3xkk
	//op4xkk
	//op5xy0
	//op6xkk
	//op7xkk
	//op8xy0
	//op8xy1
	//op8xy2
	//op8xy3
	//op8xy4
	//op8xy5
	//op8xy6
	//op8xy7
	//op8xyE
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
	opcodePrefix := uint16(cpu.bus.ram.read(cpu.PC))
	cpu.PC += 1
	opcodeSuffix := uint16(cpu.bus.ram.read(cpu.PC))
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

/*
1nnn - JP addr
Jump to location nnn.
The interpreter sets the program counter to nnn.
*/
func (cpu *Cpu) op1nnn() {
	nnn := getnnn(cpu.opcode)
	cpu.PC = nnn
}

/*
2nnn - CALL addr
Call subroutine at nnn.
The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
*/
func (cpu *Cpu) op2nnn() {
	nnn := getnnn(cpu.opcode)
	cpu.SP++
	cpu.stack[cpu.SP] = cpu.PC
	cpu.PC = nnn
}

/*
3xkk - SE Vx, byte
Skip next instruction if Vx = kk.
The interpreter compares register Vx to kk, and if they are equal, increments the program counter by 2.
*/
func (cpu *Cpu) op3xkk() {
	x := getx(cpu.opcode)
	kk := getkk(cpu.opcode)
	if cpu.registers[x] == kk {
		cpu.PC += 2
	}
}

/*
4xkk - SNE Vx, byte
Skip next instruction if Vx != kk.
The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
*/
func (cpu *Cpu) op4xkk() {
	x := getx(cpu.opcode)
	kk := getkk(cpu.opcode)
	if cpu.registers[x] != kk {
		cpu.PC += 2
	}
}

/*
5xy0 - SE Vx, Vy
Skip next instruction if Vx = Vy.
The interpreter compares register Vx to register Vy, and if they are equal, increments the program counter by 2.
*/
func (cpu *Cpu) op5xy0() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	if cpu.registers[x] == cpu.registers[y] {
		cpu.PC += 2
	}
}

/*
6xkk - LD Vx, byte
Set Vx = kk.
The interpreter puts the value kk into register Vx.
*/
func (cpu *Cpu) op6xkk() {
	x := getx(cpu.opcode)
	kk := getkk(cpu.opcode)
	cpu.registers[x] = kk
}

/*
7xkk - ADD Vx, byte
Set Vx = Vx + kk.
Adds the value kk to the value of register Vx, then stores the result in Vx.
*/
func (cpu *Cpu) op7xkk() {
	x := getx(cpu.opcode)
	kk := getkk(cpu.opcode)
	cpu.registers[x] += kk
}

/*
8xy0 - LD Vx, Vy
Set Vx = Vy.
Stores the value of register Vy in register Vx.
*/
func (cpu *Cpu) op8xy0() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	cpu.registers[x] = cpu.registers[y]
}

/*
8xy1 - OR Vx, Vy
Set Vx = Vx OR Vy.
Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A bitwise OR compares the corrseponding bits from two values, and if either bit is 1, then the same bit in the result is also 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy1() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] | cpu.registers[y]
}

/*
8xy2 - AND Vx, Vy
Set Vx = Vx AND Vy.
Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx. A bitwise AND compares the corrseponding bits from two values, and if both bits are 1, then the same bit in the result is also 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy2() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] & cpu.registers[y]
}

/*
8xy3 - XOR Vx, Vy
Set Vx = Vx XOR Vy.
Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx. An exclusive OR compares the corrseponding bits from two values, and if the bits are not both the same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy3() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] ^ cpu.registers[y]
}

/*
8xy4 - ADD Vx, Vy
Set Vx = Vx + Vy, set VF = carry.
The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
*/
func (cpu *Cpu) op8xy4() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	res := uint16(x) + uint16(y)
	cpu.registers[x] = uint8(res)

	if res > 255 {
		cpu.registers[0xf] = 1
	} else {
		cpu.registers[0xf] = 0
	}
}

/*
8xy5 - SUB Vx, Vy
Set Vx = Vx - Vy, set VF = NOT borrow.
If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results stored in Vx.
*/
func (cpu *Cpu) op8xy5() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	if cpu.registers[x] > cpu.registers[y] {
		cpu.registers[0xf] = 1
	} else {
		cpu.registers[0xf] = 0
	}
	cpu.registers[x] = cpu.registers[x] - cpu.registers[y]
}

/*
8xy6 - SHR Vx {, Vy}
Set Vx = Vx SHR 1.
If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is divided by 2.
*/
func (cpu *Cpu) op8xy6() {
	x := getx(cpu.opcode)
	if cpu.registers[x]&1 == 1 {
		cpu.registers[0xf] = 1
	} else {
		cpu.registers[0xf] = 0
	}
	cpu.registers[x] = cpu.registers[x] >> 1
}

/*
8xy7 - SUBN Vx, Vy
Set Vx = Vy - Vx, set VF = NOT borrow.
If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
*/
func (cpu *Cpu) op8xy7() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	if cpu.registers[y] > cpu.registers[x] {
		cpu.registers[0xf] = 1
	} else {
		cpu.registers[0xf] = 0
	}
	cpu.registers[x] = cpu.registers[y] - cpu.registers[x]
}

/*
8xyE - SHL Vx {, Vy}
Set Vx = Vx SHL 1.
If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is multiplied by 2.
*/
func (cpu *Cpu) op8xyE() {
	x := getx(cpu.opcode)
	if cpu.registers[x]&0x80 == 0x80 {
		cpu.registers[0xf] = 1
	} else {
		cpu.registers[0xf] = 0
	}
	cpu.registers[x] = cpu.registers[x] << 1
}

/*
9xy0 - SNE Vx, Vy
Skip next instruction if Vx != Vy.
The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
*/
func (cpu *Cpu) op9xy0() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	if cpu.registers[x] != cpu.registers[y] {
		cpu.PC += 2
	}
}

/*
Annn - LD I, addr
Set I = nnn.
The value of register I is set to nnn.
*/
func (cpu *Cpu) opAnnn() {
	nnn := getnnn(cpu.opcode)
	cpu.I = nnn
}

/*
Bnnn - JP V0, addr
Jump to location nnn + V0.
The program counter is set to nnn plus the value of V0.
*/
func (cpu *Cpu) opBnnn() {
	nnn := getnnn(cpu.opcode)
	cpu.PC = uint16(cpu.registers[0]) + nnn
}

/*
Cxkk - RND Vx, byte
Set Vx = random byte AND kk.
The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
*/
func (cpu *Cpu) opCxkk() {
	x := getx(cpu.opcode)
	kk := getkk(cpu.opcode)
	//TODO:place it outside the function
	rand.Seed(time.Now().UnixNano())
	t := uint8(rand.Intn(256))
	cpu.registers[x] = kk & t
}

/*
Dxyn - DRW Vx, Vy, nibble
Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
The interpreter reads n bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen at coordinates (Vx, Vy). Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of the display, it wraps around to the opposite side of the screen. See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information on the Chip-8 screen and sprites.
*/
func (cpu *Cpu) opDxyn() {
	x := getx(cpu.opcode)
	y := getx(cpu.opcode)
	_ = x
	_ = y
	//TODO:finish thi function
}

/*
Ex9E - SKP Vx
Skip next instruction if key with the value of Vx is pressed.
Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
*/
func (cpu *Cpu) opEx9E() {
	x := getx(cpu.opcode)
	_ = x
	//TODO:finish thi function
}

/*
ExA1 - SKNP Vx
Skip next instruction if key with the value of Vx is not pressed.
Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
*/
func (cpu *Cpu) opSKNP() {
	x := getx(cpu.opcode)
	_ = x
	//TODO:finish thi function
}

/*
Fx07 - LD Vx, DT
Set Vx = delay timer value.
The value of DT is placed into Vx.
*/
func (cpu *Cpu) opFx07() {
	x := getx(cpu.opcode)
	cpu.registers[x] = cpu.dt.value
}

/*
Fx0A - LD Vx, K
Wait for a key press, store the value of the key in Vx.
All execution stops until a key is pressed, then the value of that key is stored in Vx.
*/
func (cpu *Cpu) opFx0A() {
	x := getx(cpu.opcode)
	_ = x
	//TODO:finish thi function
}

/*
Fx15 - LD DT, Vx
Set delay timer = Vx.
DT is set equal to the value of Vx.
*/
func (cpu *Cpu) opFx15() {
	x := getx(cpu.opcode)
	cpu.dt.value = cpu.registers[x]
}

/*
Fx18 - LD ST, Vx
Set sound timer = Vx.
ST is set equal to the value of Vx.
*/
func (cpu *Cpu) opFx18() {
	x := getx(cpu.opcode)
	cpu.st.value = cpu.registers[x]
}

/*
Fx1E - ADD I, Vx
Set I = I + Vx.
The values of I and Vx are added, and the results are stored in I.
*/
func (cpu *Cpu) opFx1E() {
	x := getx(cpu.opcode)
	cpu.I = cpu.I + uint16(cpu.registers[x])
}

/*
Fx29 - LD F, Vx
Set I = location of sprite for digit Vx.
The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
*/
func (cpu *Cpu) opFx29() {
	x := getx(cpu.opcode)
	_ = x
	//TODO:finish this function
}

/*
Fx33 - LD B, Vx
Store BCD representation of Vx in memory locations I, I+1, and I+2.
The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
*/
func (cpu *Cpu) opFx33() {
	x := getx(cpu.opcode)
	v := cpu.registers[x]
	o := v % 10
	v = v / 10
	t := v % 10
	v = v / 10
	h := v % 10

	cpu.bus.ram.write(cpu.I, h)
	cpu.bus.ram.write(cpu.I+1, t)
	cpu.bus.ram.write(cpu.I+2, o)

}

/*
Fx55 - LD [I], Vx
Store registers V0 through Vx in memory starting at location I.
The interpreter copies the values of registers V0 through Vx into memory, starting at the address in I.
*/
func (cpu *Cpu) opFx55() {
	x := uint16(getx(cpu.opcode))
	for i := uint16(0); i <= x; i++ {
		cpu.bus.ram.write(cpu.I+i, cpu.registers[i])
	}

}

/*
Fx65 - LD Vx, [I]
Read registers V0 through Vx from memory starting at location I.
The interpreter reads values from memory starting at location I into registers V0 through Vx.
*/
func (cpu *Cpu) opFx65() {
	x := uint16(getx(cpu.opcode))
	for i := uint16(0); i <= x; i++ {
		cpu.registers[i] = cpu.bus.ram.read(cpu.I + i)
	}

}
