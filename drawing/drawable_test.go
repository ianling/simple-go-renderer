package drawing

import (
	"math"
	"testing"
)

const floatEquivalenceMargin = 1e-9

func aboutEqual(x, y float64) bool {
	return math.Abs(y-x) < floatEquivalenceMargin
}

func TestDrawable_ApplyRotation(t *testing.T) {
	d := NewDrawable(Vector3{}, Vector3{})

	// X-axis rotations
	d.Rotation.X = 90
	x, y, z := d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, -5) {
		t.Errorf("X-axis 90 degrees, expected 5,5,-5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.X = 180
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, -5) || !aboutEqual(z, -5) {
		t.Errorf("X-axis 180 degrees, expected 5,-5,-5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.X = 270
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, -5) || !aboutEqual(z, 5) {
		t.Errorf("X-axis 270 degrees, expected 5,-5,5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.X = 360
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("X-axis 360 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.X = 720
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("X-axis 720 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	d.ResetRotation()

	// Y-axis rotations
	d.Rotation.Y = 90
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, -5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("Y-axis 90 degrees, expected -5,5,5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Y = 180
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, -5) || !aboutEqual(y, 5) || !aboutEqual(z, -5) {
		t.Errorf("Y-axis 180 degrees, expected -5,5,-5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Y = 270
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, -5) {
		t.Errorf("Y-axis 270 degrees, expected 5,5,-5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Y = 360
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("Y-axis 360 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Y = 720
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("Y-axis 720 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	// Z-axis rotations
	d.Rotation.Z = 90
	x, y, z = d.ApplyRotation(5, 5, 0)
	if !aboutEqual(x, 5) || !aboutEqual(y, -5) || !aboutEqual(z, 0) {
		t.Errorf("Z-axis 90 degrees, expected 5,-5,0 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Z = 180
	x, y, z = d.ApplyRotation(5, 5, 0)
	if !aboutEqual(x, -5) || !aboutEqual(y, -5) || !aboutEqual(z, 0) {
		t.Errorf("Z-axis 180 degrees, expected -5,-5,0 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Z = 270
	x, y, z = d.ApplyRotation(5, 5, 0)
	if !aboutEqual(x, -5) || !aboutEqual(y, 5) || !aboutEqual(z, 0) {
		t.Errorf("Z-axis 270 degrees, expected -5,5,0 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Z = 360
	x, y, z = d.ApplyRotation(5, 5, 0)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 0) {
		t.Errorf("Z-axis 360 degrees, expected 5,5,0 got %f,%f,%f", x, y, z)
	}

	d.Rotation.Z = 720
	x, y, z = d.ApplyRotation(5, 5, 0)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 0) {
		t.Errorf("Z-axis 720 degrees, expected 5,5,0 got %f,%f,%f", x, y, z)
	}

	d.ResetRotation()

	// composite rotations
	// XY
	d.Rotation.X = 90
	d.Rotation.Y = 90
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("XY 90 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	d.ResetRotation()

	// XZ
	d.Rotation.X = 90
	d.Rotation.Z = 90
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, -5) || !aboutEqual(z, -5) {
		t.Errorf("XZ 90 degrees, expected 5,-5,-5 got %f,%f,%f", x, y, z)
	}

	d.ResetRotation()

	// YZ
	d.Rotation.Y = 90
	d.Rotation.Z = 90
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, 5) || !aboutEqual(z, 5) {
		t.Errorf("YZ 90 degrees, expected 5,5,5 got %f,%f,%f", x, y, z)
	}

	d.ResetRotation()

	// XYZ
	d.Rotation.X = 90
	d.Rotation.Y = 90
	d.Rotation.Z = 90
	x, y, z = d.ApplyRotation(5, 5, 5)
	if !aboutEqual(x, 5) || !aboutEqual(y, -5) || !aboutEqual(z, 5) {
		t.Errorf("XYZ 90 degrees, expected 5,-5,5 got %f,%f,%f", x, y, z)
	}

}
