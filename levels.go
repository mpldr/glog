package glog

import (
	"fmt"
	"strings"
)

type Level uint8

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// levelsBelow returns a list of levels equal or above the indicated levels
func levelsAbove(lvl Level) (levels []Level) {
	if !isValid(lvl) {
		return
	}

	for i := lvl; i < 6; i++ {
		levels = append(levels, i)
	}
	return
}

// levelsBelow returns a list of levels equal or below the indicated levels
func levelsBelow(lvl Level) (levels []Level) {
	if !isValid(lvl) {
		return
	}

	// can't go below 0, so we need an upper bound
	for i := lvl; isValid(i); i-- {
		levels = append(levels, i)
	}
	return
}

// levelsBetween returns a list of levels between (inclusive) the indicated levels
func levelsBetween(lowerBound, upperBound Level) (levels []Level) {
	if !isValid(lowerBound) || !isValid(upperBound) {
		return
	}

	if lowerBound > upperBound {
		// why is this variable named like this? simple: if you are here
		// you are an idiot.
		idiot := upperBound
		upperBound = lowerBound
		lowerBound = idiot
	}

	for i := lowerBound; i <= upperBound; i++ {
		levels = append(levels, i)
	}
	return
}

// isValid returns whether a given level is valid
func isValid(lvl Level) bool {
	// can't go below 0 because it's an unsigned integer
	return lvl <= FATAL
}

// SetLevel allows setting the loglevel while preserving the level set using the
// environment
func SetLevel(lvl Level) bool {
	if levelSetFromEnv {
		return false
	}
	LogLevel = lvl
	return true
}

// levelSetFromEnv tells us if the loglevel was set from the env. Just a hack to
// get rid of the error handling now
var levelSetFromEnv bool

// String implements the fmt.Stringer interface to allow use of levels in fmt
// calls.
func (lvl Level) String() string {
	switch lvl {
	case 0:
		return "TRACE"
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARNING"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return "ID-10T"
	}
}

// Short returns a levels representation in log
func (lvl Level) Short() string {
	switch lvl {
	case 0:
		return "TRACE"
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return "CRIT"
	}
}

// ParseLevel takes in a string and returns the corresponding loglevel. If it
// does not exist, default is returned instead.
func ParseLevel(level string, fallback Level) Level {
	var lvls map[string]Level = map[string]Level{
		"TRACE":   TRACE,
		"VERBOSE": TRACE,
		"DEBUG":   DEBUG,
		"INFO":    INFO,
		"WARN":    WARNING,
		"WARNING": WARNING,
		"ERROR":   ERROR,
		"ERR":     ERROR,
		"FATAL":   FATAL,
	}

	if l, ok := lvls[strings.ToUpper(level)]; ok {
		return l
	}

	return fallback
}

// ensure we conform to the stringer interface
var _ fmt.Stringer = new(Level)
