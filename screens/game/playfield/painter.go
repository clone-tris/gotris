package playfield

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Painter struct {
	Width  int16
	Height int16
}

func newPainter(width int16, height int16) *Painter {
	return &Painter{Width: width, Height: height}
}

func (this *Painter) drawGrid(canvas *ebiten.Image) {
	rows := this.Height / config.SQUARE_WIDTH
	columns := this.Width / config.SQUARE_WIDTH

	var guideColor config.UiColor = config.GUIDE

	for i := int16(0); i < rows+1; i++ {
		lineY := i * config.SQUARE_WIDTH
		vector.StrokeLine(
			canvas, 0, float32(lineY), float32(this.Width), float32(lineY), 1, guideColor, true,
		)
	}

	for i := int16(0); i < columns+1; i++ {
		lineX := i * config.SQUARE_WIDTH
		vector.StrokeLine(
			canvas, float32(lineX), 0, float32(lineX), float32(this.Height), 1, guideColor, true,
		)
	}
}

func (this *Painter) drawBackground(canvas *ebiten.Image) {
	canvas.Fill(config.BACKGROUND)
}
