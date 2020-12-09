package bus

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type color struct {
	r, g, b byte
}

func (display *Display) setPixel(x, y int, c color) {
	index := (y*display.height + x) * 4

	display.pixels[index] = c.r
	display.pixels[index+1] = c.g
	display.pixels[index+2] = c.b
}

type Display struct {
	width    int
	height   int
	window   *sdl.Window
	renderer *sdl.Renderer
	tex      *sdl.Texture
	pixels   []byte
}

func (display *Display) draw() {
	display.tex.Update(nil, display.pixels, display.width*4)
	display.renderer.Copy(display.tex, nil, nil)
	display.renderer.Present()
}

func (display *Display) config() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	display.window, err = sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(display.width), int32(display.height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		display.turnOff(0)
		return
	}

	display.renderer, err = sdl.CreateRenderer(display.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		display.turnOff(1)
		return
	}

	display.tex, err = display.renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(display.width), int32(display.height))
	if err != nil {
		fmt.Println(err)
		display.turnOff(2)
		return
	}

	display.pixels = make([]byte, int32(display.width)*int32(display.height)*4)
}
func (display *Display) test() {
	for y := 0; y < display.height; y++ {
		for x := 0; x < display.width; x++ {
			display.setPixel(x, y, color{255, 0, 0})
		}
	}
	display.tex.Update(nil, display.pixels, display.width*4)
	display.renderer.Copy(display.tex, nil, nil)
	display.renderer.Present()
}

//destroys the sdl components
func (display *Display) turnOff(flag uint8) {

	if flag > 2 {
		display.tex.Destroy()
	}
	if flag > 1 {
		display.renderer.Destroy()
	}
	if flag > 0 {
		display.window.Destroy()
	}

	sdl.Quit()
}
