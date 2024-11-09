package game

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type Painter struct {
	Canvas *ebiten.Image
}

func newPainter() *Painter {
	return &Painter{Canvas: ebiten.NewImage(int(config.CANVAS_WIDTH), int(config.CANVAS_HEIGHT))}
}
