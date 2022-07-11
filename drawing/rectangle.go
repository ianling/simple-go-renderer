package drawing

type Rectangle struct {
	Drawable
	Position Vector3
	Width    float64
	Height   float64
	Color    Color
}

func NewRectangle(x, y, z, width, height float64, r, g, b, a uint8) Rectangle {
	originX := x + (width / 2)
	originY := y + (height / 2)
	return Rectangle{
		Drawable: NewDrawable(Vector3{X: originX, Y: originY, Z: z}),
		Color:    Color{R: r, G: g, B: b, A: a},
		Position: Vector3{X: x, Y: y, Z: z},
		Width:    width,
		Height:   height,
	}
}

func (rectangle *Rectangle) Triangles() []Triangle {
	return []Triangle{
		{
			Drawable: rectangle.Drawable,
			VertexA:  Vector3{X: rectangle.Position.X, Y: rectangle.Position.Y, Z: rectangle.Position.Z},
			VertexB:  Vector3{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y, Z: rectangle.Position.Z},
			VertexC:  Vector3{X: rectangle.Position.X, Y: rectangle.Position.Y + rectangle.Height, Z: rectangle.Position.Z},
			Color:    rectangle.Color,
		},
		{
			Drawable: rectangle.Drawable,
			VertexA:  Vector3{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y + rectangle.Height, Z: rectangle.Position.Z},
			VertexB:  Vector3{X: rectangle.Position.X + rectangle.Width, Y: rectangle.Position.Y, Z: rectangle.Position.Z},
			VertexC:  Vector3{X: rectangle.Position.X, Y: rectangle.Position.Y + rectangle.Height, Z: rectangle.Position.Z},
			Color:    rectangle.Color,
		},
	}
}

func (rectangle *Rectangle) Draw(screenBuffer []byte) {
	for _, triangle := range rectangle.Triangles() {
		triangle.Draw(screenBuffer)
	}
}
