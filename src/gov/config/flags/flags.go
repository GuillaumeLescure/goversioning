package flags

import (
	"flag"
)

type (
	flags struct {
		verbose uint
		color   bool
	}
)

var (
	this flags
)

func Parse() {
	flag.UintVar(&this.verbose, "v", 1, "verbose mode (0 to 3)")
	flag.BoolVar(&this.color, "c", true, "color mode (true or false)")
	flag.Parse()
}

func Verbose() uint {
	return this.verbose
}

func Color() bool {
	return this.color
}
