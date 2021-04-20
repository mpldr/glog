package glog

import (
	"fmt"
	"time"
)

// Trace logs a message at the TRACE level
func Trace(message ...interface{}) {
	metalog("received TRACE message", message)
	if LogLevel > TRACE {
		metalog("Level is higher than trace. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(TRACE, msg)
	writeToOutput(TRACE, logLine)
}

// Tracef formats the input values as specified and writes them to the according channels
func Tracef(format string, values ...interface{}) {
	metalog("received TRACE message", format, values)
	if LogLevel > TRACE {
		metalog("Level is higher than trace. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(TRACE, msg)
	writeToOutput(TRACE, logLine)
}

// Debug logs a message at the DEBUG level
func Debug(message ...interface{}) {
	metalog("received DEBUG message", message)
	if LogLevel > DEBUG {
		metalog("Level is higher than debug. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(DEBUG, msg)
	writeToOutput(DEBUG, logLine)
}

// Debugf formats the input values as specified and writes them to the according channels
func Debugf(format string, values ...interface{}) {
	metalog("received DEBUG message", format, values)
	if LogLevel > DEBUG {
		metalog("Level is higher than debug. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(DEBUG, msg)
	writeToOutput(DEBUG, logLine)
}

// Info logs a message at the INFO level
func Info(message ...interface{}) {
	metalog("received INFO message", message)
	if LogLevel > INFO {
		metalog("Level is higher than info. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(INFO, msg)
	writeToOutput(INFO, logLine)
}

// Infof formats the input values as specified and writes them to the according channels
func Infof(format string, values ...interface{}) {
	metalog("received INFO message", format, values)
	if LogLevel > INFO {
		metalog("Level is higher than info. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(INFO, msg)
	writeToOutput(INFO, logLine)
}

// Warn logs a message at the WARNING level
func Warn(message ...interface{}) {
	metalog("received WARNING message", message)
	if LogLevel > WARNING {
		metalog("Level is higher than warning. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(WARNING, msg)
	writeToOutput(WARNING, logLine)
}

// Warnf formats the input values as specified and writes them to the according channels
func Warnf(format string, values ...interface{}) {
	metalog("received WARNING message", format, values)
	if LogLevel > WARNING {
		metalog("Level is higher than warning. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(WARNING, msg)
	writeToOutput(WARNING, logLine)
}

// Error logs a message at the ERROR level
func Error(message ...interface{}) {
	metalog("received ERROR message", message)
	if LogLevel > ERROR {
		metalog("Level is higher than error. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(ERROR, msg)
	writeToOutput(ERROR, logLine)
}

// Errorf formats the input values as specified and writes them to the according channels
func Errorf(format string, values ...interface{}) {
	metalog("received ERROR message", format, values)
	if LogLevel > ERROR {
		metalog("Level is higher than error. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(ERROR, msg)
	writeToOutput(ERROR, logLine)
}

// Fatal logs a message at the FATAL level and Panics afterwards
func Fatal(message ...interface{}) {
	metalog("received FATAL message", message)
	if LogLevel > FATAL {
		metalog("Level is higher than fatal. message discarded")
		return
	}

	msg := fmt.Sprint(message...)
	logLine := getLogLine(FATAL, msg)
	writeToOutput(FATAL, logLine)
}

// Fatalf formats the input values as specified and writes them to the according channels
func Fatalf(format string, values ...interface{}) {
	metalog("received FATAL message", format, values)
	if LogLevel > FATAL {
		metalog("Level is higher than fatal. message discarded")
		return
	}

	msg := fmt.Sprintf(format, values...)
	logLine := getLogLine(FATAL, msg)
	writeToOutput(FATAL, logLine)
}

func getLogLine(lvl Level, message string) string {
	timestamp := time.Now() // get time here so getting the caller does not change it
	caller := ""
	if showCaller[lvl] {
		metalog("caller requested")
		caller = GetCaller(2)
	}
	return LogFormatter(lvl, timestamp, caller, message)
}
