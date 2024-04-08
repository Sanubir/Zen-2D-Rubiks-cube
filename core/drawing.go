package core

import (
	"log"
	"math"
	"rubikGUI/rubik"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func drawText(renderer *sdl.Renderer, color sdl.Color, fontSize int, text string, x, y int32) {
	ttf.Init()
	font, err := ttf.OpenFont("fonts/mono-condensed-black.ttf", fontSize)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	defer font.Close()

	surface, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		log.Fatalf("Failed to render text: %v", err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Failed to create texture: %v", err)
	}
	defer texture.Destroy()

	_, _, width, height, _ := texture.Query()
	dstRect := sdl.Rect{X: x - width/2, Y: y - height/2, W: width, H: height}
	renderer.Copy(texture, nil, &dstRect)
}

// Make sure vertices are in order: top left, top right, bottom left, right
func drawQuad(renderer *sdl.Renderer, vertices [4]sdl.Vertex) {
	// Divide 4 quadrangle points into two triangles
	triangles := [][]sdl.Vertex{
		{vertices[0], vertices[1], vertices[2]},
		{vertices[2], vertices[3], vertices[1]},
	}
	// Render triangles
	renderer.RenderGeometry(nil, triangles[0], nil)
	renderer.RenderGeometry(nil, triangles[1], nil)
}

func drawCube(renderer *sdl.Renderer, cube *rubik.Cube, cfg *Config) {
	middleX := cfg.WindowWidth / 2
	// middleY := cfg.WindowHeight / 2
	faceSizeX := cfg.WindowWidth / 7
	faceSizeY := cfg.WindowHeight / 5
	faceSize := int32(math.Min(float64(faceSizeX), float64(faceSizeY)))
	stickerSize := faceSize / int32(cfg.Size)
	gap := stickerSize / 5

	// Drawing top face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.T[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX - faceSize/2 + int32(i)*gap + int32(i)*stickerSize,
				Y: int32(0.75*float64(faceSize)) + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})
		}
	}
	// Drawing back face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.B[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX - int32(3*float64(faceSize)) + int32(i)*gap + int32(i)*stickerSize,
				Y: 2*faceSize + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})
		}
	}
	// Drawing left face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.L[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX - int32(1.75*float64(faceSize)) + int32(i)*gap + int32(i)*stickerSize,
				Y: 2*faceSize + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})
		}
	}
	// Drawing front face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.F[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX - faceSize/2 + int32(i)*gap + int32(i)*stickerSize,
				Y: 2*faceSize + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})
		}
	}
	// Drawing right face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.R[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX + int32(0.75*float64(faceSize)) + int32(i)*gap + int32(i)*stickerSize,
				Y: 2*faceSize + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})
		}
	}
	// Drawing down face
	for i := 0; i < cfg.Size; i++ {
		for j := 0; j < cfg.Size; j++ {
			R, G, B := getColor(cube.D[i][j], cfg.Colors)
			renderer.SetDrawColor(R, G, B, 255)
			renderer.FillRect(&sdl.Rect{
				X: middleX - faceSize/2 + int32(i)*gap + int32(i)*stickerSize,
				Y: int32(3.25*float64(faceSize)) + int32(j)*gap + int32(j)*stickerSize,
				W: stickerSize, H: stickerSize,
			})

		}
	}

}

func getColor(color byte, colors Colors) (R, G, B uint8) {
	switch color {
	case 'y':
		return colors.Yellow.R, colors.Yellow.G, colors.Yellow.B
	case 'b':
		return colors.Blue.R, colors.Blue.G, colors.Blue.B
	case 'r':
		return colors.Red.R, colors.Red.G, colors.Red.B
	case 'g':
		return colors.Green.R, colors.Green.G, colors.Green.B
	case 'o':
		return colors.Orange.R, colors.Orange.G, colors.Orange.B
	case 'w':
		return colors.White.R, colors.White.G, colors.White.B
	}
	// log.Panic("Wrong color in getColor()")
	return 0, 0, 0
}
