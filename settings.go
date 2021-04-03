package glog

var (
	// LogLevel indicates the level of verbosity to use when logging.
	// Messages below the specified level are discarded.
	LogLevel = WARNING
	// TimeFormat specifies the formatting of the timestamp to use. The
	// default is ISO-8601 which corresponds to:
	TimeFormat = "2006-01-02T15:04:05-0700"
	//           Level:\tTimestamp – Message
	logFormat = "%s:\t%s – %s\n"
	//                 Level:\tTimestamp – Caller – Message
	logFormatCaller = "%s:\t%s – %s – %s\n"
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

func SetShowCaller(lvl Level, show bool) {
	if !isValid(lvl) {
		return
	}

	showCaller[lvl] = show
}
