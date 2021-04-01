package glog

import "git.sr.ht/~poldi1405/go-ansi"

// init attempts to automatically enable ANSI-Escape-Sequences on Windows
// Consoles. I just can't bring myself to calling them "terminals"
func init() {
	ansi.EnableANSI()
}
