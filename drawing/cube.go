package drawing

type Cube struct {
	Drawable
	Color    Color
	Position Vector3
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
		Color:    color,
		Position: Vector3{X: x, Y: y, Z: z},
		Width:    width,
		Height:   height,
		Depth:    depth,
	}
}

func (cube *Cube) Rectangles() []Rectangle {
	// a cube is made up of 6 rectangular faces, 5 of which are just the front face rotated in some way.
	// the default positions are given assuming the cube is being viewed front-on (so only the front face is visible),
	// without regard for rotation or perspective
	frontFace := Rectangle{
		Drawable: cube.Drawable,
		Position: cube.Position,
		Width:    cube.Width,
		Height:   cube.Height,
		Color:    cube.Color,
	}
	leftFace := frontFace
	rightFace := frontFace
	topFace := frontFace
	bottomFace := frontFace
	backFace := frontFace

	// from front-on, the left face is rotated 90 degrees (counter-clockwise) about the Y-axis,
	leftFace.Rotation.Y += 90

	// right face is rotated 270 degrees about the Y-axis
	rightFace.Rotation.Y += 270

	// top face is rotated 90 degrees about the X-axis
	topFace.Rotation.X += 90

	// bottom face is rotated 270 degrees about the X-axis
	bottomFace.Rotation.X += 270

	//// back face is rotated 180 degrees about the X-axis
	backFace.Rotation.X += 180

	// debug: different colors for each face
	// left green
	leftFace.Color.R = 0
	leftFace.Color.G = 255
	// right blue
	rightFace.Color.R = 0
	rightFace.Color.B = 255
	// top red-green
	topFace.Color.G = 200
	// bottom red-blue
	bottomFace.Color.B = 200
	// back white
	backFace.Color.G = 255
	backFace.Color.B = 255

	return []Rectangle{
		frontFace,
		leftFace,
		rightFace,
		backFace,
		topFace,
		bottomFace,
	}
}

func (cube *Cube) Draw(pixelBuffer []byte) {
	for _, rectangle := range cube.Rectangles() {
		rectangle.Draw(pixelBuffer)
	}
}
