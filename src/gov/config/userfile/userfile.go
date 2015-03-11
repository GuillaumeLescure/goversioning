package userfile

import (
	"gov/base/logs"
	"encoding/xml"
	"os"
)

type (
	userFile struct {
		name    string
		mail    string
		editor  string
		verbose uint
		color   bool
	}
)

var (
	this userFile
)

func ReadAndSetFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = xml.NewDecoder(file).Decode(&this)
	if err != nil {
		logs.Error(err)
	}
	return nil
}

func Name() string {
	return this.name
}

func Mail() string {
	return this.mail
}

func Editor() string {
	return this.editor
}

func Verbose() uint {
	return this.verbose
}

func Color() bool {
	return this.color
}
