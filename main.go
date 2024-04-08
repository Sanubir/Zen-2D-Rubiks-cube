package main

import (
	"os"
	"rubikGUI/core"

	"github.com/veandco/go-sdl2/sdl"
)

func config() *core.Config {
	cfg := &core.Config{
		WindowTitle:  "RubikGUI",
		WindowWidth:  800,
		WindowHeight: 600,
		BG:           sdl.Color{R: 0, G: 0, B: 0},
		FG:           sdl.Color{R: 255, G: 255, B: 255},
		Colors: core.Colors{ // Cube colors
			Yellow: sdl.Color{R: 255, G: 255, B: 0},
			Blue:   sdl.Color{R: 0, G: 0, B: 255},
			Red:    sdl.Color{R: 255, G: 0, B: 0},
			Green:  sdl.Color{R: 0, G: 221, B: 0},
			Orange: sdl.Color{R: 255, G: 170, B: 0},
			White:  sdl.Color{R: 255, G: 255, B: 255},
		},
		Framerate: 50,
		Size:      3,
	}
	cfg.Frametime = 1000 / cfg.Framerate

	return cfg
}

func main() {
	cfg := config()
	// os.Exit(..) must run AFTER sdl.Main(..) below; so keep track of exit
	// status manually outside the closure passed into sdl.Main(..) below
	var exitcode int
	sdl.Main(func() {
		exitcode = core.Run(cfg)
	})
	// os.Exit(..) must run here! If run in sdl.Main(..) above, it will cause
	// premature quitting of sdl.Main(..) function; resource cleaning deferred
	// calls/closing of channels may never run
	os.Exit(exitcode)
}
