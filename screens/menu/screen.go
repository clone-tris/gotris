package menu

import (
	"gotris/components"
	"gotris/config"
	"gotris/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const SW = config.SQUARE_WIDTH
const BW = config.SQUARE_BORDER_WIDTH
const IW = SW - (BW * 2)

type Menu struct {
	shape components.Shape
}

func NewMenu() *Menu {
	return &Menu{
		shape: components.Random(),
	}
}

func (self *Menu) Update() {
	print("update")
}

func (self *Menu) Draw(screen *ebiten.Image) {
	engine.DrawGuide(screen, 0, 0, config.CANVAS_WIDTH, config.CANVAS_HEIGHT)

	// todo use draw squares
	for _, square := range graphic {
		engine.DrawSquare(screen, square.Column*SW, square.Row*SW, square.Color)
	}
}

func (self *Menu) KeyDown() {
	print("keyDown")
}

func (self *Menu) MouseButtonUp() {
	print("keyDown")
}
