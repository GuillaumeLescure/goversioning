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

	userfile.ReadOrCreateFile(userfile.DefaultPath)
}
