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

func AveragePoints(vertices []Vector3) Vector3 {
	var x, y, z float64
	for _, vertex := range vertices {
		x += vertex.X
		y += vertex.Y
		z += vertex.Z
	}

	numVertices := float64(len(vertices))

	return Vector3{
		X: x / numVertices,
		Y: y / numVertices,
		Z: z / numVertices,
	}
}
