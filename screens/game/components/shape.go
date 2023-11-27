package components

import "image/color"

type Shape struct {
	Row    int8
	Column int8
	Width  int8
	Height int8
	Color  color.Color
}
