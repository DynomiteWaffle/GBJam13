package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// controlls arrows/z/x

// Game implements ebiten.Game interface.
type Game struct {
	palette    *ebiten.Image
	logo       *ebiten.Image
	background *ebiten.Image
	buttons    *ebiten.Image
	numbers    *ebiten.Image
	plates     *ebiten.Image
	roll       *ebiten.Image

	EPPS        *ebiten.Shader
	EPPSOptions *ebiten.DrawRectShaderOptions

	buffer *ebiten.Image
}

// load assets
//
//go:embed EPPS.kage
var EPPS []byte
var game = &Game{}

//go:embed assets/GBJam13Logo.png
var rawLogo []byte

//go:embed assets/BG.png
var rawBackground []byte

//go:embed assets/Buttons.png
var rawButtons []byte

//go:embed assets/numbers.png
var rawNumbers []byte

//go:embed assets/Plates.png
var rawPlates []byte

//go:embed assets/Roll.png
var rawRoll []byte

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.buffer.Clear()
	// draw to buffer

	g.buffer.DrawImage(g.logo, &ebiten.DrawImageOptions{})
	g.buffer.DrawImage(g.background, &ebiten.DrawImageOptions{})

	// real draw
	screen.DrawRectShader(g.buffer.Bounds().Dx(), g.buffer.Bounds().Dy(), g.EPPS, g.EPPSOptions)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

func main() {

	var err error
	// load assets
	// load shader
	game.EPPS, err = ebiten.NewShader(EPPS)
	if err != nil {
		log.Fatal(err)
	}

	// game.logo, _, err = ebitenutil.NewImageFromFile("GBJam13Logo.png")
	game.logo, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawLogo))
	if err != nil {
		log.Fatal(err)
	}
	game.background, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawBackground))
	if err != nil {
		log.Fatal(err)
	}
	game.buttons, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawButtons))
	if err != nil {
		log.Fatal(err)
	}
	game.numbers, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawNumbers))
	if err != nil {
		log.Fatal(err)
	}
	game.plates, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawPlates))
	if err != nil {
		log.Fatal(err)
	}
	game.roll, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(rawRoll))
	if err != nil {
		log.Fatal(err)
	}
	// init palette
	game.palette = ebiten.NewImage(160, 144)
	game.buffer = ebiten.NewImage(160, 144)

	game.palette.Set(1, 0, color.RGBA{R: 56, G: 28, B: 46, A: 255})
	game.palette.Set(2, 0, color.RGBA{R: 105, G: 109, B: 109, A: 255})
	game.palette.Set(3, 0, color.RGBA{R: 255, G: 166, B: 84, A: 255})
	game.palette.Set(4, 0, color.RGBA{R: 255, G: 215, B: 101, A: 255})

	// init shader opts
	game.EPPSOptions = &ebiten.DrawRectShaderOptions{}
	game.EPPSOptions.Images[0] = game.palette
	game.EPPSOptions.Images[1] = game.buffer

	// Specify the window size as you like. Here, a doubled size is specified.
	// ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GB Jam 13")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
