package components

import (
	"image/color"
	"math/rand/v2"

	"github.com/clone-tris/gotris/config"
)

type Type int

const (
	I Type = iota
	O
	T
	J
	L
	S
	Z
)

var COLORS = [...]color.RGBA{
	config.TETROMINO_CYAN,   // I
	config.TETROMINO_YELLOW, // O
	config.TETROMINO_PURPLE, // T
	config.TETROMINO_BLUE,   // J
	config.TETROMINO_ORANGE, // L
	config.TETROMINO_GREEN,  // S
	config.TETROMINO_RED,    // Z
}

var GRIDS = [...][SQUARES_IN_SHAPE][2]int{
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, // I
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, // O
	{{0, 0}, {0, 1}, {0, 2}, {1, 1}}, // T
	{{0, 0}, {1, 0}, {1, 1}, {1, 2}}, // J
	{{0, 0}, {0, 1}, {0, 2}, {1, 0}}, // L
	{{0, 1}, {0, 2}, {1, 0}, {1, 1}}, // S
	{{0, 0}, {0, 1}, {1, 1}, {1, 2}}, // Z
}

func makeSquares(t Type) [4]Square {
	var color = COLORS[t]
	var grid = GRIDS[t]
	var squares [SQUARES_IN_SHAPE]Square

	for i := range SQUARES_IN_SHAPE {
		var cell = grid[i]
		squares[i] = Square{
			row:    cell[0],
			column: cell[1],
			color:  color,
		}

	}

	return squares
}

func random() Shape {
	t := Type(rand.IntN(len(GRIDS)))
	squares := makeSquares(t)
	return newShape(0, 0, squares)
}
