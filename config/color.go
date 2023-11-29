package config

import (
	"fmt"
	"image/color"
)

func ParseHexColor(s string) *color.RGBA {
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
	return &c
}

type TetrominoColor = *color.RGBA
type SquareColor = *color.RGBA
type UiColor = *color.RGBA

var CYAN TetrominoColor = ParseHexColor("#6DECEE")
var BLUE TetrominoColor = ParseHexColor("#0014E6")
var ORANGE TetrominoColor = ParseHexColor("#E4A338")
var YELLOW TetrominoColor = ParseHexColor("#F0EF4F")
var GREEN TetrominoColor = ParseHexColor("#6EEB47")
var PURPLE TetrominoColor = ParseHexColor("#9225E7")
var RED TetrominoColor = ParseHexColor("#DC2F20")

var DEFAULT_SQUARE_COLOR SquareColor = ParseHexColor("#cc8081")
var BORDER_TOP SquareColor = ParseHexColor("#ffffffb3")
var BORDER_BOTTOM SquareColor = ParseHexColor("#00000080")
var BORDER_SIDE SquareColor = ParseHexColor("#0000001a")

var BACKGROUND UiColor = ParseHexColor("#333333")
var SIDEBAR_BACKGROUND UiColor = ParseHexColor("#545454")
var POPUP_BACKGROUND UiColor = ParseHexColor("#212121")
var GUIDE UiColor = ParseHexColor("#555555")
var WHITE_TEXT UiColor = ParseHexColor("#FFFFFF")
var POPUP_TEXT UiColor = ParseHexColor("#EFEFEF")
