package bus

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Key struct {
	name  string
	value uint8
	//indicates if the key is currently pressed
	pressed bool
}
type Joypad struct {
	keys map[string]Key
}

func (joypad *Joypad) config() {
	joypad.keys = make(map[string]Key)
	joypad.keys["0"] = Key{name: "0", value: 0}
	joypad.keys["1"] = Key{name: "1", value: 1}
	joypad.keys["2"] = Key{name: "2", value: 2}
	joypad.keys["3"] = Key{name: "3", value: 3}
	joypad.keys["4"] = Key{name: "4", value: 4}
	joypad.keys["5"] = Key{name: "5", value: 5}
	joypad.keys["6"] = Key{name: "6", value: 6}
	joypad.keys["7"] = Key{name: "7", value: 7}
	joypad.keys["8"] = Key{name: "8", value: 8}
	joypad.keys["9"] = Key{name: "9", value: 9}
	joypad.keys["A"] = Key{name: "A", value: 0xA}
	joypad.keys["B"] = Key{name: "B", value: 0xB}
	joypad.keys["C"] = Key{name: "C", value: 0xC}
	joypad.keys["D"] = Key{name: "D", value: 0xD}
	joypad.keys["E"] = Key{name: "E", value: 0xE}
	joypad.keys["F"] = Key{name: "F", value: 0xF}
}
func (joypad *Joypad) getKey() Key {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch et := event.(type) {

			case *sdl.KeyboardEvent:
				if event.GetType() == sdl.KEYDOWN {
					if key, ok := joypad.keys[sdl.GetKeyName(et.Keysym.Sym)]; ok {
						return key
					}

				}

				break
			}
		}
		sdl.Delay(16)

	}
}
