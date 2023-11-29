package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gotris/config"
	"gotris/screens/game/playfield"
	"log"
)

type Game struct {
	playfield *playfield.Playfield
}

func (this *Game) Update() error {
	return nil
}

func (this *Game) Draw(screen *ebiten.Image) {
	this.playfield.Paint(screen)
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CANVAS_WIDTH, config.CANVAS_HEIGHT
}

func main() {
	game := &Game{
		playfield: playfield.NewPlayfield(config.WAR_ZONE_WIDTH, config.CANVAS_HEIGHT),
	}

	ebiten.SetWindowSize(config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Gotris")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
