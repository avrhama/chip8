package bus

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Ram struct {
	mem [4096]uint8
}

func (ram *Ram) write(address uint16, data uint8) {
	if address < 4096 {
		ram.mem[address] = data
	}

}
func (ram *Ram) read(address uint16) uint8 {
	if address < 4096 {
		return ram.mem[address]
	}
	return 0
}

func (ram *Ram) config() {
	//fill the mem with digit representtions
	//0
	ram.mem[0] = 0xF0
	ram.mem[1] = 0x90
	ram.mem[2] = 0x90
	ram.mem[3] = 0x90
	ram.mem[4] = 0xF0
	//1
	ram.mem[5] = 0x20
	ram.mem[6] = 0x60
	ram.mem[7] = 0x20
	ram.mem[8] = 0x20
	ram.mem[9] = 0x70
	//2
	ram.mem[10] = 0xF0
	ram.mem[11] = 0x10
	ram.mem[12] = 0xF0
	ram.mem[13] = 0x80
	ram.mem[14] = 0xF0
	//3
	ram.mem[15] = 0xF0
	ram.mem[16] = 0x10
	ram.mem[17] = 0xF0
	ram.mem[18] = 0x10
	ram.mem[19] = 0xF0
	//4
	ram.mem[20] = 0x90
	ram.mem[21] = 0x90
	ram.mem[22] = 0xF0
	ram.mem[23] = 0x10
	ram.mem[24] = 0x10
	//5
	ram.mem[25] = 0xF0
	ram.mem[26] = 0x80
	ram.mem[27] = 0xF0
	ram.mem[28] = 0x10
	ram.mem[29] = 0xF0
	//6
	ram.mem[30] = 0xF0
	ram.mem[31] = 0x80
	ram.mem[32] = 0xF0
	ram.mem[33] = 0x90
	ram.mem[34] = 0xF0
	//7
	ram.mem[35] = 0xF0
	ram.mem[36] = 0x10
	ram.mem[37] = 0x20
	ram.mem[38] = 0x40
	ram.mem[39] = 0x40
	//8
	ram.mem[40] = 0xF0
	ram.mem[41] = 0x90
	ram.mem[42] = 0xF0
	ram.mem[43] = 0x90
	ram.mem[44] = 0xF0
	//9
	ram.mem[45] = 0xF0
	ram.mem[46] = 0x90
	ram.mem[47] = 0xF0
	ram.mem[48] = 0x10
	ram.mem[49] = 0xF0
	//A
	ram.mem[50] = 0xF0
	ram.mem[51] = 0x90
	ram.mem[52] = 0xF0
	ram.mem[53] = 0x90
	ram.mem[54] = 0x90
	//B
	ram.mem[55] = 0xE0
	ram.mem[56] = 0x90
	ram.mem[57] = 0xE0
	ram.mem[58] = 0x90
	ram.mem[59] = 0xE0
	//C
	ram.mem[60] = 0xF0
	ram.mem[61] = 0x80
	ram.mem[62] = 0x80
	ram.mem[63] = 0x80
	ram.mem[64] = 0xF0
	//D
	ram.mem[65] = 0xE0
	ram.mem[66] = 0x90
	ram.mem[67] = 0x90
	ram.mem[68] = 0x90
	ram.mem[69] = 0xE0
	//E
	ram.mem[70] = 0xF0
	ram.mem[71] = 0x80
	ram.mem[72] = 0xF0
	ram.mem[73] = 0x80
	ram.mem[74] = 0xF0
	//F
	ram.mem[75] = 0xF0
	ram.mem[76] = 0x80
	ram.mem[77] = 0xF0
	ram.mem[78] = 0x80
	ram.mem[79] = 0x80

}
func (ram *Ram) loadRom(path string) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	size := uint16(len(data))
	for i := uint16(0); i < size; i++ {
		ram.write(i+0x200, data[i])
	}
}
