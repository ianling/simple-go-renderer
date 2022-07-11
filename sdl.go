package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
	"time"
)

var sdlPerformanceCounter uint64
var frameCounter float64
var deltaCounter time.Duration

var initOnce sync.Once
var initOnceFunc = func() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func newSDLWindow(width, height int32) (*sdl.Window, *sdl.Renderer, *sdl.Texture, func()) {
	initOnce.Do(initOnceFunc)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, windowWidth, windowHeight)
	if err != nil {
		panic(err)
	}

	return window, renderer, texture, func() {
		texture.Destroy()
		renderer.Destroy()
		window.Destroy()
	}
}

// deltaTime returns the time delta between frames
func deltaTime() time.Duration {
	last := sdlPerformanceCounter
	sdlPerformanceCounter = sdl.GetPerformanceCounter()
	return time.Duration(float64(time.Millisecond) * float64((sdlPerformanceCounter-last)*1000) / float64(sdl.GetPerformanceFrequency()))
}

func handleEvents() (continueRunning bool) {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			return KeyboardHandler(event)
		}
	}

	return true
}

func exit(sdlCleanupFunc func()) {
	sdlCleanupFunc()
	sdl.Quit()
}
