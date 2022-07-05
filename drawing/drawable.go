package drawing

import "math"

type Drawer interface {
	Draw([]byte)
}

type Drawable struct {
	// the origin coordinates
	Origin Vector2
	// the translation along each axis in pixels
	Translation Vector2
	// the rotation about the Origin in degrees
	Rotation float64
	// the scale factor along each axis (1 = no change, 2 = double, etc)
	Scale Vector2
}

func NewDrawable(origin Vector2) Drawable {
	return Drawable{Origin: origin, Scale: Vector2{X: 1, Y: 1}}
}

func (drawable *Drawable) ApplyTranslation(x, y int) (int, int) {
	return x + drawable.Translation.X, y + drawable.Translation.Y
}

func (drawable *Drawable) ApplyRotation(x, y int) (int, int) {
	// convert angle from degrees -> radians
	angleRad := drawable.Rotation * (math.Pi / float64(180))

	x -= drawable.Origin.X
	y -= drawable.Origin.Y

	newX := int(math.Cos(angleRad)*float64(x)) + int(math.Sin(angleRad)*float64(y))
	newY := -int(math.Sin(angleRad)*float64(x)) + int(math.Cos(angleRad)*float64(y))

	newX += drawable.Origin.X
	newY += drawable.Origin.Y

	return newX, newY
}

func (drawable *Drawable) ApplyScale(x, y int) (int, int) {
	return x * drawable.Scale.X, y * drawable.Scale.Y
}

func (drawable *Drawable) ApplyTransformations(x, y int) (int, int) {
	return drawable.ApplyTranslation(drawable.ApplyScale(drawable.ApplyRotation(x, y)))
}
