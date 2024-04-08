package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R, G, B uint8
}

type Colors struct {
	Yellow, Blue, Red, Green, Orange, White sdl.Color
}

type Config struct {
	WindowTitle  string
	WindowWidth  int32
	WindowHeight int32
	BG           sdl.Color
	FG           sdl.Color
	Colors       Colors
	Framerate    float64
	Frametime    float64
	Size         int
}

type Graphics struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}
