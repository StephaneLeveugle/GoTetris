package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
	_ "image/png"
)

const WINDOW_WIDTH = 480
const WINDOW_HEIGHT = 480
const GAME_WIDTH = 240
const GAME_HEIGHT = 480

// Game implements ebiten.Game interface.
type Game struct{}

var img *ebiten.Image
var gameBackground *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("resources/l.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}



// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	
	// grey background
	screen.Fill(color.NRGBA{0x44, 0x44, 0x44, 0xff})

	if img != nil {
		gameBackground = ebiten.NewImage(GAME_WIDTH, GAME_HEIGHT)
	}
	gameBackground.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff});
	gameBackgroundOptions := &ebiten.DrawImageOptions{}
	screen.DrawImage(gameBackground, gameBackgroundOptions)

	// tetromino
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.1, 0.1)
	options.GeoM.Translate(GAME_WIDTH/2, 0)
	screen.DrawImage(img, options)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Le super Tetris de Stété et Sysy")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
