package logs

import (
	"gov/base/colors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

const (
	information = "INFORMATION ->"
	warning     = "WARNING ->"
	error       = "ERROR ->"
)

type logs struct {
	color   bool
	verbose uint
}

var (
	this logs
)

func Color() bool {
	return this.color
}

func SetColor(color bool) {
	this.color = color
}

func Verbose() uint {
	return this.verbose
}

func SetVerbose(verbose uint) {
	this.verbose = verbose
}

func Information(verbose uint, msg ...interface{}) {
	if this.verbose >= verbose {
		if this.color == true {
			log.Println(colors.StringColored(colors.Bold, colors.Green, colors.Black, information) + " " + fmt.Sprint(msg...))
		} else {
			log.Println(information, fmt.Sprint(msg...))
		}
	}
}

func Warning(msg ...interface{}) {
	if this.verbose > 0 {
		if this.color == true {
			fmt.Println(colors.StringColored(colors.Bold, colors.Orange, colors.Black, warning) + " " + fmt.Sprint(msg...))
		} else {
			log.Println(warning + " " + fmt.Sprint(msg...))
		}
	}
}

func Error(msg ...interface{}) {
	if this.color == true {
		log.Println(colors.StringColored(colors.Bold, colors.Red, colors.Black, error) + " " + fmt.Sprint(msg...))
	} else {
		log.Println(error + " " + fmt.Sprint(msg...))
	}
	debug.PrintStack()
	os.Exit(1)
}
