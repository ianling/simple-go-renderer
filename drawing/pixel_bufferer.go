package drawing

type pixelBufferer interface {
	SetPixel(x, y int, z float64, r, g, b, a uint8)
}
