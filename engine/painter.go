package engine

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGuide(canvas *ebiten.Image, width int, height int) {
	vector.DrawFilledRect(canvas, 0, 0, float32(width), float32(height), config.BACKGROUND, false)
}
