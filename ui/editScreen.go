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
		editingTile = editingTileSlice[0]
	}
	if ui.input.currKeyState[sdl.SCANCODE_2] != 0 && ui.input.prevKeyState[sdl.SCANCODE_2] == 0 {
		editingTile = editingTileSlice[1]
	}
	if ui.input.currKeyState[sdl.SCANCODE_3] != 0 && ui.input.prevKeyState[sdl.SCANCODE_3] == 0 {
		editingTile = editingTileSlice[2]
	}
	if ui.input.currKeyState[sdl.SCANCODE_4] != 0 && ui.input.prevKeyState[sdl.SCANCODE_4] == 0 {
		editingTile = editingTileSlice[3]
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

var editingTileSlice = []game.Tile{
	game.DirtFloor,
	game.StoneWall,
	game.DoorC,
	game.MainCharacter,
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

func editTile(ui *UI2d) {
	if ui.input.leftButton { // && !ui.input.prevLeftButton
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		l := getTileType()
		if ui.background.dstRect[y][x] == nil && l == 0 {
			ui.background.srcRect[y][x] = &textureIndex[editingTile][rand.Intn(len(textureIndex[editingTile]))]
			ui.background.dstRect[y][x] = &sdl.Rect{X: int32(x) * 32, Y: int32(y) * 32, W: 32, H: 32}
			addToGridWorld(x, y, l, editingTile)
		} else if l == 1 && !ui.input.prevLeftButton {
			globalLevel.Entities = append(globalLevel.Entities, game.Entity{x * 32, y * 32, editingTile})
			ui.background.entities = append(ui.background.entities, getEntity(globalLevel.Entities[len(globalLevel.Entities)-1]))
		}

	}
	if ui.input.rightButton { //&&
		//isDeleted := false
		x := int(math.Floor(float64(ui.input.x) / 32))
		y := int(math.Floor(float64(ui.input.y) / 32))
		if len(ui.background.entities) > 0 && !ui.input.prevRightButton {
			for i, intf := range ui.background.entities {
				obj := intf.(entityInterface)
				if ui.input.x < obj.getX()+32 && ui.input.x >= obj.getX() && ui.input.y < obj.getY()+32 && ui.input.y >= obj.getY() {
					if len(ui.background.entities) > 1 {
						ui.background.entities = append(ui.background.entities[0:i], ui.background.entities[i+1:len(ui.background.entities)]...)
						globalLevel.Entities = append(globalLevel.Entities[0:i], globalLevel.Entities[i+1:len(globalLevel.Entities)]...)
					} else {
						ui.background.entities = ui.background.entities[0:0]
						globalLevel.Entities = globalLevel.Entities[0:0]
					}
					//isDeleted = true
				}
			}
		}
		if ui.background.dstRect[y][x] != nil { //!isDeleted &&
			ui.background.dstRect[y][x] = nil
			ui.background.srcRect[y][x] = nil
			globalLevel.GridWorld.Rows[y].Grids[x].Background = game.Blank
		}
	}

	if ui.input.currKeyState[sdl.SCANCODE_BACKSPACE] != 0 && ui.input.prevKeyState[sdl.SCANCODE_BACKSPACE] == 0 {
		fmt.Println("Level Reloaded")
		globalLevel = globalLevel.ReLoadTheLevel()
		createLayers(globalLevel, ui)

	}
}

func currTileChangeMenu(ui *UI2d) {
	var x int32 = 300
	var y int32 = 200
	var w int = 240
	var h int = 152 + 48
	var tileTabDst []*sdl.Rect
	renderer.Copy(ui.mainMenu.infoTab, &sdl.Rect{0, 0, 1, 1}, &sdl.Rect{x, y, int32(w), int32(h)})
	for i, tile := range editingTileSlice {
		if 76*(i+1) > w {
			y = y + 76
		}
		if editingTile == tile {
			px := createOnePixel(255, 255, 255, 200)
			renderer.Copy(px, nil, &sdl.Rect{x + 6 + (int32(76*i))%int32(w-12), y, 76, 76})
			renderer.Copy(ui.mainMenu.infoTab, &sdl.Rect{0, 0, 1, 1}, &sdl.Rect{x + 12 + (int32(76*i))%int32(w-12), y + 6, 64, 64})
		}
		tileTabDst = append(tileTabDst, &sdl.Rect{x + 12 + (int32(76*i))%int32(w-12), y + 6, 64, 64})
		renderer.Copy(textureAtlas, &textureIndex[tile][0], tileTabDst[i])
	}

	x = 360
	y += 76
	renderer.Copy(uiAtlas, ui.selectMenu.start.srcRect[0], &sdl.Rect{x, y, 25*1.5 - 0.5, 32 * 1.5})
	renderer.Copy(uiAtlas, ui.selectMenu.start.srcRect[1], &sdl.Rect{x + 24*1.5, y + 4, 55*1.5 - .5, 25 * 1.6})
	renderer.Copy(ui.selectMenu.start.str, nil, &sdl.Rect{x + 28*1.5, y + 7, 45*1.5 - 0.5, 20 * 1.5})

	if ui.input.leftButton && !ui.input.prevLeftButton {
		clickRect := &sdl.Rect{int32(ui.input.x), int32(ui.input.y), 1, 1}
		for i, rect := range tileTabDst {
			if clickRect.HasIntersection(rect) {
				editingTile = editingTileSlice[i]
				break
			}
		}
		if clickRect.HasIntersection(&sdl.Rect{x, y, 25*1.5 - 0.5, 32 * 1.5}) || clickRect.HasIntersection(&sdl.Rect{x + 24*1.5, y + 4, 55*1.5 - .5, 25 * 1.6}) {
			game.Save(globalLevel)
			currentState = endScreen
		}
	}
}

func editMenu(ui *UI2d) stateFunc {
	//fmt.Println("edit Menu")
	if ui.input.currKeyState[sdl.SCANCODE_S] == 0 && ui.input.prevKeyState[sdl.SCANCODE_S] != 0 {
		game.Save(globalLevel)
		fmt.Println("saving done")
	}
	checkEditingTileChange(ui)

	renderer.Copy(mainMenuBackground, nil, nil)
	//renderer.Copy(blackPixel, nil, &sdl.Rect{0, 0, winWidth, winHeight})

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			if ui.background.dstRect[y][x] != nil {
				renderer.Copy(textureAtlas, ui.background.srcRect[y][x], ui.background.dstRect[y][x])
			}
		}
	}
	for _, intf := range ui.background.entities {
		obj := intf.(entityInterface)
		renderer.Copy(textureAtlas, obj.getRect(), &sdl.Rect{int32(obj.getX()), int32(obj.getY()), 32, 32})
	}

	if ui.input.currKeyState[sdl.SCANCODE_TAB] != 0 || ui.input.currKeyState[sdl.SCANCODE_LSHIFT] != 0 {
		currTileChangeMenu(ui)
	} else {
		editTile(ui)
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
	edit   button
	rand   button
}

func createSelectMenu(ui *UI2d) {
	ui.selectMenu = selectMenuObj{}
	ui.selectMenu.start = button{pos: pos{x: winWidth * .4, y: winHeight * .4}, isClicked: false}
	ui.selectMenu.start.srcRect = append(ui.selectMenu.start.srcRect, &sdl.Rect{310, 349, 25, 32})
	ui.selectMenu.start.srcRect = append(ui.selectMenu.start.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{0, 480 - 50, 25*1.5 - 0.5, 32 * 1.5})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{0 + 24*1.5, 480 + 3*1.5 - 50 - 0.5, 55*1.5 - .5, 25 * 1.6})
	ui.selectMenu.start.str = getTextTexture("Start", sdl.Color{255, 255, 255, 0})
	ui.selectMenu.start.dstRect = append(ui.selectMenu.start.dstRect, &sdl.Rect{0 + 28*1.5, 480 + 5*1.5 - .5 - 50, 45*1.5 - 0.5, 20 * 1.5})

	ui.selectMenu.edit = button{pos: pos{x: winWidth * 0.4, y: winHeight * .4}, isClicked: false}
	ui.selectMenu.edit.srcRect = append(ui.selectMenu.edit.srcRect, &sdl.Rect{336, 349, 25, 32})
	ui.selectMenu.edit.srcRect = append(ui.selectMenu.edit.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.selectMenu.edit.dstRect = append(ui.selectMenu.edit.dstRect, &sdl.Rect{0, 480, 25*1.5 - 0.5, 32 * 1.5})
	ui.selectMenu.edit.dstRect = append(ui.selectMenu.edit.dstRect, &sdl.Rect{0 + 24*1.5, 480 + 3*1.5 - 0.5, 46 * 1.5, 25 * 1.6})
	ui.selectMenu.edit.str = getTextTexture("Edit", sdl.Color{255, 255, 255, 0})
	ui.selectMenu.edit.dstRect = append(ui.selectMenu.edit.dstRect, &sdl.Rect{0 + 28*1.5, 480 + 5*1.5 - .5, 36 * 1.5, 20 * 1.5})

	ui.selectMenu.rand = button{pos: pos{x: winWidth * .4, y: winHeight * .4}, isClicked: false}
	ui.selectMenu.rand.srcRect = append(ui.selectMenu.rand.srcRect, &sdl.Rect{362, 349, 25, 32})
	ui.selectMenu.rand.srcRect = append(ui.selectMenu.rand.srcRect, &sdl.Rect{313, 381, 70, 25})
	ui.selectMenu.rand.dstRect = append(ui.selectMenu.rand.dstRect, &sdl.Rect{0, 480 + 50, 25*1.5 - 0.5, 32 * 1.5})
	ui.selectMenu.rand.dstRect = append(ui.selectMenu.rand.dstRect, &sdl.Rect{0 + 24*1.5, 480 + 3*1.5 + 50 - 0.5, 82 * 1.5, 25 * 1.6})
	ui.selectMenu.rand.str = getTextTexture("ReCreate", sdl.Color{255, 255, 255, 0})
	ui.selectMenu.rand.dstRect = append(ui.selectMenu.rand.dstRect, &sdl.Rect{0 + 28*1.5, 480 + 5*1.5 - .5 + 50, 72 * 1.5, 20 * 1.5})

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
		if ui.selectMenu.edit.dstRect[0].HasIntersection(clickRect) || ui.selectMenu.edit.dstRect[1].HasIntersection(clickRect) {
			for _, level := range ui.selectMenu.levels {
				if level.isClicked {
					ui.input.updateMouseState()
					globalLevel = &game.Level{}
					globalLevel.LevelName = level.levelName
					editBeforeStart = true
					break
				}
			}
		}
		if ui.selectMenu.start.dstRect[0].HasIntersection(clickRect) || ui.selectMenu.start.dstRect[1].HasIntersection(clickRect) {
			for _, level := range ui.selectMenu.levels {
				if level.isClicked {
					ui.input.updateMouseState()
					globalLevel = &game.Level{}
					globalLevel.LevelName = level.levelName
					editBeforeStart = false
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
			for _, intf := range ui.levelPreviews[i].entities {
				obj := intf.(entityInterface)
				renderer.Copy(textureAtlas, obj.getRect(), &sdl.Rect{int32(obj.getX() + 150), int32(obj.getY()), 32, 32})
			}
		}
		renderer.Copy(level.texture, nil, level.rect)
	}

	for i := 0; i < 2; i++ {
		renderer.Copy(uiAtlas, ui.selectMenu.start.srcRect[i], ui.selectMenu.start.dstRect[i])
		renderer.Copy(uiAtlas, ui.selectMenu.edit.srcRect[i], ui.selectMenu.edit.dstRect[i])
		renderer.Copy(uiAtlas, ui.selectMenu.rand.srcRect[i], ui.selectMenu.rand.dstRect[i])
	}
	renderer.Copy(ui.selectMenu.start.str, nil, ui.selectMenu.start.dstRect[2])
	renderer.Copy(ui.selectMenu.edit.str, nil, ui.selectMenu.edit.dstRect[2])
	renderer.Copy(ui.selectMenu.rand.str, nil, ui.selectMenu.rand.dstRect[2])

	if ui.input.currKeyState[sdl.SCANCODE_ESCAPE] != 0 && ui.input.prevKeyState[sdl.SCANCODE_ESCAPE] == 0 {
		currentState = mainScreen
	}

	return determineToken
}
