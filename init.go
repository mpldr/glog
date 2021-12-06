package glog

import (
	"os"

	"mpldr.codes/ansi"
)

// init attempts to automatically enable ANSI-Escape-Sequences on Windows
// Consoles. I just can't bring myself to calling them "terminals"
func init() {
	ansi.EnableANSI()
	// if stdout is a file, add stdout to all levels. the user probably
	// wants all logmessages
	if fi, err := os.Stdout.Stat(); err == nil {
		if fi.Mode().IsRegular() {
			AddOutputAbove(WARNING, os.Stdout)
		}
	}
}
