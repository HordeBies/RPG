package ui

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/veandco/go-sdl2/sdl"
)

var editingTile game.Tile = game.DirtFloor

func checkEditingTileChange(ui *UI2d) {
	if ui.input.currKeyState[sdl.SCANCODE_1] != 0 && ui.input.prevKeyState[sdl.SCANCODE_1] == 0 {
		editingTile = game.DirtFloor
	}
	if ui.input.currKeyState[sdl.SCANCODE_2] != 0 && ui.input.prevKeyState[sdl.SCANCODE_2] == 0 {
		editingTile = game.StoneWall
	}
	if ui.input.currKeyState[sdl.SCANCODE_3] != 0 && ui.input.prevKeyState[sdl.SCANCODE_3] == 0 {
		editingTile = game.Door
	}
}

func editMenu(ui *UI2d) stateFunc {
	//fmt.Println("edit Menu")
	if ui.input.currKeyState[sdl.SCANCODE_S] != 0 && ui.input.prevKeyState[sdl.SCANCODE_S] == 0 {
		fmt.Println("pressed s in edit menu")
	}
	checkEditingTileChange(ui)

	if ui.input.leftButton && !ui.input.prevLeftButton {
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		for l := range ui.layers {
			if ui.layers[l].dstRect[y][x] == nil {
				ui.layers[l].srcRect[y][x] = &textureIndex[editingTile][rand.Intn(len(textureIndex[editingTile]))]
				ui.layers[l].dstRect[y][x] = &sdl.Rect{X: int32(x) * 32, Y: int32(y) * 32, W: 32, H: 32}
				break
			}
		}
	}
	if ui.input.rightButton && !ui.input.prevRightButton {
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		for i := len(ui.layers) - 1; i >= 0; i-- {
			if ui.layers[i].dstRect[y][x] != nil {
				ui.layers[i].dstRect[y][x] = nil
				ui.layers[i].srcRect[y][x] = nil
				break
			}
		}
	}
	return determineToken
}
