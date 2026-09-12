package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const SW = SQUARE_WIDTH
const BW = SQUARE_BORDER_WIDTH
const IW = SW - (BW * 2)

func drawGuide(canvas *ebiten.Image, x int, y int, width int, height int) {
	vector.FillRect(canvas, float32(x), float32(y), float32(width), float32(height), UI_BACKGROUND, false)

	rows := height / SW
	columns := width / SW

	for i := 0; i < rows; i++ {
		lineY := y + i*SW
		vector.StrokeLine(canvas, float32(x), float32(lineY), float32(x+width), float32(lineY), 1, UI_GUIDE, false)
	}

	for i := 0; i < columns; i++ {
		lineX := x + i*SW
		vector.StrokeLine(canvas, float32(lineX), float32(y), float32(lineX), float32(y+height), 1, UI_GUIDE, false)
	}
}
