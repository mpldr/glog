package glog

var (
	// TimeFormat specifies the formatting of the timestamp to use. The
	// default is ISO-8601 which corresponds to:
	TimeFormat = "2006-01-02T15:04:05.000000-0700"
	// NoPanic indicates that after a call to Fatal(f) no panic-stop shall
	// be executed.
	NoPanic bool
	// LogFormatter contains the function used to actually format logged
	// statements. Changing this allows complete control over the format of
	// log messages.
	LogFormatter FormatFunction = defaultLogFormat
	//           Level:\tTimestamp – Caller (optional) – Message
	logFormat = "%s:\t%s – %s%s\n"
	// showCaller contains whether a certain loglevel should show it's caller.
	showCaller = [...]bool{
		true,  // TRACE
		true,  // DEBUG
		false, // INFO
		false, // WARN
		true,  // ERROR
		true,  // FATAL
	}
)

// SetShowCaller allows defining for what levels the caller is displayed in the
// Log. By default the caller is shown for TRACE, DEBUG, ERROR, and FATAL.
func SetShowCaller(lvl Level, show bool) {
	if !isValid(lvl) {
		return
	}

	showCaller[lvl] = show
}
