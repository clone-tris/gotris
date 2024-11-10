package game

import (
	"gotris/config"

	"github.com/hajimehoshi/ebiten/v2"
)

var playfieldOptions *ebiten.DrawImageOptions
var sidebarOptions *ebiten.DrawImageOptions

func init() {
	playfieldOptions = &ebiten.DrawImageOptions{}
	playfieldOptions.GeoM.Translate(config.SIDEBAR_WIDTH, 0)
	sidebarOptions = &ebiten.DrawImageOptions{}
}

type GameScreen struct {
	Width   int16
	Height  int16
	Painter *Painter
}

func NewScreen() *GameScreen {
	return &GameScreen{
		Painter: NewPainter(),
	}
}

func (this *GameScreen) Update() {

}

func (this *GameScreen) Paint(screen *ebiten.Image) {
	screen.DrawImage(this.Painter.DrawSidebar(), sidebarOptions)
	screen.DrawImage(this.Painter.DrawPlayfield(), playfieldOptions)
}
