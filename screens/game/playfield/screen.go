package playfield

import (
	"gotris/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Playfield struct {
	canvas *ebiten.Image
	Width  int16
	Height int16
}

func NewPlayfield(width int16, height int16) *Playfield {
	return &Playfield{Width: width, Height: height}
}

func (playfield *Playfield) Paint(canvas *ebiten.Image) {
	rows := playfield.Height / config.SQUARE_WIDTH
	columns := playfield.Width / config.SQUARE_WIDTH

	for i := int16(0); i < rows+1; i++ {
		lineY := i * config.SQUARE_WIDTH
		vector.StrokeLine(canvas, 0, float32(lineY), float32(playfield.Width), float32(lineY), 1, color.White, true)
	}

	for i := int16(0); i < columns+1; i++ {
		lineX := i * config.SQUARE_WIDTH

		vector.StrokeLine(canvas, float32(lineX), 0, float32(lineX), float32(playfield.Height), 1, color.White, true)
	}
}
