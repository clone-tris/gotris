package config

const SQUARE_WIDTH = 24
const SQUARE_BORDER_WIDTH = 3
const PUZZLE_HEIGHT int8 = 20
const PUZZLE_WIDTH int8 = 10

const SIDEBAR_WIDTH = SQUARE_WIDTH * 6
const WAR_ZONE_WIDTH = int(PUZZLE_WIDTH) * SQUARE_WIDTH

const CANVAS_WIDTH = SIDEBAR_WIDTH + WAR_ZONE_WIDTH
const CANVAS_HEIGHT = int(PUZZLE_HEIGHT) * SQUARE_WIDTH
