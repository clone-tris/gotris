package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gotris/config"
	"gotris/screens/game/playfield"
)

type GameScreen struct {
	Width     int16
	Height    int16
	Painter   *Painter
	playfield *playfield.Playfield
	sidebar   *ebiten.Image
}

func NewGameScreen() *GameScreen {
	return &GameScreen{
		Painter:   newPainter(config.CANVAS_WIDTH, config.CANVAS_HEIGHT),
		playfield: playfield.NewPlayfield(config.WAR_ZONE_WIDTH, config.CANVAS_HEIGHT),
		sidebar:   &ebiten.Image{},
	}
}

func (this *GameScreen) Paint() {
	this.playfield.Paint()
}

func (this *GameScreen) Canvas() *ebiten.Image {
	this.Painter.stitch(this.sidebar, this.playfield.Canvas())
	return this.Painter.Canvas
}
