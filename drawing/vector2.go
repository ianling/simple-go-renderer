package drawing

type Vector2 struct {
	X int
	Y int
}

func (vec *Vector2) MultiplyInt(n int) {
	vec.X = vec.X * n
	vec.Y = vec.Y * n
}

func (vec *Vector2) MultiplyFloat(n float64) {
	vec.X = int(float64(vec.X) * n)
	vec.Y = int(float64(vec.Y) * n)
}
