package game

import (
	"gotris/config"
	"gotris/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Painter struct {
	MainCanvas      *ebiten.Image
	PlayfieldCanvas *ebiten.Image
	SidebarCanvas   *ebiten.Image
}

func NewPainter() *Painter {
	return &Painter{
		MainCanvas: ebiten.NewImage(
			int(config.CANVAS_WIDTH),
			int(config.CANVAS_HEIGHT),
		),
		PlayfieldCanvas: ebiten.NewImage(config.WAR_ZONE_WIDTH, config.CANVAS_HEIGHT),
		SidebarCanvas:   ebiten.NewImage(config.SIDEBAR_WIDTH, config.CANVAS_HEIGHT),
	}
}

func (this *Painter) DrawSidebar() {
	this.DrawSidebarBackground()

	this.MainCanvas.DrawImage(this.SidebarCanvas, &ebiten.DrawImageOptions{})
}

func (this *Painter) DrawSidebarBackground() {
	vector.DrawFilledRect(this.SidebarCanvas, 0, 0, config.SIDEBAR_WIDTH, float32(config.CANVAS_HEIGHT), config.SIDEBAR_BACKGROUND, false)
}

func (this *Painter) DrawPlayfield() {
	engine.DrawGuide(this.PlayfieldCanvas, int(config.WAR_ZONE_WIDTH), config.CANVAS_HEIGHT)

	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(config.SIDEBAR_WIDTH, 0)
	this.MainCanvas.DrawImage(this.PlayfieldCanvas, drawOptions)
}
