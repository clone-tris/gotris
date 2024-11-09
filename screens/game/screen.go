package game

import "github.com/hajimehoshi/ebiten/v2"

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
	this.Painter.DrawSidebar()

	screen.DrawImage(this.Painter.Canvas, &ebiten.DrawImageOptions{})

}
