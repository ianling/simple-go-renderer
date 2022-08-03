package drawing

type Cube struct {
	Drawable
	Position Vector3
	Color    Color
	Width    float64
	Height   float64
	Depth    float64
}

func NewCube(x, y, z, width, height, depth float64, color Color) Cube {
	originX := x + (width / 2)
	originY := y + (height / 2)
	originZ := z + (depth / 2)
	return Cube{
		Drawable: NewDrawable(Vector3{X: originX, Y: originY, Z: originZ}),
		Position: Vector3{X: x, Y: y, Z: z},
		Color:    color,
		Width:    width,
		Height:   height,
		Depth:    depth,
	}
}

func (cube *Cube) Faces() []Polygon {
	leftBottomFrontVertex := Vector3{
		X: cube.Position.X,
		Y: cube.Position.Y,
		Z: cube.Position.Z,
	}
	rightBottomFrontVertex := Vector3{
		X: cube.Position.X + cube.Width,
		Y: cube.Position.Y,
		Z: cube.Position.Z,
	}
	leftBottomBackVertex := Vector3{
		X: cube.Position.X,
		Y: cube.Position.Y,
		Z: cube.Position.Z + cube.Depth,
	}
	rightBottomBackVertex := Vector3{
		X: cube.Position.X + cube.Width,
		Y: cube.Position.Y,
		Z: cube.Position.Z + cube.Depth,
	}
	leftTopFrontVertex := Vector3{
		X: cube.Position.X,
		Y: cube.Position.Y + cube.Height,
		Z: cube.Position.Z,
	}
	rightTopFrontVertex := Vector3{
		X: cube.Position.X + cube.Width,
		Y: cube.Position.Y + cube.Height,
		Z: cube.Position.Z,
	}
	leftTopBackVertex := Vector3{
		X: cube.Position.X,
		Y: cube.Position.Y + cube.Height,
		Z: cube.Position.Z + cube.Depth,
	}
	rightTopBackVertex := Vector3{
		X: cube.Position.X + cube.Width,
		Y: cube.Position.Y + cube.Height,
		Z: cube.Position.Z + cube.Depth,
	}

	frontVertices := []Vector3{leftBottomFrontVertex, rightBottomFrontVertex, rightTopFrontVertex, leftTopFrontVertex}
	leftVertices := []Vector3{leftBottomFrontVertex, leftTopFrontVertex, leftTopBackVertex, leftBottomBackVertex}
	rightVertices := []Vector3{rightBottomFrontVertex, rightTopFrontVertex, rightTopBackVertex, rightBottomBackVertex}
	topVertices := []Vector3{leftTopFrontVertex, rightTopFrontVertex, rightTopBackVertex, leftTopBackVertex}
	bottomVertices := []Vector3{leftBottomFrontVertex, rightBottomFrontVertex, rightBottomBackVertex, leftBottomBackVertex}
	backVertices := []Vector3{leftBottomBackVertex, rightBottomBackVertex, rightTopBackVertex, leftTopBackVertex}

	frontFace := NewPolygon(frontVertices, cube.Drawable, cube.Color)
	leftFace := NewPolygon(leftVertices, cube.Drawable, cube.Color)
	rightFace := NewPolygon(rightVertices, cube.Drawable, cube.Color)
	topFace := NewPolygon(topVertices, cube.Drawable, cube.Color)
	bottomFace := NewPolygon(bottomVertices, cube.Drawable, cube.Color)
	backFace := NewPolygon(backVertices, cube.Drawable, cube.Color)

	return []Polygon{
		frontFace,
		leftFace,
		rightFace,
		backFace,
		topFace,
		bottomFace,
	}
}

func (cube *Cube) Draw(window pixelBufferer) {
	for _, rectangle := range cube.Faces() {
		rectangle.Draw(window)
	}
}
