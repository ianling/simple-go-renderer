package drawing

type Pyramid struct {
	Drawable
	Position Vector3
	Color    Color
	Width    float64
	Height   float64
	Depth    float64
}

func NewPyramid(x, y, z, width, height, depth float64, color Color) Pyramid {
	originX := x + (width / 2)
	originY := y + (height / 2)
	originZ := z + (depth / 2)

	return Pyramid{
		Drawable: NewDrawable(Vector3{X: originX, Y: originY, Z: originZ}),
		Position: Vector3{X: x, Y: y, Z: z},
		Color:    color,
		Width:    width,
		Height:   height,
		Depth:    depth,
	}
}

func (pyramid *Pyramid) Faces() []Polygon {
	bottomLeftVertex := Vector3{
		X: pyramid.Position.X,
		Y: pyramid.Position.Y,
		Z: pyramid.Position.Z,
	}
	bottomRightVertex := Vector3{
		X: pyramid.Position.X + pyramid.Width,
		Y: pyramid.Position.Y,
		Z: pyramid.Position.Z,
	}
	backVertex := Vector3{
		X: pyramid.Position.X + pyramid.Width/2,
		Y: pyramid.Position.Y,
		Z: pyramid.Position.Z + pyramid.Depth,
	}
	topVertex := Vector3{
		X: pyramid.Position.X + pyramid.Width/2,
		Y: pyramid.Position.Y + pyramid.Height,
		Z: pyramid.Position.Z + pyramid.Depth/2,
	}

	frontVertices := []Vector3{bottomLeftVertex, bottomRightVertex, topVertex}
	leftVertices := []Vector3{bottomLeftVertex, backVertex, topVertex}
	rightVertices := []Vector3{bottomRightVertex, backVertex, topVertex}
	bottomVertices := []Vector3{bottomLeftVertex, bottomRightVertex, backVertex}

	frontFace := NewPolygon(frontVertices, pyramid.Drawable, pyramid.Color)
	leftFace := NewPolygon(leftVertices, pyramid.Drawable, pyramid.Color)
	rightFace := NewPolygon(rightVertices, pyramid.Drawable, pyramid.Color)
	bottomFace := NewPolygon(bottomVertices, pyramid.Drawable, pyramid.Color)

	return []Polygon{
		frontFace,
		leftFace,
		rightFace,
		bottomFace,
	}
}

func (pyramid *Pyramid) Draw(pixelBuffer []byte) {
	for _, face := range pyramid.Faces() {
		face.Draw(pixelBuffer)
	}
}
