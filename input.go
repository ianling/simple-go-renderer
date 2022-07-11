package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func KeyboardHandler(event *sdl.KeyboardEvent) bool {
	if event.Type != sdl.KEYDOWN {
		return true
	}

	switch event.Keysym.Sym {
	case sdl.K_SPACE:
		paused = !paused
	case sdl.K_ESCAPE:
		return false
	}

	return true
}
