package main

var emptyBuffer = make([]byte, windowWidth*windowHeight*4)

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
