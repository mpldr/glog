package glog

import (
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~poldi1405/go-ansi"
)

var (
	stdout        os.FileInfo
	stderr        os.FileInfo
	overwriteColor int8
)

func init() {
	var err error
	stdout, err = os.Stdout.Stat()
	if err != nil {
		panic(fmt.Sprint("glog: cannot stat stdout:", err))
	}
	stdout, err = os.Stderr.Stat()
	if err != nil {
		panic(fmt.Sprint("glog: cannot stat stderr:", err))
	}

	// Just because
	ansi.EnableANSI()

	switch strings.ToUpper(os.Getenv("GLOG_COLOR")) {
	case "ON":
		fallthrough
	case "ALWAYS":
		overwriteColor = 1
	case "OFF":
		fallthrough
	case "NEVER":
		overwriteColor = -1
	}

	switch strings.ToUpper(os.Getenv("GLOG_LEVEL")) {
	case "TRACE":
		fallthrough
	case "VERBOSE":
		LogLevel = TRACE
	case "DEBUG":
		LogLevel = DEBUG
	case "INFO":
		LogLevel = INFO
	case "ERROR":
		LogLevel = ERROR
	case "FATAL":
		LogLevel = FATAL
	case "MUTE":
		fallthrough
	case "SILENT":
		LogLevel = 42
	}
}
