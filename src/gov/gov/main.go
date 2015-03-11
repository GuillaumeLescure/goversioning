package main

import (
	"gov/base/logs"
	"gov/config/flags"
	"gov/config/userfile"
)

func main() {
	flags.Parse()
	logs.SetColor(flags.Color())
	logs.SetVerbose(flags.Verbose())

	err := userfile.ReadAndSetFile("~/.gov.conf")
	if err != nil {
		logs.Error(err)
	}
}
