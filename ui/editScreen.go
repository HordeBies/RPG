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
	if ui.input.currKeyState[sdl.SCANCODE_4] != 0 && ui.input.prevKeyState[sdl.SCANCODE_4] == 0 {
		editingTile = game.MainCharacter
	}
}

func getTileLayer() int {
	switch editingTile {
	case game.DirtFloor:
		return 0
	case game.StoneWall:
		return 0
	case game.MainCharacter:
		return 1
	case game.Door:
		return 1
	default:
		panic("unknown tile in getTileLayer")
	}
}

var globalLevel *game.Level

func addToGridWorld(x, y, l int, tile game.Tile) {
	gridY := len(globalLevel.GridWorld.Rows)
	for gridY < y+1 {
		globalLevel.GridWorld.Rows = append(globalLevel.GridWorld.Rows, game.Row{})
		gridY++
	}
	gridX := len(globalLevel.GridWorld.Rows[y].Grids)
	for gridX < x+1 {
		globalLevel.GridWorld.Rows[y].Grids = append(globalLevel.GridWorld.Rows[y].Grids, game.Grid{})
		gridX++
	}
	gridL := len(globalLevel.GridWorld.Rows[y].Grids[x].Layers)
	if gridL == 0 || globalLevel.GridWorld.Rows[y].Grids[x].Layers == nil {
		globalLevel.GridWorld.Rows[y].Grids[x].Layers = make([]game.Tile, l+1)
	}
	for gridL < l+1 {
		globalLevel.GridWorld.Rows[y].Grids[x].Layers = append(globalLevel.GridWorld.Rows[y].Grids[x].Layers, game.Blank)
		gridL++
	}
	globalLevel.GridWorld.Rows[y].Grids[x].Layers[l] = tile
}

func editMenu(ui *UI2d) stateFunc {
	//fmt.Println("edit Menu")
	if ui.input.currKeyState[sdl.SCANCODE_S] == 0 && ui.input.prevKeyState[sdl.SCANCODE_S] != 0 {
		game.Save(globalLevel)
		fmt.Println("saving done")
	}
	if ui.input.currKeyState[sdl.SCANCODE_BACKSPACE] == 0 && ui.input.prevKeyState[sdl.SCANCODE_BACKSPACE] != 0 {
		createLayers(globalLevel, ui)
	}
	checkEditingTileChange(ui)

	if ui.input.leftButton && !ui.input.prevLeftButton {
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		l := getTileLayer()
		if ui.layers[l].dstRect[y][x] == nil {
			ui.layers[l].srcRect[y][x] = &textureIndex[editingTile][rand.Intn(len(textureIndex[editingTile]))]
			ui.layers[l].dstRect[y][x] = &sdl.Rect{X: int32(x) * 32, Y: int32(y) * 32, W: 32, H: 32}
			addToGridWorld(x, y, l, editingTile)
		}

	}
	if ui.input.rightButton && !ui.input.prevRightButton {
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		for i := len(ui.layers) - 1; i >= 0; i-- {
			if ui.layers[i].dstRect[y][x] != nil {
				ui.layers[i].dstRect[y][x] = nil
				ui.layers[i].srcRect[y][x] = nil
				globalLevel.GridWorld.Rows[y].Grids[x].Layers[i] = game.Blank
				break
			}
		}
	}
	return determineToken
}
