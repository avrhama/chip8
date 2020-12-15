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
	keys       map[sdl.Scancode]Key
	keysMapper [16]sdl.Scancode
}

func (joypad *Joypad) config() {
	joypad.keys = make(map[sdl.Scancode]Key)

	joypad.keys[sdl.SCANCODE_0] = Key{name: "0", value: 0}
	joypad.keys[sdl.SCANCODE_1] = Key{name: "1", value: 1}
	joypad.keys[sdl.SCANCODE_2] = Key{name: "2", value: 2}
	joypad.keys[sdl.SCANCODE_3] = Key{name: "3", value: 3}
	joypad.keys[sdl.SCANCODE_4] = Key{name: "4", value: 4}
	joypad.keys[sdl.SCANCODE_5] = Key{name: "5", value: 5}
	joypad.keys[sdl.SCANCODE_6] = Key{name: "6", value: 6}
	joypad.keys[sdl.SCANCODE_7] = Key{name: "7", value: 7}
	joypad.keys[sdl.SCANCODE_8] = Key{name: "8", value: 8}
	joypad.keys[sdl.SCANCODE_9] = Key{name: "9", value: 9}
	joypad.keys[sdl.SCANCODE_A] = Key{name: "A", value: 0xA}
	joypad.keys[sdl.SCANCODE_B] = Key{name: "B", value: 0xB}
	joypad.keys[sdl.SCANCODE_C] = Key{name: "C", value: 0xC}
	joypad.keys[sdl.SCANCODE_D] = Key{name: "D", value: 0xD}
	joypad.keys[sdl.SCANCODE_E] = Key{name: "E", value: 0xE}
	joypad.keys[sdl.SCANCODE_F] = Key{name: "F", value: 0xF}

	for key, v := range joypad.keys {
		joypad.keysMapper[v.value] = key
	}
}
func (joypad *Joypad) getKey() Key {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch et := event.(type) {

			case *sdl.KeyboardEvent:
				if event.GetType() == sdl.KEYDOWN {
					if key, ok := joypad.keys[sdl.GetScancodeFromKey(et.Keysym.Sym)]; ok {
						return key
					}
				}

				break
			}
		}
		sdl.Delay(16)

	}
}
