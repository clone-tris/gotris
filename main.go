package main

import (
	"gotris/config"
	"gotris/engine"
	"gotris/screens/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gotris struct {
	gameScreen engine.Screen
}

func (this *Gotris) Update() error {
	this.gameScreen.Update()
	return nil
}

func (this *Gotris) Draw(screen *ebiten.Image) {
	this.gameScreen.Paint(screen)
}

func (this *Gotris) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CANVAS_WIDTH, config.CANVAS_HEIGHT
}

func main() {

	gotris := &Gotris{
		gameScreen: game.NewScreen(),
	}

	ebiten.SetWindowSize(config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Gotris")

	if err := ebiten.RunGame(gotris); err != nil {
		log.Fatal(err)
	}
}
