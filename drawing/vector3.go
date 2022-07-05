package drawing

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (vec *Vector3) MultiplyInt(n int) {
	vec.X = vec.X * float64(n)
	vec.Y = vec.Y * float64(n)
	vec.Z = vec.Z * float64(n)
}

func (vec *Vector3) MultiplyFloat(n float64) {
	vec.X = vec.X * n
	vec.Y = vec.Y * n
	vec.Z = vec.Z * n
}
