package engine

import (
	"image/color"

	"github.com/clone-tris/gotris/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const SW = config.SQUARE_WIDTH
const BW = config.SQUARE_BORDER_WIDTH
const IW = SW - (BW * 2)

func DrawGuide(canvas *ebiten.Image, x int, y int, width int, height int) {
	vector.FillRect(canvas, float32(x), float32(y), float32(width), float32(height), config.UI_BACKGROUND, false)

	rows := height / SW
	columns := width / SW

	for i := 0; i < rows+1; i++ {
		lineY := y + i*SW
		vector.StrokeLine(canvas, float32(x), float32(lineY), float32(x+width), float32(lineY), 1, config.UI_GUIDE, false)
	}

	for i := 0; i < columns+1; i++ {
		lineX := x + i*SW
		vector.StrokeLine(canvas, float32(lineX), float32(y), float32(lineX), float32(y+height), 1, config.UI_GUIDE, false)
	}
}

func DrawSquare(canvas *ebiten.Image, x int, y int, color color.RGBA) {
	vector.FillRect(canvas, float32(x), float32(y), float32(SW), float32(SW), color, false)
	// left
	fillPolygon(canvas, []point{{x, y}, {x + BW, y + BW}, {x + BW, y + SW - BW}, {x, y + SW}}, config.SQUARE_BORDER_SIDE)
	// right
	fillPolygon(canvas, []point{{x + SW, y}, {x + SW - BW, y + BW}, {x + SW - BW, y + SW - BW}, {x + SW, y + SW}}, config.SQUARE_BORDER_SIDE)
	// top
	fillPolygon(canvas, []point{{x, y}, {x + BW, y + BW}, {x + SW - BW, y + BW}, {x + SW, y}}, config.SQUARE_BORDER_TOP)
	// bottom
	fillPolygon(canvas, []point{{x, y + SW}, {x + BW, y + SW - BW}, {x + SW - BW, y + SW - BW}, {x + SW, y + SW}}, config.SQUARE_BORDER_BOTTOM)
}

type point struct {
	x int
	y int
}

func fillPolygon(canvas *ebiten.Image, points []point, clr color.RGBA) {
	if len(points) == 0 {
		return
	}

	var path vector.Path
	path.MoveTo(float32(points[0].x), float32(points[0].y))
	for _, p := range points[1:] {
		path.LineTo(float32(p.x), float32(p.y))
	}
	path.Close()

	var cs ebiten.ColorScale
	cs.ScaleWithColor(clr)
	vector.FillPath(canvas, &path, nil, &vector.DrawPathOptions{ColorScale: cs})
}
