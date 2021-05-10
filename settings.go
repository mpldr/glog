package glog

var (
	// TimeFormat specifies the formatting of the timestamp to use. The
	// default is ISO-8601 which corresponds to:
	TimeFormat = "2006-01-02T15:04:05.000000-0700"
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
	// EnableMetaLogging starts logging the inner workings of this library
	// and is usually only used when developing but can be helpful if there
	// are issues with your logging. If you happen to find a bug, please
	// don't hesitate it at: https://todo.sr.ht/~poldi1405/issues
	// (no account needed)
	EnableMetaLogging bool
	// ShortCaller indicates whether the printed caller should be shortened
	// to only include the package and it's function instead of the entire
	// import-path.
	ShortCaller = true
	// ShowCallerLine displays file and line as the caller instead of the
	// function that called.
	ShowCallerLine bool
)

// SetShowCaller allows defining for what levels the caller is displayed in the
// Log. By default the caller is shown for TRACE, DEBUG, ERROR, and FATAL.
func SetShowCaller(lvl Level, show bool) {
	if !isValid(lvl) {
		return
	}

	metalog("set caller for", lvl, "to", show)
	showCaller[lvl] = show
}
