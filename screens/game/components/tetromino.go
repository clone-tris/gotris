package components

import (
	"gotris/config"
)

type Tetromino = string

var T Tetromino = "T"
var Z Tetromino = "Z"
var S Tetromino = "S"
var L Tetromino = "L"
var J Tetromino = "J"
var O Tetromino = "O"
var I Tetromino = "I"

var TetrominoGrid = map[Tetromino][][]int{
	T: {{0, 0}, {0, 1}, {0, 2}, {1, 1}},
	Z: {{0, 0}, {0, 1}, {1, 1}, {1, 2}},
	S: {{0, 1}, {0, 2}, {1, 0}, {1, 1}},
	L: {{0, 0}, {0, 1}, {0, 2}, {1, 0}},
	J: {{0, 0}, {1, 0}, {1, 1}, {1, 2}},
	O: {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	I: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
}

var TetrominoColor = map[Tetromino]config.TetrominoColor{
	T: config.PURPLE,
	Z: config.RED,
	S: config.GREEN,
	L: config.ORANGE,
	J: config.BLUE,
	O: config.YELLOW,
	I: config.CYAN,
}
