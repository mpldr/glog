package glog

import (
	"os"
	"strings"
)

func init() {
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
