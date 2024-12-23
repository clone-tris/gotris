package components

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Square struct {
	Row    int8
	Column int8
	Color  config.SquareColor
}

func NewSquare(row int8, column int8, color config.SquareColor) *Square {
	return &Square{Row: row, Column: column, Color: color}
}

func (this *Square) Draw(canvas *ebiten.Image, refRow int8, refColumn int8) {
	var width = config.SQUARE_WIDTH
	var x = int(refColumn+this.Column) * width
	var y = int(refRow+this.Row) * width

	vector.DrawFilledRect(canvas, float32(x), float32(y), float32(width), float32(width), this.Color, false)
}
