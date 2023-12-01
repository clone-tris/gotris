package playfield

import (
	"gotris/config"
	"gotris/screens/game/components"

	"github.com/hajimehoshi/ebiten/v2"
)

type Playfield struct {
	Width   int
	Height  int
	Painter *Painter
	square  *components.Square
}

func NewPlayfield(width int, height int) *Playfield {
	return &Playfield{
		Width:   width,
		Height:  height,
		Painter: newPainter(width, height),
		square:  components.NewSquare(2, 2, config.RED),
	}
}

func (this *Playfield) Paint() {
	this.Painter.drawBackground()
	this.Painter.drawGrid()
	this.square.Draw(this.Painter.Canvas, 0, 0)

}

func (this *Playfield) Canvas() *ebiten.Image {
	return this.Painter.Canvas
}
