package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (self *Game) Update() error {
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	drawGuide(screen, 0, 0, CANVAS_WIDTH, CANVAS_HEIGHT)
	for i := range 4 {
		drawSquare(screen, i*SW, 0, TETROMINO_CYAN)
	}
}

func (self *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return CANVAS_WIDTH, CANVAS_HEIGHT
}

func main() {
	ebiten.SetWindowSize(CANVAS_WIDTH, CANVAS_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
