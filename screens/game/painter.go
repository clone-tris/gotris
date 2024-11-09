package game

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Painter struct {
	Canvas *ebiten.Image
}

func NewPainter() *Painter {
	return &Painter{Canvas: ebiten.NewImage(int(config.CANVAS_WIDTH), int(config.CANVAS_HEIGHT))}
}

func (this *Painter) DrawSidebar() {
	this.DrawSidebarBackground()
}

func (this *Painter) DrawSidebarBackground() {
	vector.DrawFilledRect(this.Canvas, 0, 0, config.SIDEBAR_WIDTH, float32(config.CANVAS_HEIGHT), config.SIDEBAR_BACKGROUND, false)
}
