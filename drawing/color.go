package drawing

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var (
	ColorRed   = Color{R: 255, G: 0, B: 0, A: 255}
	ColorGreen = Color{R: 0, G: 255, B: 0, A: 255}
	ColorBlue  = Color{R: 0, G: 0, B: 255, A: 255}
	ColorBlack = Color{R: 0, G: 0, B: 0, A: 255}
	ColorWhite = Color{R: 255, G: 255, B: 255, A: 255}
	ColorGrey  = Color{R: 127, G: 127, B: 127, A: 255}
)
