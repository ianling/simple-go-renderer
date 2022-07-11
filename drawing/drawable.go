package drawing

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Drawer interface {
	Draw([]byte)
}

type Drawable struct {
	// the origin coordinates
	Origin Vector3
	// the translation along each axis in pixels
	Translation Vector3
	// the rotation in degrees around the Origin, about each axis
	Rotation Vector3
	// the scale factor along each axis (1 = no change, 2 = double, etc)
	Scale Vector3
}

func NewDrawable(origin Vector3) Drawable {
	return Drawable{Origin: origin, Scale: Vector3{X: 1, Y: 1, Z: 1}}
}

func (drawable *Drawable) ApplyTranslation(x, y, z float64) (float64, float64, float64) {
	return x + drawable.Translation.X, y + drawable.Translation.Y, z + drawable.Translation.Z
}

func (drawable *Drawable) ApplyRotation(x, y, z float64) (float64, float64, float64) {
	// convert angle from degrees -> radians, precompute some constants
	angleXRad := drawable.Rotation.X * (math.Pi / float64(180))
	angleYRad := drawable.Rotation.Y * (math.Pi / float64(180))
	angleZRad := drawable.Rotation.Z * (math.Pi / float64(180))
	cosX := float32(math.Cos(angleXRad))
	cosY := float32(math.Cos(angleYRad))
	cosZ := float32(math.Cos(angleZRad))
	sinX := float32(math.Sin(angleXRad))
	sinY := float32(math.Sin(angleYRad))
	sinZ := float32(math.Sin(angleZRad))

	// construct rotation matrices
	rotationX := mgl32.NewMatrixFromData([]float32{
		1, 0, 0,
		0, cosX, -sinX,
		0, sinX, cosX,
	}, 3, 3)
	rotationY := mgl32.NewMatrixFromData([]float32{
		cosY, 0, sinY,
		0, 1, 0,
		-sinY, 0, cosY,
	}, 3, 3)
	rotationZ := mgl32.NewMatrixFromData([]float32{
		cosZ, -sinZ, 0,
		sinZ, cosZ, 0,
		0, 0, 1,
	}, 3, 3)
	compositeRotation := mgl32.NewMatrix(3, 3)

	rotationZ.MulMxN(compositeRotation, rotationY)
	compositeRotation.MulMxN(compositeRotation, rotationX)

	// translate to origin
	x -= drawable.Origin.X
	y -= drawable.Origin.Y
	z -= drawable.Origin.Z

	coordinates := mgl32.NewMatrixFromData(
		[]float32{
			float32(x),
			float32(y),
			float32(z),
		}, 3, 1)

	compositeRotation.MulMxN(coordinates, coordinates)

	newX := float64(coordinates.At(0, 0))
	newY := float64(coordinates.At(1, 0))
	newZ := float64(coordinates.At(2, 0))

	// translate back from origin
	newX += drawable.Origin.X
	newY += drawable.Origin.Y
	newZ += drawable.Origin.Z

	return newX, newY, newZ
}

func (drawable *Drawable) ApplyScale(x, y, z float64) (float64, float64, float64) {
	return x * drawable.Scale.X, y * drawable.Scale.Y, z * drawable.Scale.Z
}

func (drawable *Drawable) ApplyTransformations(x, y, z float64) (float64, float64, float64) {
	return drawable.ApplyTranslation(drawable.ApplyScale(drawable.ApplyRotation(x, y, z)))
}

func (drawable *Drawable) ResetRotation() {
	drawable.Rotation = Vector3{}
}

func (drawable *Drawable) ResetScale() {
	drawable.Scale = Vector3{X: 1, Y: 1, Z: 1}
}

func (drawable *Drawable) ResetTranslation() {
	drawable.Translation = Vector3{}
}

func (drawable *Drawable) ResetTransformations() {
	drawable.ResetRotation()
	drawable.ResetScale()
	drawable.ResetTranslation()
}
