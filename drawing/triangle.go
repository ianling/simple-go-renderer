package drawing

type Triangle struct {
	Drawable
	VertexA Vector2
	VertexB Vector2
	VertexC Vector2
	Color   Color
}

func NewTriangle(x1, y1, x2, y2, x3, y3 int, r, g, b, a uint8) Triangle {
	return Triangle{
		Drawable: NewDrawable(),
		Color:    Color{R: r, G: g, B: b, A: a},
		VertexA:  Vector2{X: x1, Y: y1},
		VertexB:  Vector2{X: x2, Y: y2},
		VertexC:  Vector2{X: x3, Y: y3},
	}
}

func (triangle *Triangle) Lines() LineSegments {
	return LineSegments{
		{
			Drawable: triangle.Drawable,
			Color:    triangle.Color,
			VertexA:  triangle.VertexA,
			VertexB:  triangle.VertexB,
		},
		{
			Drawable: triangle.Drawable,
			Color:    triangle.Color,
			VertexA:  triangle.VertexA,
			VertexB:  triangle.VertexC,
		},
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
