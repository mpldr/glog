package glog

import (
	"fmt"
	"sync"
	"time"

	"git.sr.ht/~poldi1405/go-ansi"
)

// StyleFunction is a kind of function that is used for styling. I know, who
// would have thought.
type StyleFunction func(...interface{}) string

// FormatFunction allow defining custom logging formats and changing around the
// order. If Caller is not set all I can do at the moment is passing an empty
// string.
//
// The string arguments correspond to caller and message.
type FormatFunction func(Level, time.Time, string, string) string

// so… I hope https://github.com/golang/go/issues/45624 goes through and we have
// a safe way of ensuring that a Caller is indeed unset
// var EmptyCaller = &string("")

func defaultLogFormat(lvl Level, timestamp time.Time, caller string, message string) string {
	if caller != "" {
		caller += " – "
	}

	return fmt.Sprintf(logFormat, lvl.Short(), timestamp.Format(TimeFormat), caller, message)
}

var (
	styleFuncMtx sync.RWMutex
	styleFuncs   = [...]StyleFunction{
		ansi.Blue,   // TRACE
		ansi.Cyan,   // DEBUG
		ansi.Green,  // INFO
		ansi.Yellow, // WARN
		ansi.Red,    // ERROR
		FatalStyle,  // FATAL
	}
)

// FatalStyle is the default style for fatal logmessages. Just in case you want
// to restore it after changing it.
func FatalStyle(content ...interface{}) string {
	return ansi.Bold(ansi.Red(content...))
}

// SetStyle allows customizing the look of the *error-level*. This can also be
// used for changing the names of the loglevels.
func SetStyle(lvl Level, styler StyleFunction) {
	if !isValid(lvl) {
		return
	}

	styleFuncMtx.Lock()
	defer styleFuncMtx.Unlock()

	styleFuncs[lvl] = styler
}

// PlainStyle is an implementation of StyleFunction that does not format output.
func PlainStyle(content ...interface{}) string {
	return fmt.Sprint(content...)
}
