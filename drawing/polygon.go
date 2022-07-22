package drawing

type Polygon struct {
	Drawable
	Vertices []Vector3
	Color    Color
}

func NewPolygon(vertices []Vector3, drawable Drawable, color Color) Polygon {
	return Polygon{
		Drawable: drawable,
		Color:    color,
		Vertices: vertices,
	}
}

func (polygon *Polygon) Lines() LineSegments {
	numVertices := len(polygon.Vertices)
	lineSegments := make(LineSegments, 0, numVertices)

	for ii := 0; ii < numVertices; ii++ {
		lineSegments = append(lineSegments, NewLineSegment(polygon.Vertices[ii], polygon.Vertices[(ii+1)%numVertices], polygon.Drawable, polygon.Color))
	}

	return lineSegments
}

func (polygon *Polygon) Draw(screenBuffer []byte) {
	polygon.Lines().Draw(screenBuffer)
}
