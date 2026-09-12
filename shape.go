package main

const SQUARES_IN_SHAPE = 4

type Shape struct {
	row     int
	column  int
	width   int
	height  int
	squares [SQUARES_IN_SHAPE]Square
}

func newShape(row int, column int, squares [SQUARES_IN_SHAPE]Square) Shape {
	shape := Shape{row: row, column: column, squares: squares}
	shape.computeSize()
	return shape
}

func (self *Shape) computeSize() {
	minRow := PUZZLE_HEIGHT
	maxRow := 0
	minColumn := PUZZLE_WIDTH
	maxColumn := 0

	for _, square := range self.squares {
		maxRow = max(square.row, maxRow)
		minRow = min(square.row, minRow)
		maxColumn = max(square.column, maxColumn)
		minColumn = min(square.column, minColumn)
	}

	self.height = maxRow - minRow + 1
	self.width = maxColumn - minColumn + 1

}
