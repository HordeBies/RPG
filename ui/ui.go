package ui

import (
	"bufio"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight = 800, 600

type stateFunc func(*UI2d) stateFunc

type gameState int

const (
	mainScreen gameState = 0
	editScreen gameState = 1
	inGame     gameState = 2
)

var currentState gameState = mainScreen

var renderer *sdl.Renderer
var textureAtlas *sdl.Texture
var textureIndex map[game.Tile][]sdl.Rect

type inputState struct {
	leftButton      bool
	prevLeftButton  bool
	rightButton     bool
	prevRightButton bool
	x, y            int
	currKeyState    []uint8
	prevKeyState    []uint8
}

func (result *inputState) updateMouseState() {
	result.prevLeftButton = result.leftButton
	result.prevRightButton = result.rightButton
	mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()
	result.x = int(mouseX)
	result.y = int(mouseY)
	result.leftButton = leftButton != 0
	result.rightButton = rightButton != 0

}
func (result *inputState) updateKeyboardState() {
	for i := range result.currKeyState {
		result.prevKeyState[i] = result.currKeyState[i]
	}
}

func init() {
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
	textureAtlas = imgFileToTexture("ui/assets/tiles.png")
	loadTextureIndex()

}

func loadTextureIndex() {
	textureIndex = make(map[game.Tile][]sdl.Rect)
	infile, err := os.Open("ui/assets/atlas-index.txt")
	if err != nil {
		panic(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		tileRune := game.Tile(line[0])
		xy := line[2:]
		splitXYC := strings.Split(xy, ",")
		x, err := strconv.ParseInt(splitXYC[0], 10, 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseInt(splitXYC[1], 10, 64)
		if err != nil {
			panic(err)
		}
		variationCount, err := strconv.ParseInt(splitXYC[2], 10, 64)
		if err != nil {
			panic(err)
		}

		var rects []sdl.Rect
		for i := 0; i < int(variationCount); i++ {
			rects = append(rects, sdl.Rect{X: int32(x * 32), Y: int32(y * 32), W: 32, H: 32})
			x++
			if x > 62 {
				x = 0
				y++
			}
		}
		textureIndex[tileRune] = rects

	}
}

func imgFileToTexture(filename string) *sdl.Texture {
	infile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		panic(err)
	}
	tex.Update(nil, pixels, w*4)

	err = tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}
	return tex
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
	layers []layer
	mc     mainCharacter
	input  *inputState
}

func createLayers(level *game.Level, ui *UI2d) {
	gridWorld := level.GridWorld
	for y, row := range gridWorld.Rows {
		for x, grid := range row.Grids {
			for i, layer := range grid.Layers {
				if layer != game.Blank {
					srcRects := textureIndex[layer]
					ui.layers[i].srcRect[y][x] = &srcRects[rand.Intn(len(srcRects))]
					ui.layers[i].dstRect[y][x] = &sdl.Rect{X: int32(x) * 32, Y: int32(y) * 32, W: 32, H: 32}

					renderer.Copy(textureAtlas, ui.layers[0].srcRect[y][x], ui.layers[0].dstRect[y][x])
				}
			}
		}
	}
}

func determineToken(ui *UI2d) stateFunc {
	switch currentState {
	case mainScreen:
		return mainMenu(ui)
	case editScreen:
		return editMenu(ui)
	default:
		return nil
	}
}

func getBlackPixel() *sdl.Texture {
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic(err)
	}
	pixels := make([]byte, 4)
	pixels[0] = 0
	pixels[1] = 0
	pixels[2] = 0
	pixels[3] = 0
	tex.Update(nil, pixels, 4)
	return tex
}

func (ui *UI2d) Draw(level *game.Level, layerCount int) {
	var input inputState
	input.updateMouseState()
	input.currKeyState = sdl.GetKeyboardState()
	input.prevKeyState = make([]uint8, len(input.currKeyState))
	input.updateKeyboardState()

	currentState = editScreen

	ui.layers = make([]layer, layerCount)
	ui.input = &input
	createLayers(level, ui)
	blackPixel := getBlackPixel()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) { // theEvent := event.(type) //remember this
			case *sdl.QuitEvent:
				return
			}
		}
		determineToken(ui)

		renderer.Copy(blackPixel, nil, &sdl.Rect{0, 0, winWidth, winHeight})
		for l := range ui.layers {
			for y := 0; y < 100; y++ {
				for x := 0; x < 100; x++ {
					if ui.layers[l].dstRect[y][x] != nil {
						renderer.Copy(textureAtlas, ui.layers[l].srcRect[y][x], ui.layers[l].dstRect[y][x])
					}
				}
			}
		}
		//fmt.Println(ui.layers[1].srcRect[0][0], ui.layers[1].dstRect[0][0])
		renderer.Present()
		//elapsedTime := time.Until(currTime).Milliseconds()
		//fmt.Println("Ms:", elapsedTime)
		sdl.Delay(16)
		input.updateKeyboardState()
		input.updateMouseState()
	}
}
