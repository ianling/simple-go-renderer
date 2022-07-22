package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"simplerenderer/drawing"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var screenRect = sdl.Rect{
	W: windowWidth,
	H: windowHeight,
}

var cube = drawing.NewCube(175, 175, -150, 150, 150, 150, drawing.Color{R: 255, G: 0, B: 0, A: 255})
var pyramid = drawing.NewPyramid(475, 175, -50, 150, 150, 250, drawing.Color{R: 255, G: 0, B: 0, A: 255})

var running bool
var paused bool

func main() {
	_, renderer, texture, sdlCleanupFunc := newSDLWindow(windowWidth, windowHeight)
	defer exit(sdlCleanupFunc)

	var screenBuffer []byte
	var err error
	for running = true; running; running = handleEvents() {
		// track average FPS over each second, print once per second
		deltaCounter += deltaTime()
		if deltaCounter.Seconds() >= 1 {
			fmt.Println(int(frameCounter / deltaCounter.Seconds()))
			deltaCounter = 0
			frameCounter = 0
		}

		if paused {
			continue
		}

		cube.Rotation.X += 0.04
		cube.Rotation.Y += 0.03
		cube.Rotation.Z += 0.02

		pyramid.Rotation.X += 0.02
		pyramid.Rotation.Y += 0.03
		pyramid.Rotation.Z += 0.04

		// track mouse position with a line for fun
		//mouseX32, mouseY32, _ := sdl.GetMouseState()
		//mouseX, mouseY := constrainWithinWindow(int(mouseX32), int(mouseY32))
		//lines[0].VertexB.X = float64(mouseX)
		//lines[0].VertexB.Y = float64(mouseY)

		// get a byte array from our render texture so we can fill in pixels
		if screenBuffer, _, err = texture.Lock(&screenRect); err != nil {
			panic(err)
		}

		clearBuffer(screenBuffer)

		cube.Draw(screenBuffer)
		pyramid.Draw(screenBuffer)

		texture.Unlock()

		if err = renderer.Copy(texture, &screenRect, &screenRect); err != nil {
			panic(err)
		}
		renderer.Present()

		frameCounter += 1
	}
}
