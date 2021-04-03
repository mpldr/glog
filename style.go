package glog

import (
	"fmt"
	"sync"

	"git.sr.ht/~poldi1405/go-ansi"
)

// StyleFunction is a kind of function that is used for styling. I know, who
// would have thought.
type StyleFunction func(...interface{}) string

var (
	styleFuncMtx sync.RWMutex
	styleFuncs   = [...]StyleFunction{
		ansi.Blue,   // TRACE
		ansi.Cyan,   // DEBUG
		ansi.Green,  // INFO
		ansi.Yellow, // WARN
		ansi.Red,    // ERROR
		fatalStyle,  // FATAL
	}
)

func fatalStyle(content ...interface{}) string {
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
