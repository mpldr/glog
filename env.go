package glog

import (
	"fmt"
	"os"
	"strings"

	"mpldr.codes/ansi"
)

var (
	stdout os.FileInfo
	stderr os.FileInfo
	// OverrideColor allows overwriting the coloring mode. 0 = auto;
	// 1 = always; -1 = never
	//
	// This can also be set using the Environment-variable `GLOG_COLOR`
	OverrideColor int8
	// LogLevel indicates the level of verbosity to use when logging.
	// Messages below the specified level are discarded. Usually you want to
	// use SetLevel() to ensure that overrides are correctly applied.
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

	setupEnv()
}

func setupEnv() {
	switch os.Getenv("GLOG_METALOGGER") {
	case "1":
		EnableMetaLogging = true
	}

	// Just because
	ansi.EnableANSI()

	switch strings.ToUpper(os.Getenv("GLOG_COLOR")) {
	case "1", "ON", "ALWAYS":
		OverrideColor = 1
		metalog("set colormode to always from environment")
		ansi.OverrideNoColor()
	case "-1", "0", "OFF", "NEVER":
		OverrideColor = -1
		metalog("set colormode to never from environment")
	default:
		OverrideColor = 0
	}

	switch strings.ToUpper(os.Getenv("GLOG_LEVEL")) {
	case "TRACE":
		fallthrough
	case "VERBOSE":
		LogLevel = TRACE
		levelSetFromEnv = true
		metalog("set loglevel to trace from environment")
	case "DEBUG":
		LogLevel = DEBUG
		levelSetFromEnv = true
		metalog("set loglevel to debug from environment")
	case "INFO":
		LogLevel = INFO
		levelSetFromEnv = true
		metalog("set loglevel to info from environment")
	case "WARN", "WARNING":
		LogLevel = WARNING
		levelSetFromEnv = true
		metalog("set loglevel to warning from environment")
	case "ERROR":
		LogLevel = ERROR
		levelSetFromEnv = true
		metalog("set loglevel to error from environment")
	case "FATAL":
		LogLevel = FATAL
		levelSetFromEnv = true
		metalog("set loglevel to fatal from environment")
	case "MUTE", "SILENT":
		LogLevel = 42
		levelSetFromEnv = true
		metalog("set loglevel to silent from environment")
	}
}
