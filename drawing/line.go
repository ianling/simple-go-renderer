package drawing

import (
	"math"
	"sync"
)

type LineSegment struct {
	Drawable
	Color   Color
	VertexA Vector2
	VertexB Vector2
}

func NewLineSegment(x1, y1, x2, y2 int, r, g, b, a uint8) LineSegment {
	originX := x1
	originY := y1
	return LineSegment{
		Drawable: NewDrawable(Vector2{X: originX, Y: originY}),
		Color:    Color{R: r, G: g, B: b, A: a},
		VertexA:  Vector2{X: x1, Y: y1},
		VertexB:  Vector2{X: x2, Y: y2},
	}
}

type LineSegments []LineSegment

type Liner interface {
	Lines() LineSegments
}

func (line *LineSegment) Draw(screenBuffer []byte) {
	aX, aY := line.ApplyTransformations(line.VertexA.X, line.VertexA.Y)
	bX, bY := line.ApplyTransformations(line.VertexB.X, line.VertexB.Y)

	// debug
	//SetPixel(screenBuffer, int(aX), int(aY), line.Color.R, line.Color.G, line.Color.B, line.Color.A)
	//SetPixel(screenBuffer, int(bX), int(bY), line.Color.R, line.Color.G, line.Color.B, line.Color.A)
	//return

	dx := float64(bX - aX)
	dy := float64(bY - aY)
	absDx := math.Abs(dx)
	absDy := math.Abs(dy)
	var step float64
	if absDx >= absDy {
		step = absDx
	} else {
		step = absDy
	}
	dx = dx / step
	dy = dy / step
	x := float64(aX)
	y := float64(aY)
	for ii := float64(1); ii <= step; ii += 1 {
		SetPixel(screenBuffer, int(x), int(y), line.Color.R, line.Color.G, line.Color.B, line.Color.A)
		x += dx
		y += dy
	}
}

func (lines LineSegments) drawAsync(screenBuffer []byte) {
	var wg sync.WaitGroup
	wg.Add(len(lines))

	for ii := range lines {
		ii := ii
		go func() {
			lines[ii].Draw(screenBuffer)
			wg.Done()
		}()
	}

	wg.Wait()
}

func (lines LineSegments) drawSync(screenBuffer []byte) {
	for ii := range lines {
		lines[ii].Draw(screenBuffer)
	}
}

func (lines LineSegments) Draw(screenBuffer []byte) {
	lines.drawSync(screenBuffer)
}
