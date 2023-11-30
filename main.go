package main

import (
	"gotris/config"
	"gotris/screens/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gotris struct {
	gameScreen *game.GameScreen
}

func (this *Gotris) Update() error {
	return nil
}

func (this *Gotris) Draw(screen *ebiten.Image) {
	this.gameScreen.Paint()
	screen.DrawImage(this.gameScreen.Canvas(), &ebiten.DrawImageOptions{})
}

func (this *Gotris) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CANVAS_WIDTH, config.CANVAS_HEIGHT
}

func main() {

	gotris := &Gotris{
		gameScreen: game.NewGameScreen(),
	}

	ebiten.SetWindowSize(config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Gotris")

	if err := ebiten.RunGame(gotris); err != nil {
		log.Fatal(err)
	}
}
