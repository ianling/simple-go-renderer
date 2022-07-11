package drawing

import (
	"math"
	"sync"
)

type LineSegment struct {
	Drawable
	Color   Color
	VertexA Vector3
	VertexB Vector3
}

func NewLineSegment(x1, y1, x2, y2 float64, r, g, b, a uint8) LineSegment {
	originX := x1
	originY := y1
	return LineSegment{
		Drawable: NewDrawable(Vector3{X: originX, Y: originY}),
		Color:    Color{R: r, G: g, B: b, A: a},
		VertexA:  Vector3{X: x1, Y: y1},
		VertexB:  Vector3{X: x2, Y: y2},
	}
}

type LineSegments []LineSegment

type Liner interface {
	Lines() LineSegments
}

func (line *LineSegment) Draw(screenBuffer []byte) {
	aX, aY, _ := line.ApplyTransformations(line.VertexA.X, line.VertexA.Y, line.VertexA.Z)
	bX, bY, _ := line.ApplyTransformations(line.VertexB.X, line.VertexB.Y, line.VertexB.Z)

	dx := bX - aX
	dy := bY - aY
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
	x := aX
	y := aY
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
