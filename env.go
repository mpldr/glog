package glog

import (
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~poldi1405/go-ansi"
)

var (
	stdout os.FileInfo
	stderr os.FileInfo
	// OverwriteColor allows overwriting the coloring mode. 0 = auto;
	// 1 = always; -1 = never
	//
	// This can also be set using the Environment-variable `GLOG_COLOR`
	OverwriteColor int8
	// LogLevel indicates the level of verbosity to use when logging.
	// Messages below the specified level are discarded.
	//
	// This can also be set using the Environment-variable `GLOG_LEVEL`
	LogLevel = WARNING
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
	case "1":
		fallthrough
	case "ON":
		fallthrough
	case "ALWAYS":
		OverwriteColor = 1
	case "-1":
		fallthrough
	case "0":
		fallthrough
	case "OFF":
		fallthrough
	case "NEVER":
		OverwriteColor = -1
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
