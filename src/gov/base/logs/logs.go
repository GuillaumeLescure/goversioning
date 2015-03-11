package logs

import (
	"gov/base/colors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
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

			if this.verbose == 3 {
				moreColoredInformation()
			}
		} else {
			log.Println(information, fmt.Sprint(msg...))

			if this.verbose == 3 {
				moreInformation()
			}
		}
	}
}

func Warning(msg ...interface{}) {
	if this.verbose > 0 {
		if this.color == true {
			fmt.Println(colors.StringColored(colors.Bold, colors.Orange, colors.Black, warning) + " " + fmt.Sprint(msg...))

			if this.verbose == 3 {
				moreColoredInformation()
			}
		} else {
			log.Println(warning + " " + fmt.Sprint(msg...))

			if this.verbose == 3 {
				moreInformation()
			}
		}
	}
}

func Error(msg ...interface{}) {
	if this.color == true {
		log.Println(colors.StringColored(colors.Bold, colors.Red, colors.Black, error) + " " + fmt.Sprint(msg...))

		if this.verbose == 3 {
			moreColoredInformation()
		}
	} else {
		log.Println(error + " " + fmt.Sprint(msg...))

		if this.verbose == 3 {
			moreInformation()
		}
	}
	os.Exit(1)
}

func moreColoredInformation() {
	function, file, line, _ := runtime.Caller(2)
	fmt.Println("(at the line " + colors.StringColored(colors.Underline, colors.White, colors.Black, strconv.FormatInt((int64)(line), 10)) + " of the function '" + colors.StringColored(colors.Underline, colors.White, colors.Black, runtime.FuncForPC(function).Name()) + "' in the file '" + colors.StringColored(colors.Underline, colors.White, colors.Black, file) + "')")
}

func moreInformation() {
	function, file, line, _ := runtime.Caller(2)
	fmt.Println("(at the line " + strconv.FormatInt((int64)(line), 10) + " of the function '" + runtime.FuncForPC(function).Name() + "' in the file '" + file + "')")
}
