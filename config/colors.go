package config

import (
	"fmt"
	"image/color"
)

func ParseHexColor(s string) color.RGBA {
	c := color.RGBA{}
	c.A = 0xff

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

var TETROMINO_CYAN = ParseHexColor("#6DECEE")
var TETROMINO_BLUE = ParseHexColor("#0014E6")
var TETROMINO_ORANGE = ParseHexColor("#E4A338")
var TETROMINO_YELLOW = ParseHexColor("#F0EF4F")
var TETROMINO_GREEN = ParseHexColor("#6EEB47")
var TETROMINO_PURPLE = ParseHexColor("#9225E7")
var TETROMINO_RED = ParseHexColor("#DC2F20")

var SQUARE_DEFAULT_COLOR = ParseHexColor("#cc8081")
var SQUARE_BORDER_TOP = ParseHexColor("#ffffffb3")
var SQUARE_BORDER_BOTTOM = ParseHexColor("#00000080")
var SQUARE_BORDER_SIDE = ParseHexColor("#0000001a")

var UI_BACKGROUND = ParseHexColor("#333333")
var UI_SIDEBAR_BACKGROUND = ParseHexColor("#545454")
var UI_POPUP_BACKGROUND = ParseHexColor("#212121")
var UI_GUIDE = ParseHexColor("#555555")
var UI_WHITE_TEXT = ParseHexColor("#FFFFFF")
var UI_POPUP_TEXT = ParseHexColor("#EFEFEF")
