package drawing

type Vector3 struct {
	X int
	Y int
	Z int
}

func (vec *Vector3) MultiplyInt(n int) {
	vec.X = vec.X * n
	vec.Y = vec.Y * n
	vec.Z = vec.Z * n
}

func (vec *Vector3) MultiplyFloat(n float64) {
	vec.X = int(float64(vec.X) * n)
	vec.Y = int(float64(vec.Y) * n)
	vec.Z = int(float64(vec.Z) * n)
}
