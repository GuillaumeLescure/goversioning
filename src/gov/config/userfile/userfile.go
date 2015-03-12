package userfile

import (
	"gov/config/version"
	"gov/base/logs"
	"encoding/xml"
	"os"
	"os/user"
)

type (
	self struct {
		firstName   string
		lastName    string
		mail        string
		editor      string
		verbose     uint
		color       bool
	}

	dataFromFile struct {
		XMLName     xml.Name    `xml:"GoVersioning"`
		Type        string      `xml:"type,attr"`
		Version     string      `xml:"version,attr"`
		FirstName   string      `xml:"user>firstName"`
		LastName    string      `xml:"user>lastName"`
		Mail        string      `xml:"user>mail"`
		Editor      string      `xml:"preferences>editor"`
		Verbose     uint        `xml:"preferences>verbose"`
		Color       bool        `xml:"preferences>color"`
	}
)

var (
	this    self

	DefaultPath string
)

func init() {
	currentUser, _ := user.Current()
	DefaultPath = currentUser.HomeDir + "/.gov.conf"
}

func ReadOrCreateFile(path string) {
	if _, err := os.Stat(path); err == nil {
		LoadFile(path)
	} else {
		WriteDefaultFile(path)
	}
}

func LoadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		logs.Error(err);
	}
	defer file.Close()

	data := dataFromFile{}
	err = xml.NewDecoder(file).Decode(&data)
	if err != nil {
		logs.Error(err)
	}

	logs.Warning(data)
}

func WriteDefaultFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		logs.Error(err);
	}
	defer file.Close()

	data := dataFromFile{XMLName:xml.Name{" ", "GoVersioning"}, Type:"userConfFile", Version:version.Version(), FirstName:"???", LastName:"???", Mail:"???", Editor:"vi", Verbose:0, Color:false}
	xmlToWrite, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		logs.Error(err)
	}

	_, err = file.Write(xmlToWrite)
	if err != nil {
		logs.Error(err);
	}
}

func FirstName() string {
	return this.firstName
}

func LastName() string {
	return this.lastName
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
