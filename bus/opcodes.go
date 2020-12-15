package bus

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//I decide to spearte opcodes into 2 types.
//those which have prefix with one(single) operation relates to its(e.g 1nnn only one operation starts with 1 at this example is JP addr)
//and those which have prefix with multiple operations relate to its(e.g 0nnn, 00e0 and 00ee)
func getOperationKey(opcode uint16) string {
	prefix := opcode >> 12

	if prefix != 0 && prefix != 8 && prefix != 0xe && prefix != 0xf {
		return fmt.Sprintf("%X", prefix)
	}
	switch prefix {
	case 0:
		s := opcode & 0xff
		if s == 0xe0 {
			return "0.1"
		} else if s == 0xee {
			return "0.2"
		}
		return "0.0"
	case 8:
		s := opcode & 0x0f
		if s == 1 {
			return "8.1"
		} else if s == 2 {
			return "8.2"
		} else if s == 3 {
			return "8.3"
		} else if s == 4 {
			return "8.4"
		} else if s == 5 {
			return "8.5"
		} else if s == 6 {
			return "8.6"
		} else if s == 7 {
			return "8.7"
		} else if s == 0xe {
			return "8.8"
		}
		return "8.0"
	case 0xe:
		s := opcode & 0xff
		if s == 0xa1 {
			return "E.1"
		}
		return "E.0"
	case 0xf:
		s := opcode & 0xff
		if s == 0x0a {
			return "F.1"
		} else if s == 0x15 {
			return "F.2"
		} else if s == 0x18 {
			return "F.3"
		} else if s == 0x1e {
			return "F.4"
		} else if s == 0x29 {
			return "F.5"
		} else if s == 0x33 {
			return "F.6"
		} else if s == 0x55 {
			return "F.7"
		} else if s == 0x65 {
			return "F.8"
		}
		return "F.0"

	}
	return ""
}

type Opcode struct {
	//used for debugging
	name      string
	operation func()
}

func (cpu *Cpu) configOpcodes() {

	cpu.opcodes["0.0"] = Opcode{name: "op0nnn", operation: cpu.op0nnn}
	cpu.opcodes["0.1"] = Opcode{name: "op00E0", operation: cpu.op00E0}
	cpu.opcodes["0.2"] = Opcode{name: "op00EE", operation: cpu.op00EE}

	cpu.opcodes["1"] = Opcode{name: "op1nnn", operation: cpu.op1nnn}
	cpu.opcodes["2"] = Opcode{name: "op2nnn", operation: cpu.op2nnn}
	cpu.opcodes["3"] = Opcode{name: "op3xkk", operation: cpu.op3xkk}
	cpu.opcodes["4"] = Opcode{name: "op4xkk", operation: cpu.op4xkk}
	cpu.opcodes["5"] = Opcode{name: "op5xy0", operation: cpu.op5xy0}
	cpu.opcodes["6"] = Opcode{name: "op6xkk", operation: cpu.op6xkk}
	cpu.opcodes["7"] = Opcode{name: "op7xkk", operation: cpu.op7xkk}

	cpu.opcodes["8.0"] = Opcode{name: "op8xy0", operation: cpu.op8xy0}
	cpu.opcodes["8.1"] = Opcode{name: "op8xy1", operation: cpu.op8xy1}
	cpu.opcodes["8.2"] = Opcode{name: "op8xy2", operation: cpu.op8xy2}
	cpu.opcodes["8.3"] = Opcode{name: "op8xy3", operation: cpu.op8xy3}
	cpu.opcodes["8.4"] = Opcode{name: "op8xy4", operation: cpu.op8xy4}
	cpu.opcodes["8.5"] = Opcode{name: "op8xy5", operation: cpu.op8xy5}
	cpu.opcodes["8.6"] = Opcode{name: "op8xy6", operation: cpu.op8xy6}
	cpu.opcodes["8.7"] = Opcode{name: "op8xy7", operation: cpu.op8xy7}
	cpu.opcodes["8.8"] = Opcode{name: "op8xyE", operation: cpu.op8xyE}

	cpu.opcodes["9"] = Opcode{name: "op9xy0", operation: cpu.op9xy0}
	cpu.opcodes["A"] = Opcode{name: "opAnnn", operation: cpu.opAnnn}
	cpu.opcodes["B"] = Opcode{name: "opBnnn", operation: cpu.opBnnn}
	cpu.opcodes["C"] = Opcode{name: "opCxkk", operation: cpu.opCxkk}
	cpu.opcodes["D"] = Opcode{name: "opDxyn", operation: cpu.opDxyn}

	cpu.opcodes["E.0"] = Opcode{name: "opEx9E", operation: cpu.opEx9E}
	cpu.opcodes["E.1"] = Opcode{name: "opExA1", operation: cpu.opExA1}

	cpu.opcodes["F.0"] = Opcode{name: "opFx07", operation: cpu.opFx07}
	cpu.opcodes["F.1"] = Opcode{name: "opFx0A", operation: cpu.opFx0A}
	cpu.opcodes["F.2"] = Opcode{name: "opFx15", operation: cpu.opFx15}
	cpu.opcodes["F.3"] = Opcode{name: "opFx18", operation: cpu.opFx18}
	cpu.opcodes["F.4"] = Opcode{name: "opFx1E", operation: cpu.opFx1E}
	cpu.opcodes["F.5"] = Opcode{name: "opFx29", operation: cpu.opFx29}
	cpu.opcodes["F.6"] = Opcode{name: "opFx33", operation: cpu.opFx33}
	cpu.opcodes["F.7"] = Opcode{name: "opFx55", operation: cpu.opFx55}
	cpu.opcodes["F.8"] = Opcode{name: "opFx65", operation: cpu.opFx65}

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
	return uint8((v >> 4) & 0xf)
}
func getkk(v uint16) uint8 {
	return uint8(v & 0xff)
}

