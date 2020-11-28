package ui

func endMenu(ui *UI2d) stateFunc {
	renderer.Copy(blackPixel, nil, nil)
	return determineToken
}
