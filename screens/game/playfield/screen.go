package playfield

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Playfield struct {
	Width   int16
	Height  int16
	Painter *Painter
}

func NewPlayfield(width int16, height int16) *Playfield {
	return &Playfield{Width: width, Height: height, Painter: newPainter(width, height)}
}

func (this *Playfield) Paint(canvas *ebiten.Image) {
	this.Painter.drawBackground(canvas)
	this.Painter.drawGrid(canvas)
}
