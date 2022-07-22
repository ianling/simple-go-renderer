package main

var emptyBuffer = make([]byte, windowWidth*windowHeight*4)

func init() {
	for ii := 0; ii < len(emptyBuffer); ii += 4 {
		emptyBuffer[ii] = 255
		emptyBuffer[ii+1] = 255
		emptyBuffer[ii+2] = 255
		emptyBuffer[ii+3] = 255
	}
}

func constrainWithinWindow(x, y int) (int, int) {
	if x < 0 {
		x = 0
	} else if x > windowWidth {
		x = windowWidth
	}

	if y < 0 {
		y = 0
	} else if y > windowHeight {
		y = windowHeight
	}

	return x, y
}

func clearBuffer(buffer []byte) {
	copy(buffer, emptyBuffer)
}
