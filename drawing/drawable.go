package drawing

import "math"

type Drawer interface {
	Draw([]byte)
}

type Drawable struct {
	// the translation along each axis in pixels
	Translation Vector2
	// the rotation in degrees
	Rotation float64
	// the scale factor along each axis (1 = no change, 2 = double, etc)
	Scale Vector2
}

func NewDrawable() Drawable {
	return Drawable{Scale: Vector2{X: 1, Y: 1}}
}

func (drawable *Drawable) ApplyTranslation(x, y int) (int, int) {
	return x + drawable.Translation.X, y + drawable.Translation.Y
}

func (drawable *Drawable) ApplyRotation(x, y int) (int, int) {
	return int(math.Cos(drawable.Rotation)*float64(x)) + int(math.Sin(drawable.Rotation)*float64(y)),
		int(math.Cos(drawable.Rotation)*float64(y)) + int(math.Sin(drawable.Rotation)*float64(-x))
}

func (drawable *Drawable) ApplyScale(x, y int) (int, int) {
	return x * drawable.Scale.X, y * drawable.Scale.Y
}

func (drawable *Drawable) ApplyTransformations(x, y int) (int, int) {
	return drawable.ApplyTranslation(drawable.ApplyScale(drawable.ApplyRotation(x, y)))
}
