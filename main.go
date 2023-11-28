package main

import (
	"gotris/config"
	"gotris/screens/game/playfield"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playfield *playfield.Playfield
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.playfield.Paint(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.CANVAS_WIDTH, config.CANVAS_HEIGHT
}

func main() {
	playfield := &playfield.Playfield{Width: config.WAR_ZONE_WIDTH, Height: config.CANVAS_HEIGHT}

	game := &Game{playfield: playfield}
	ebiten.SetWindowSize(config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Gotris")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
