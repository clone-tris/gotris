package menu

import (
	"github.com/clone-tris/gotris/config"
	"github.com/clone-tris/gotris/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

const SW = config.SQUARE_WIDTH
const BW = config.SQUARE_BORDER_WIDTH
const IW = SW - (BW * 2)

type Menu struct {
}

func (self *Menu) Update() {
	print("update")
}

func (self *Menu) Draw(screen *ebiten.Image) {
	engine.DrawGuide(screen, 0, 0, config.CANVAS_WIDTH, config.CANVAS_HEIGHT)
	for i := range 4 {
		engine.DrawSquare(screen, i*SW, 0, config.TETROMINO_CYAN)
	}
}

func (self *Menu) KeyDown() {
	print("keyDown")
}

func (self *Menu) MouseButtonUp() {
	print("keyDown")
}
