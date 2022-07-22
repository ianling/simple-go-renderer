package drawing

type Rectangle struct {
	Polygon
}

func NewRectangle(x, y, z float64, width, height float64, color Color) Rectangle {
	vertices := []Vector3{
		{
			X: x,
			Y: y,
			Z: z,
		},
		{
			X: x + width,
			Y: y,
			Z: z,
		},
		{
			X: x,
			Y: y + height,
			Z: z,
		},
		{
			X: x + width,
			Y: y + height,
			Z: z,
		},
	}

	return Rectangle{
		Polygon: NewPolygon(vertices, NewDrawable(AveragePoints(vertices)), color),
	}
}
