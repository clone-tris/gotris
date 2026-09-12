package config

import (
	"fmt"
	"image/color"
)

func ParseHexColor(s string) color.RGBA {
	c := color.RGBA{
		A: 0xff}

	switch len(s) {
	case 7:
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 9:
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	default:
		_ = fmt.Errorf("invalid length, must be 7 or 9")
	}
	return c
}

var TETROMINO_CYAN = ParseHexColor("#00F0F0")
var TETROMINO_BLUE = ParseHexColor("#0000F0")
var TETROMINO_ORANGE = ParseHexColor("#F0A000")
var TETROMINO_YELLOW = ParseHexColor("#F0F000")
var TETROMINO_GREEN = ParseHexColor("#00F000")
var TETROMINO_PURPLE = ParseHexColor("#A000F0")
var TETROMINO_RED = ParseHexColor("#F00000")

var SQUARE_DEFAULT_COLOR = ParseHexColor("#cc8081FF")
var SQUARE_BORDER_TOP = ParseHexColor("#FFFFFF99")
var SQUARE_BORDER_BOTTOM = ParseHexColor("#00000080")
var SQUARE_BORDER_SIDE = ParseHexColor("#0000001A")

var UI_BACKGROUND = ParseHexColor("#333333")
var UI_SIDEBAR_BACKGROUND = ParseHexColor("#545454")
var UI_POPUP_BACKGROUND = TETROMINO_CYAN
var UI_GUIDE = ParseHexColor("#555555")
var UI_WHITE_TEXT = ParseHexColor("#FFFFFF")
var UI_POPUP_TEXT = ParseHexColor("#EFEFEF")
