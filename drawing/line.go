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

func NewLineSegment(vertexA, vertexB Vector3, drawable Drawable, color Color) LineSegment {
	return LineSegment{
		Drawable: drawable,
		Color:    color,
		VertexA:  vertexA,
		VertexB:  vertexB,
	}
}

type LineSegments []LineSegment

func (line *LineSegment) Draw(screenBuffer []byte) {
	aX, aY, aZ := line.ApplyTransformations(line.VertexA.X, line.VertexA.Y, line.VertexA.Z)
	bX, bY, bZ := line.ApplyTransformations(line.VertexB.X, line.VertexB.Y, line.VertexB.Z)

	dx := bX - aX
	dy := bY - aY
	dz := bZ - aZ
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
	dz = dz / step
	x := aX
	y := aY
	z := aZ
	for ii := float64(1); ii <= step; ii += 1 {
		//SetPixel(screenBuffer, int(x), int(y), line.Color.R, line.Color.G, line.Color.B, uint8(math.Min(math.Max(float64(line.Color.A)*(z/100), 0), 255)))
		SetPixel(screenBuffer, int(x), int(y), line.Color.R, line.Color.G, line.Color.B, line.Color.A)
		x += dx
		y += dy
		z += dz
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
