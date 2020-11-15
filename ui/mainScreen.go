package ui

import "fmt"

func mainMenu(ui *UI2d) stateFunc {
	for currentState == mainScreen {
		fmt.Println("mainMenu")
		renderer.Present()
	}
	return determineToken
}
