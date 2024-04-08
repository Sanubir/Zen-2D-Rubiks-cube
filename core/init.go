package core

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func InitWindow(windowTitle string, windowWidth, windowHeight int32) *sdl.Window {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %v", err)
	}
	window, err := sdl.CreateWindow(
		windowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE,
	)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}
	return window
}

func InitRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %v", err)
	}
	return renderer
}

func (g *Graphics) Shutdown() {
	g.Renderer.Destroy()
	g.Window.Destroy()
	sdl.Quit()
	fmt.Println("Graphics closed")
}
