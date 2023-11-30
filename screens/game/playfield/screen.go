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

func (this *Playfield) Paint() {
	this.Painter.drawBackground()
	this.Painter.drawGrid()
}

func (this *Playfield) Canvas() *ebiten.Image {
	return this.Painter.Canvas
}
