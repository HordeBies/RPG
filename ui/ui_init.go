package ui

import (
	"math/rand"
	"time"

	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const winWidth, winHeight = 800, 600

type stateFunc func(*UI2d) stateFunc

type gameState int

const (
	mainScreen   gameState = 0
	selectScreen gameState = 1
	inGame       gameState = 2
	editLevel    gameState = 3
)

var currentState gameState = mainScreen

var renderer *sdl.Renderer
var textureAtlas *sdl.Texture
var textureIndex map[game.Tile][]sdl.Rect
var blackPixel *sdl.Texture
var font *ttf.Font

type inputState struct {
	leftButton      bool
	prevLeftButton  bool
	rightButton     bool
	prevRightButton bool
	x, y            int
	currKeyState    []uint8
	prevKeyState    []uint8
}

type layer struct {
	srcRect [100][100]*sdl.Rect
	dstRect [100][100]*sdl.Rect
}

type entity struct {
	x, y, layer int
	srcRect     *sdl.Rect
	dstRect     *sdl.Rect
}

type mainCharacter struct {
	entity
}

type UI2d struct {
	levelPreviews []layer
	background    layer
	mc            mainCharacter
	input         *inputState
	mainMenu      mainMenuObj
	selectMenu    selectMenuObj
}

func (ui *UI2d) Init() {
	createMainMenu(ui)
	var input inputState
	input.updateMouseState()
	input.currKeyState = sdl.GetKeyboardState()
	input.prevKeyState = make([]uint8, len(input.currKeyState))
	input.updateKeyboardState()
	ui.input = &input
	createSelectMenu(ui)
}

func init() {
	ttf.Init()
	font, _ = ttf.OpenFont("ui/assets/OpenSans-Regular.ttf", 64)

	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("RPG", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	/*explosionBytes, audioSpec := sdl.LoadWAV("29301__junggle__btn121.wav")
	audioID, err := sdl.OpenAudioDevice("", false, audioSpec, nil, 0)
	if err != nil {
		panic(err)
	}
	defer sdl.FreeWAV(explosionBytes)*/
	rand.Seed(time.Now().UTC().UnixNano())
	blackPixel = createOnePixel(0, 0, 0, 0)
	textureAtlas = imgFileToTexture("ui/assets/tiles.png")
	loadTextureIndex()

}
