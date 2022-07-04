package drawing

func SetPixel(buffer []byte, x, y int, r, g, b, a uint8) {
	// 800 is the window width!
	pixelIndex := y*800*4 + x*4
	buffer[pixelIndex+0] = b
	buffer[pixelIndex+1] = g
	buffer[pixelIndex+2] = r
	buffer[pixelIndex+3] = a
}
