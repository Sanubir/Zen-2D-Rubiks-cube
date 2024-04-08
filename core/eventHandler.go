package core

import (
	"fmt"
	"rubikGUI/rubik"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleEvents(running *bool, cfg *Config, cube *rubik.Cube) {
	// Track the state of each key
	keyState := make(map[sdl.Keycode]bool)

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			*running = false
		case *sdl.WindowEvent:
			windowEvent := event.(*sdl.WindowEvent)
			if windowEvent.Event == sdl.WINDOWEVENT_RESIZED {
				cfg.WindowWidth = windowEvent.Data1
				cfg.WindowHeight = windowEvent.Data2
			}
		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			keys := ""

			// Modifier keys
			switch t.Keysym.Mod {
			case sdl.KMOD_LALT:
				keys += "Left Alt"
			case sdl.KMOD_LCTRL:
				keys += "Left Control"
			case sdl.KMOD_LSHIFT:
				keys += "Left Shift"
			case sdl.KMOD_LGUI:
				keys += "Left Meta or Windows key"
			case sdl.KMOD_RALT:
				keys += "Right Alt"
			case sdl.KMOD_RCTRL:
				keys += "Right Control"
			case sdl.KMOD_RSHIFT:
				keys += "Right Shift"
			case sdl.KMOD_RGUI:
				keys += "Right Meta or Windows key"
			case sdl.KMOD_NUM:
				keys += "Num Lock"
			case sdl.KMOD_CAPS:
				keys += "Caps Lock"
			case sdl.KMOD_MODE:
				keys += "AltGr Key"
			}

			switch t.State {
			case sdl.RELEASED:
				// Update the key state to released
				keyState[keyCode] = false
			case sdl.PRESSED:
				// Check if the key was already pressed
				if !keyState[keyCode] {
					// Update the key state to pressed
					keyState[keyCode] = true

					switch keyCode {
					case sdl.K_j:
						cube.Rotate_U()
					case sdl.K_f:
						cube.Rotate_Uc()
					case sdl.K_k:
						cube.Rotate_Rc()
					case sdl.K_i:
						cube.Rotate_R()
					case sdl.K_d:
						cube.Rotate_L()
					case sdl.K_e:
						cube.Rotate_Lc()
					case sdl.K_DOWN:
						cube.Rotate_DownCube()
					case sdl.K_UP:
						cube.Rotate_UpCube()
					}

					if keyCode < 10000 {
						if keys != "" {
							keys += " + "
						}
						// If the key is held down, this will fire
						if t.Repeat > 0 {
							keys += string(keyCode) + " repeating"
						} else {
							keys += string(keyCode) + " pressed"
						}
					}
					if keys != "" {
						// fmt.Println(keys)
					}
				}
			}
		}
	}
}

func updateCurrentTime(currentTime *string) {
	for {
		*currentTime = time.Now().Format("15:04:05")
		time.Sleep(100 * time.Millisecond)
	}
}
