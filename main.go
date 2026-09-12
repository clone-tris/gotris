package main

import (
	"log"

	"github.com/clone-tris/gotris/config"
	"github.com/clone-tris/gotris/engine"
	"github.com/clone-tris/gotris/screens/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	screen engine.Screen
}

func NewGame() *Game {
	screen := &menu.Menu{}
	return &Game{
		screen: screen,
	}
}

func (self *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		print("q")
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	self.screen.Draw(screen)
}

func (self *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CANVAS_WIDTH, config.CANVAS_HEIGHT
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Gotris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
