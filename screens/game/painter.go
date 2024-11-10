package game

import (
	"gotris/config"
	"gotris/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Painter struct {
	Canvas          *ebiten.Image
	PlayfieldCanvas *ebiten.Image
	SidebarCanvas   *ebiten.Image
}

func NewPainter() *Painter {
	return &Painter{
		Canvas: ebiten.NewImage(
			int(config.CANVAS_WIDTH),
			int(config.CANVAS_HEIGHT),
		),
		PlayfieldCanvas: ebiten.NewImage(config.WAR_ZONE_WIDTH, config.CANVAS_HEIGHT),
		SidebarCanvas:   ebiten.NewImage(config.SIDEBAR_WIDTH, config.CANVAS_HEIGHT),
	}
}

func (this *Painter) DrawSidebar() *ebiten.Image {
	this.DrawSidebarBackground()

	this.Canvas.DrawImage(this.SidebarCanvas, &ebiten.DrawImageOptions{})

	return this.SidebarCanvas
}

func (this *Painter) DrawSidebarBackground() {
	vector.DrawFilledRect(this.SidebarCanvas, 0, 0, config.SIDEBAR_WIDTH, float32(config.CANVAS_HEIGHT), config.SIDEBAR_BACKGROUND, false)
}

func (this *Painter) DrawPlayfield() *ebiten.Image {
	engine.DrawGuide(this.PlayfieldCanvas, int(config.WAR_ZONE_WIDTH), config.CANVAS_HEIGHT)

	return this.PlayfieldCanvas
}
