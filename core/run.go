package core

import (
	"rubikGUI/rubik"

	"github.com/veandco/go-sdl2/sdl"
)

func Run(cfg *Config) int {
	G := new(Graphics)
	G.Window = InitWindow(cfg.WindowTitle, cfg.WindowWidth, cfg.WindowHeight)
	G.Renderer = InitRenderer(G.Window)
	defer G.Shutdown()

	cube := new(rubik.Cube)
	cube.InitCube(cfg.Size)

	var currentTime string
	go updateCurrentTime(&currentTime)

	// main loop
	running := true
	for running {
		HandleEvents(&running, cfg, cube)

		// Color background
		G.Renderer.SetDrawColor(cfg.BG.R, cfg.BG.G, cfg.BG.B, 10)
		G.Renderer.Clear()

		// ------------------Drawing GUI elements------------------

		drawCube(G.Renderer, cube, cfg)

		drawText(G.Renderer, cfg.FG, int(cfg.WindowWidth)/42,
			"Moves [key->action]: J->U, F->U', I->R, K->R', E->L', D->L, ↓->x', ↑->x",
			int32(float64(cfg.WindowWidth)*0.5),
			int32(float64(cfg.WindowHeight)*0.95),
		)

		drawText(G.Renderer, cfg.FG, int(cfg.WindowWidth)/30,
			"Current time:",
			int32(float64(cfg.WindowWidth)*0.8),
			int32(float64(cfg.WindowHeight)*0.22),
		)

		drawText(G.Renderer, cfg.FG, int(cfg.WindowWidth)/42,
			currentTime,
			int32(float64(cfg.WindowWidth)*0.8),
			int32(float64(cfg.WindowHeight)*0.28),
		)

		// -----------------------------------------------------
		// Update the screen
		G.Renderer.Present()
		sdl.Delay(uint32(cfg.Frametime))
	}

	return 0
}
