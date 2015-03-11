package version

import (
	"strconv"
)

type (
	self struct {
		version string
	}
)

const (
	MajorVersion    uint    = 1
	MinorVersion    uint    = 0
	BuildVersion    uint    = 0
)

var (
	this    self
)

func init() {
	this.version = strconv.FormatInt((int64)(MajorVersion), 10) + "." +
		strconv.FormatInt((int64)(MinorVersion), 10) + "-" +
		strconv.FormatInt((int64)(BuildVersion), 10)
}

func Version() string {
	return this.version
}
