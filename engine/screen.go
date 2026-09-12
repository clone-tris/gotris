package engine

import "github.com/hajimehoshi/ebiten/v2"

type Screen interface {
	Update()
	Draw(screen *ebiten.Image)
	KeyDown()
	MouseButtonUp()
}
