package components

import "gotris/config"

const SQUARES_IN_SHAPE = 4

type Shape struct {
	Row     int
	Column  int
	Width   int
	Height  int
	Squares [SQUARES_IN_SHAPE]Square
}

func newShape(row int, column int, squares [SQUARES_IN_SHAPE]Square) Shape {
	shape := Shape{Row: row, Column: column, Squares: squares}
	shape.computeSize()
	return shape
}

func (self *Shape) computeSize() {
	minRow := config.PUZZLE_HEIGHT
	maxRow := 0
	minColumn := config.PUZZLE_WIDTH
	maxColumn := 0

	for _, square := range self.Squares {
		maxRow = max(square.Row, maxRow)
		minRow = min(square.Row, minRow)
		maxColumn = max(square.Column, maxColumn)
		minColumn = min(square.Column, minColumn)
	}

	self.Height = maxRow - minRow + 1
	self.Width = maxColumn - minColumn + 1

}
