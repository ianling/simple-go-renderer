package drawing

type Triangle struct {
	Drawable
	VertexA Vector3
	VertexB Vector3
	VertexC Vector3
	Color   Color
}

func NewTriangle(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64, r, g, b, a uint8) Triangle {
	originX := (x1 + x2 + x3) / 3
	originY := (y1 + y2 + y3) / 3
	originZ := (z1 + z2 + z3) / 3

	return Triangle{
		Drawable: NewDrawable(Vector3{X: originX, Y: originY, Z: originZ}),
		Color:    Color{R: r, G: g, B: b, A: a},
		VertexA:  Vector3{X: x1, Y: y1, Z: z1},
		VertexB:  Vector3{X: x2, Y: y2, Z: z2},
		VertexC:  Vector3{X: x3, Y: y3, Z: z3},
	}
}

func (triangle *Triangle) Lines() LineSegments {
	return LineSegments{
		// A->B
		{
			Drawable: triangle.Drawable,
			Color:    triangle.Color,
			VertexA:  triangle.VertexA,
			VertexB:  triangle.VertexB,
		},
		// A->C
		{
			Drawable: triangle.Drawable,
			Color:    triangle.Color,
			VertexA:  triangle.VertexA,
			VertexB:  triangle.VertexC,
		},
		// B->C
		{
			Drawable: triangle.Drawable,
			Color:    triangle.Color,
			VertexA:  triangle.VertexB,
			VertexB:  triangle.VertexC,
		},
	}
}

func (triangle *Triangle) Draw(screenBuffer []byte) {
	triangle.Lines().Draw(screenBuffer)
}
