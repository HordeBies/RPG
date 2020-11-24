package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var mainMenuBackground *sdl.Texture
var uiAtlas *sdl.Texture

type pos struct {
	x, y int
}

type button struct {
	pos
	srcRect   []*sdl.Rect
	dstRect   []*sdl.Rect
	str       *sdl.Texture
	isClicked bool
}

type mainMenuObj struct {
	play    button
	info    button
	infoTab *sdl.Texture
}

func createMainMenu(ui *UI2d) {
	mainMenuBackground = imgFileToTexture("ui/assets/main_menu_background.png")
	uiAtlas = imgFileToTexture("ui/assets/ui_split.png")
	ui.mainMenu = mainMenuObj{}

	ui.mainMenu.play = button{pos: pos{x: winWidth * 0.4, y: winHeight * .4}, isClicked: false}
	ui.mainMenu.play.srcRect = append(ui.mainMenu.play.srcRect, &sdl.Rect{310, 349, 25, 32})
	ui.mainMenu.play.srcRect = append(ui.mainMenu.play.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.mainMenu.play.dstRect = append(ui.mainMenu.play.dstRect, &sdl.Rect{winWidth * .4, winHeight * .65, 25 * 2, 32 * 2})
	ui.mainMenu.play.dstRect = append(ui.mainMenu.play.dstRect, &sdl.Rect{(winWidth * .4) + 24*2, (winHeight * .65) + 3*2, 70 * 2, 25 * 2})
	ui.mainMenu.play.str = getTextTexture("Play", sdl.Color{255, 255, 255, 0})
	ui.mainMenu.play.dstRect = append(ui.mainMenu.play.dstRect, &sdl.Rect{(winWidth * .4) + 28*2, (winHeight * .65) + 5*2, 45 * 2, 20 * 2})

	ui.mainMenu.info = button{pos: pos{x: winWidth * 0.4, y: winHeight * .5}, isClicked: false}
	ui.mainMenu.info.srcRect = append(ui.mainMenu.info.srcRect, &sdl.Rect{336, 349, 25, 32})
	ui.mainMenu.info.srcRect = append(ui.mainMenu.info.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.mainMenu.info.dstRect = append(ui.mainMenu.info.dstRect, &sdl.Rect{winWidth * .4, winHeight * .8, 25 * 2, 32 * 2})
	ui.mainMenu.info.dstRect = append(ui.mainMenu.info.dstRect, &sdl.Rect{(winWidth * .4) + 24*2, (winHeight * .8) + 3*2, 70 * 2, 25 * 2})
	ui.mainMenu.info.str = getTextTexture("Info", sdl.Color{255, 255, 255, 0})
	ui.mainMenu.info.dstRect = append(ui.mainMenu.info.dstRect, &sdl.Rect{(winWidth * .4) + 28*2, (winHeight * .8) + 5*2, 45 * 2, 20 * 2})

}

func drawMenuButtons(ui *UI2d) {
	for i := 0; i < 2; i++ {
		renderer.Copy(uiAtlas, ui.mainMenu.play.srcRect[i], ui.mainMenu.play.dstRect[i])
		renderer.Copy(uiAtlas, ui.mainMenu.info.srcRect[i], ui.mainMenu.info.dstRect[i])
	}

	renderer.Copy(ui.mainMenu.play.str, nil, ui.mainMenu.play.dstRect[2])
	renderer.Copy(ui.mainMenu.info.str, nil, ui.mainMenu.info.dstRect[2])
}

func createInfo(ui *UI2d) {

}

func updateMenu(ui *UI2d) {
	if !ui.mainMenu.play.isClicked && !ui.mainMenu.info.isClicked {
		drawMenuButtons(ui)
	} else if ui.mainMenu.play.isClicked {
		currentState = editScreen
		ui.mainMenu.play.isClicked = false
	} else if ui.mainMenu.info.isClicked {
		ui.mainMenu.info.isClicked = false
	}
}

func mainMenu(ui *UI2d) stateFunc {
	renderer.Copy(mainMenuBackground, nil, nil)
	updateMenu(ui)
	if ui.input.leftButton && !ui.input.prevLeftButton {
		fmt.Println("left clicked")
		for i := 0; i < 2; i++ {
			if ui.mainMenu.play.dstRect[i].HasIntersection(&sdl.Rect{int32(ui.input.x), int32(ui.input.y), 1, 1}) {
				ui.mainMenu.play.isClicked = true
			}
			if ui.mainMenu.info.dstRect[i].HasIntersection(&sdl.Rect{int32(ui.input.x), int32(ui.input.y), 1, 1}) {
				ui.mainMenu.info.isClicked = true
			}
		}
	}

	return determineToken
}
