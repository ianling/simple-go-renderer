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

var lines = drawing.LineSegments{
	drawing.NewLineSegment(300, 300, 250, 250, 255, 0, 0, 255),
	drawing.NewLineSegment(100, 100, 200, 200, 0, 255, 0, 255),
	drawing.NewLineSegment(100, 100, 650, 200, 255, 255, 255, 255),
	drawing.NewLineSegment(100, 100, 200, 550, 255, 255, 255, 255),
	drawing.NewLineSegment(300, 300, 300, 400, 160, 100, 255, 255),
}

var triangles = []drawing.Triangle{
	drawing.NewTriangle(50, 250, 0, 150, 250, 0, 100, 100, 0, 255, 0, 0, 255),
	drawing.NewTriangle(250, 250, 0, 350, 250, 0, 300, 100, 0, 255, 0, 0, 255),
	drawing.NewTriangle(450, 250, 0, 550, 250, 0, 500, 100, 0, 255, 0, 0, 255),
	drawing.NewTriangle(650, 250, 0, 750, 250, 0, 700, 100, 0, 255, 0, 0, 255),
}

var rectangles = []drawing.Rectangle{
	drawing.NewRectangle(50, 350, 0, 75, 40, 255, 0, 255, 255),
	drawing.NewRectangle(250, 350, 0, 75, 40, 255, 0, 255, 255),
	drawing.NewRectangle(450, 350, 0, 75, 40, 255, 0, 255, 255),
	drawing.NewRectangle(650, 350, 0, 75, 40, 255, 0, 255, 255),
}

var cube = drawing.NewCube(175, 175, 175, 150, 150, 150, drawing.Color{R: 255, G: 0, B: 0, A: 255})

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

		//lines[4].Rotation += 0.06
		triangles[0].Rotation.Z += 0.06
		triangles[1].Rotation.X += 0.06
		triangles[2].Rotation.Y += 0.06
		triangles[3].Rotation.X += 0.06
		triangles[3].Rotation.Y += 0.04
		triangles[3].Rotation.Z += 0.03
		rectangles[0].Rotation.Z += 0.07
		rectangles[1].Rotation.X += 0.07
		rectangles[2].Rotation.Y += 0.07
		rectangles[3].Rotation.X += 0.07
		rectangles[3].Rotation.Y += 0.04
		rectangles[3].Rotation.Z += 0.03

		//cube.Rotation.X += 0.03
		//cube.Rotation.Y += 0.03

		// track mouse position with a line for fun
		mouseX32, mouseY32, _ := sdl.GetMouseState()
		mouseX, mouseY := constrainWithinWindow(int(mouseX32), int(mouseY32))
		lines[0].VertexB.X = float64(mouseX)
		lines[0].VertexB.Y = float64(mouseY)

		// get a byte array from our render texture so we can fill in pixels
		if screenBuffer, _, err = texture.Lock(&screenRect); err != nil {
			panic(err)
		}

		clearBuffer(screenBuffer)

		//lines.Draw(screenBuffer)

		for _, triangle := range triangles {
			triangle.Draw(screenBuffer)
		}

		for _, rectangle := range rectangles {
			rectangle.Draw(screenBuffer)
		}

		//cube.Draw(screenBuffer)

		texture.Unlock()

		if err = renderer.Copy(texture, &screenRect, &screenRect); err != nil {
			panic(err)
		}
		renderer.Present()

		frameCounter += 1
	}
}
