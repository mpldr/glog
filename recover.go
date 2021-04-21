package glog

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// PanicHandler logs a panic if it occurs. A panic is always written to
// panic.log and then passed on. Under *no* circumstances should this be used as
// a way to ensure a programs stability. Make your own function for that.
func PanicHandler() {
	r := recover()

	if r == nil {
		return
	}

	panicLog, err := os.OpenFile("panic.log", os.O_SYNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		// we tried, not possible. bye
		panic(r)
	}
	defer panicLog.Close()

	// if any error happens here, we do not care. Shit is burning already.
	panicLog.WriteString(strings.Repeat("#", 80) + "\n")
	panicLog.WriteString(strings.Repeat(" ", 34))
	panicLog.WriteString("PANIC CAUGHT!\n")
	panicLog.WriteString(strings.Repeat(" ", 24) + time.Now().Format(TimeFormat) + "\n")
	panicLog.WriteString(strings.Repeat("#", 80) + "\n\n")
	panicLog.WriteString(fmt.Sprintf("Error: %v\n\n", r))
	panicLog.Write(debug.Stack())
	panic(r)
}