func (cpu *Cpu) op0nnn() {
	fmt.Printf("%x\n", cpu.opcode)
	os.Exit(1)
}

/*
00E0 - CLS
Clear the display.
*/
func (cpu *Cpu) op00E0() {
	cpu.bus.display.clear()
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
	y := gety(cpu.opcode)
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
	y := gety(cpu.opcode)
	cpu.registers[x] = cpu.registers[y]
}

/*
8xy1 - OR Vx, Vy
Set Vx = Vx OR Vy.
Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A bitwise OR compares the corrseponding bits from two values, and if either bit is 1, then the same bit in the result is also 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy1() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] | cpu.registers[y]
}

/*
8xy2 - AND Vx, Vy
Set Vx = Vx AND Vy.
Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx. A bitwise AND compares the corrseponding bits from two values, and if both bits are 1, then the same bit in the result is also 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy2() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] & cpu.registers[y]
}

/*
8xy3 - XOR Vx, Vy
Set Vx = Vx XOR Vy.
Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx. An exclusive OR compares the corrseponding bits from two values, and if the bits are not both the same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
*/
func (cpu *Cpu) op8xy3() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
	cpu.registers[x] = cpu.registers[x] ^ cpu.registers[y]
}

/*
8xy4 - ADD Vx, Vy
Set Vx = Vx + Vy, set VF = carry.
The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
*/
func (cpu *Cpu) op8xy4() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
	res := uint16(cpu.registers[x]) + uint16(cpu.registers[y])
	cpu.registers[x] = cpu.registers[x] + cpu.registers[y]

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
	y := gety(cpu.opcode)
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
	cpu.registers[0xf] = cpu.registers[x] & 1
	cpu.registers[x] = cpu.registers[x] >> 1
}

/*
8xy7 - SUBN Vx, Vy
Set Vx = Vy - Vx, set VF = NOT borrow.
If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
*/
func (cpu *Cpu) op8xy7() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
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
	cpu.registers[0xf] = cpu.registers[x] >> 7
	cpu.registers[x] = cpu.registers[x] << 1
}

/*
9xy0 - SNE Vx, Vy
Skip next instruction if Vx != Vy.
The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
*/
func (cpu *Cpu) op9xy0() {
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
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
	n := uint16(getn(cpu.opcode))
	x := getx(cpu.opcode)
	y := gety(cpu.opcode)
	posX := int(cpu.registers[x])
	posY := int(cpu.registers[y])
	cpu.registers[0xf] = 0
	for i := uint16(0); i < n; i++ {
		data := cpu.bus.ram.read(cpu.I + i)
		//pointer to the curr pixel(bit) in the data
		stepPixel := 7
		posY = posY % cpu.bus.display.height
		for ; stepPixel >= 0; stepPixel-- {
			newPixelBit := (data >> stepPixel) & 0x1
			if newPixelBit == 1 {

				currPosX := posX + (7 - stepPixel)
				currPosX = currPosX % cpu.bus.display.width

				oldpixel := cpu.bus.display.getPixel(currPosX, posY)
				//indecates if the curr present pixel is black pixel or not
				oldPixelBit := uint8(0)
				if oldpixel.r == 255 && oldpixel.b == 255 && oldpixel.g == 255 {
					oldPixelBit = uint8(1)
				}

				//check if there is collision
				if oldPixelBit+newPixelBit == 2 {
					cpu.registers[0xf] = 1
				}
				c := Black

				if (oldPixelBit ^ newPixelBit) == 1 {
					c = White

				}
				cpu.bus.display.setPixel(currPosX, posY, c)
			}

		}
		posY++
	}
}

/*
Ex9E - SKP Vx
Skip next instruction if key with the value of Vx is pressed.
Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
*/
func (cpu *Cpu) opEx9E() {
	x := getx(cpu.opcode)
	//keyCode := fmt.Sprintf("%X", cpu.registers[x])
	keyCode := cpu.bus.joypad.keysMapper[cpu.registers[x]]
	if cpu.bus.joypad.keys[keyCode].pressed {
		cpu.PC = cpu.PC + 2
	}
}

/*
ExA1 - SKNP Vx
Skip next instruction if key with the value of Vx is not pressed.
Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
*/
func (cpu *Cpu) opExA1() {
	x := getx(cpu.opcode)
	//keyCode := fmt.Sprintf("%X", cpu.registers[x])
	keyCode := cpu.bus.joypad.keysMapper[cpu.registers[x]]
	if !cpu.bus.joypad.keys[keyCode].pressed {
		cpu.PC = cpu.PC + 2
	}
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
	key := cpu.bus.joypad.getKey()
	cpu.registers[x] = key.value
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
	cpu.I = uint16(5 * cpu.registers[x])

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
	cpu.I = cpu.I + x + 1
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
	cpu.I = cpu.I + x + 1
}
