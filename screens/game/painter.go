package game

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type Painter struct {
	Width  int
	Height int
	Canvas *ebiten.Image
}

func newPainter(width int, height int) *Painter {
	return &Painter{Width: width, Height: height, Canvas: ebiten.NewImage(int(width), int(height))}
}

func (this *Painter) stitch(sidebarCanvas *ebiten.Image, playfieldCanvas *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(config.SIDEBAR_WIDTH, 0)

	this.Canvas.DrawImage(playfieldCanvas, drawOptions)
}
