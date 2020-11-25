package ui

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"

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
		editingTile = game.DoorC
	}
	if ui.input.currKeyState[sdl.SCANCODE_4] != 0 && ui.input.prevKeyState[sdl.SCANCODE_4] == 0 {
		editingTile = game.MainCharacter
	}
}

func getTileType() int {
	switch editingTile {
	case game.DirtFloor:
		return 0
	case game.StoneWall:
		return 0
	case game.MainCharacter:
		return 1
	case game.DoorC:
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
		globalLevel.GridWorld.Rows[y].Grids = append(globalLevel.GridWorld.Rows[y].Grids, game.Grid{Layers: []game.Tile{}})
		gridX++
	}
	//gridL := len(globalLevel.GridWorld.Rows[y].Grids[x].Layers)
	/*	if gridL == 0 || globalLevel.GridWorld.Rows[y].Grids[x].Layers == nil {
		globalLevel.GridWorld.Rows[y].Grids[x].Layers = make([]game.Tile, l+1)
	}*/
	if tile == game.Blank || tile == game.DirtFloor || tile == game.StoneWall {
		globalLevel.GridWorld.Rows[y].Grids[x].Background = tile
	}
}

func editMenu(ui *UI2d) stateFunc {
	//fmt.Println("edit Menu")
	if ui.input.currKeyState[sdl.SCANCODE_S] == 0 && ui.input.prevKeyState[sdl.SCANCODE_S] != 0 {
		game.Save(globalLevel)
		fmt.Println("saving done")
	}
	checkEditingTileChange(ui)

	if ui.input.leftButton { // && !ui.input.prevLeftButton
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		l := getTileType()
		if ui.background.dstRect[y][x] == nil {
			ui.background.srcRect[y][x] = &textureIndex[editingTile][rand.Intn(len(textureIndex[editingTile]))]
			ui.background.dstRect[y][x] = &sdl.Rect{X: int32(x) * 32, Y: int32(y) * 32, W: 32, H: 32}
			addToGridWorld(x, y, l, editingTile)
		}

	}
	if ui.input.rightButton { //&& !ui.input.prevRightButton
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		if ui.background.dstRect[y][x] != nil {
			ui.background.dstRect[y][x] = nil
			ui.background.srcRect[y][x] = nil
			globalLevel.GridWorld.Rows[y].Grids[x].Background = game.Blank
		}
	}

	renderer.Copy(mainMenuBackground, nil, nil)
	//renderer.Copy(blackPixel, nil, &sdl.Rect{0, 0, winWidth, winHeight})

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			if ui.background.dstRect[y][x] != nil {
				renderer.Copy(textureAtlas, ui.background.srcRect[y][x], ui.background.dstRect[y][x])
			}
		}
	}

	return determineToken
}

var buttonSelected *sdl.Texture

type levelButton struct {
	levelName string
	isClicked bool
	texture   *sdl.Texture
	rect      *sdl.Rect
}

type selectMenuObj struct {
	levels []levelButton
	start  button
}

func createSelectMenu(ui *UI2d) {
	ui.selectMenu = selectMenuObj{}
	ui.selectMenu.start = button{pos: pos{x: winWidth * 0.4, y: winHeight * .4}, isClicked: false}
	ui.selectMenu.start.srcRect = append(ui.selectMenu.start.srcRect, &sdl.Rect{310, 349, 25, 32})
	ui.selectMenu.start.srcRect = append(ui.selectMenu.start.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{winWidth - 200, winHeight - 70, 25 * 2, 32 * 2})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{(winWidth - 200) + 24*2, (winHeight - 70) + 3*2, 70 * 2, 25 * 2})
	ui.selectMenu.start.str = getTextTexture("Start", sdl.Color{255, 255, 255, 0})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{(winWidth - 200) + 28*2, (winHeight - 70) + 5*2, 45 * 2, 20 * 2})

	files, err := ioutil.ReadDir("./game/maps/")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		name := f.Name()
		if strings.Contains(name, "level") && !strings.Contains(name, "Entities") && strings.HasSuffix(name, ".map") {
			ui.selectMenu.levels = append(ui.selectMenu.levels, levelButton{strings.TrimSuffix(name, ".map"), false, nil, nil})
		}
	}
	for i := range ui.selectMenu.levels {
		ui.selectMenu.levels[i].texture = getTextTexture(ui.selectMenu.levels[i].levelName, sdl.Color{255, 255, 255, 0})
		ui.selectMenu.levels[i].rect = &sdl.Rect{5, int32(40 * i), 90, 40}
	}
}

func updateSelections(ui *UI2d) {
	if ui.input.leftButton && !ui.input.prevLeftButton {
		clickRect := &sdl.Rect{int32(ui.input.x), int32(ui.input.y), 1, 1}
		if ui.selectMenu.start.dstRect[0].HasIntersection(clickRect) || ui.selectMenu.start.dstRect[1].HasIntersection(clickRect) {
			for _, level := range ui.selectMenu.levels {
				if level.isClicked {
					ui.input.updateMouseState()
					globalLevel = &game.Level{}
					globalLevel.LevelName = level.levelName
					break
				}
			}
		}
		for i, level := range ui.selectMenu.levels {
			if level.rect.HasIntersection(clickRect) {
				for i := range ui.selectMenu.levels {
					ui.selectMenu.levels[i].isClicked = false
				}
				ui.selectMenu.levels[i].isClicked = true
				break
			}
		}
	}
}

func selectMenu(ui *UI2d) stateFunc {
	renderer.Copy(mainMenuBackground, nil, nil)
	renderer.Copy(ui.mainMenu.infoTab, nil, &sdl.Rect{0, 0, 150, winHeight})
	updateSelections(ui)

	for i, level := range ui.selectMenu.levels {
		if level.isClicked {
			px := createOnePixel(255, 255, 255, 200)
			renderer.Copy(px, nil, &sdl.Rect{0, int32(40 * i), 110, 40})
			renderer.Copy(ui.mainMenu.infoTab, nil, &sdl.Rect{5, int32(i*40) + 5, 100, 30})
			for y := 0; y < 100; y++ {
				for x := 0; x < 100; x++ {
					if ui.levelPreviews[i].dstRect[y][x] != nil {
						renderer.Copy(textureAtlas, ui.levelPreviews[i].srcRect[y][x], ui.levelPreviews[i].dstRect[y][x])
					}
				}
			}
		}
		renderer.Copy(level.texture, nil, level.rect)
	}

	for i := 0; i < 2; i++ {
		renderer.Copy(uiAtlas, ui.selectMenu.start.srcRect[i], ui.selectMenu.start.dstRect[i])
	}
	renderer.Copy(ui.selectMenu.start.str, nil, ui.selectMenu.start.dstRect[2])

	if ui.input.currKeyState[sdl.SCANCODE_ESCAPE] != 0 && ui.input.prevKeyState[sdl.SCANCODE_ESCAPE] == 0 {
		currentState = mainScreen
	}

	return determineToken
}
