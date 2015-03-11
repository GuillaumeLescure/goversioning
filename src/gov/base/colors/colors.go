package colors

import (
	"strconv"
)

type (
	Color uint
	Style uint
)

const (
	Black Color = iota
	Red
	Green
	Orange
	Blue
	Magenta
	Cyan
	White
	LightGray
)

const (
	Normal Style = iota
	Bold
	_
	Italic
	Underline
	Blink
	_
	Reverse
)

const (
	begin     = "\033["
	end       = "m"
	separator = ";"
	reset     = begin + "0" + end
)

func StringColored(style Style, foregroundColor Color, backgroundColor Color, msg string) string {
	return begin + strconv.FormatUint((uint64)(style), 10) + separator + strconv.FormatUint(30+(uint64)(foregroundColor), 10) + separator + strconv.FormatUint(30+(uint64)(foregroundColor), 10) + end + msg + reset
}
