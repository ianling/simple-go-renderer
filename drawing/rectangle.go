package drawing

type Rectangle struct {
	Drawable
	Position Vector2
	Width    int
	Height   int
	Color    Color
}

func NewRectangle(x, y, width, height int, r, g, b, a uint8) Rectangle {
	originX := x + (width / 2)
	originY := y + (height / 2)
	return Rectangle{
		Drawable: NewDrawable(Vector2{X: originX, Y: originY}),
		Color:    Color{R: r, G: g, B: b, A: a},
		Position: Vector2{X: x, Y: y},
		Width:    width,
		Height:   height,
	}
}

func (rectangle *Rectangle) Triangles() []Triangle {
	return []Triangle{
		{
			Drawable: rectangle.Drawable,
			VertexA:  Vector2{X: rectangle.Position.X, Y: rectangle.Position.Y},
			VertexB:  Vector2{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y},
			VertexC:  Vector2{X: rectangle.Position.X, Y: rectangle.Position.Y + rectangle.Height},
			Color:    rectangle.Color,
		},
		{
			Drawable: rectangle.Drawable,
			VertexA:  Vector2{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y + rectangle.Height},
			VertexB:  Vector2{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y},
			VertexC:  Vector2{X: rectangle.Position.X, Y: rectangle.Position.Y + rectangle.Height},
			Color:    rectangle.Color,
		},
	}
}

func (rectangle *Rectangle) Draw(screenBuffer []byte) {
	for _, triangle := range rectangle.Triangles() {
		triangle.Draw(screenBuffer)
	}
}
