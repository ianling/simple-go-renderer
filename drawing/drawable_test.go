package drawing

import "testing"

func TestDrawable_ApplyRotation(t *testing.T) {
	d := NewDrawable(Vector2{X: 0, Y: 0})

	d.Rotation = 90
	x, y := d.ApplyRotation(5, 5)
	if x != 5 || y != -5 {
		t.Errorf("incorrect rotation of 90 degrees, expected 5,-5 got %d,%d", x, y)
	}

	d.Rotation = 180
	x, y = d.ApplyRotation(5, 5)
	if x != -5 || y != -5 {
		t.Errorf("incorrect rotation of 180 degrees, expected -5,-5 got %d,%d", x, y)
	}

	d.Rotation = 270
	x, y = d.ApplyRotation(5, 5)
	if x != -5 || y != 5 {
		t.Errorf("incorrect rotation of 270 degrees, expected -5,5 got %d,%d", x, y)
	}

	d.Rotation = 360
	x, y = d.ApplyRotation(5, 5)
	if x != 5 || y != 5 {
		t.Errorf("incorrect rotation of 360 degrees, expected 5,5 got %d,%d", x, y)
	}

	d.Rotation = 720
	x, y = d.ApplyRotation(5, 5)
	if x != 5 || y != 5 {
		t.Errorf("incorrect rotation of 720 degrees, expected 5,5 got %d,%d", x, y)
	}
}
