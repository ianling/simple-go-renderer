package main

// Window holds all the pixel values and does some other rendering things, independent of the actual renderer used.
type Window struct {
	width       int
	height      int
	buffer      []byte
	emptyBuffer []byte
	zCache      [][]float64
}

func NewWindow(width, height int) *Window {
	window := Window{
		width:       width,
		height:      height,
		buffer:      make([]byte, width*height*4),
		emptyBuffer: make([]byte, width*height*4),
	}

	window.zCache = make([][]float64, height)
	for ii := range window.zCache {
		window.zCache[ii] = make([]float64, width)
	}

	for ii := 0; ii < len(window.emptyBuffer); ii += 4 {
		window.emptyBuffer[ii] = 255
		window.emptyBuffer[ii+1] = 255
		window.emptyBuffer[ii+2] = 255
		window.emptyBuffer[ii+3] = 255
	}

	window.ClearBuffer()

	return &window
}

func (window *Window) ClearBuffer() {
	copy(window.buffer, window.emptyBuffer)

	for ii := range window.zCache {
		for jj := range window.zCache[ii] {
			window.zCache[ii][jj] = -99999999999 // arbitrary small value
		}
	}
}

func (window *Window) SetPixel(x, y int, z float64, r, g, b, a uint8) {
	if window.zCache[y][x] >= z {
		return
	}

	window.zCache[y][x] = z

	pixelIndex := y*window.width*4 + x*4

	if pixelIndex < 0 || pixelIndex+3 > len(window.buffer) {
		return
	}

	window.buffer[pixelIndex+0] = b
	window.buffer[pixelIndex+1] = g
	window.buffer[pixelIndex+2] = r
	window.buffer[pixelIndex+3] = a
}

func (window *Window) CopyPixelBufferTo(destination []byte) {
	copy(destination, window.buffer)
}
