package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawGuide(canvas *ebiten.Image, x int, y int, width int, height int) {
	vector.FillRect(canvas, float32(x), float32(y), float32(width), float32(height), TETROMINO_BLUE, false)
}
