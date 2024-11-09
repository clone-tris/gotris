package engine

import "github.com/hajimehoshi/ebiten/v2"

type Screen interface {
	Update()
	Paint(screen *ebiten.Image)
}
