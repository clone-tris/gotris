package playfield

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gotris/config"
)

type Painter struct {
	Width  int16
	Height int16
	Canvas *ebiten.Image
}

func newPainter(width int16, height int16) *Painter {
	return &Painter{Width: width, Height: height, Canvas: ebiten.NewImage(int(width), int(height))}
}

func (this *Painter) drawBackground() {
	this.Canvas.Fill(config.BACKGROUND)
}

func (this *Painter) drawGrid() {
	rows := this.Height / config.SQUARE_WIDTH
	columns := this.Width / config.SQUARE_WIDTH

	var guideColor config.UiColor = config.GUIDE

	for i := int16(0); i < rows+1; i++ {
		lineY := i * config.SQUARE_WIDTH
		vector.StrokeLine(
			this.Canvas, 0, float32(lineY), float32(this.Width), float32(lineY), 1, guideColor, true,
		)
	}

	for i := int16(0); i < columns+1; i++ {
		lineX := i * config.SQUARE_WIDTH
		vector.StrokeLine(
			this.Canvas, float32(lineX), 0, float32(lineX), float32(this.Height), 1, guideColor, true,
		)
	}
}
