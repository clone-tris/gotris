package game

type GameScreen struct {
	Width   int16
	Height  int16
	Painter *Painter
}

func NewGameScreen() *GameScreen {
	return &GameScreen{
		Painter: newPainter(),
	}
}

func (this *GameScreen) Update() {

}

func (this *GameScreen) Paint() {
}
