package glog

import (
	"fmt"
	"io"
	"os"
	"sync"

	"git.sr.ht/~poldi1405/go-ansi"
)

// set default outputs per loglevel
var outputs = [...][]io.Writer{
	{os.Stdout}, // TRACE
	{os.Stdout}, // DEBUG
	{os.Stdout}, // INFO
	{os.Stderr}, // WARNING
	{os.Stderr}, // ERROR
	{os.Stderr}, // FATAL
}
var outputMtx sync.RWMutex

// SetOutput removes all outputs and replaces them with the specified output. To
// discard log messages to this level set the loglevel accordingly or use
// `io.Discard`
func SetOutput(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = []io.Writer{output}
}

// AddOutput adds the specified output to the list of outputs for the level
func AddOutput(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = append(outputs[lvl], output)
}

// SetOutputBetween removes all outputs and replaces them with the specified
// output. This is executed for all specified levels.
func SetOutputBetween(lowerLevel, upperLevel Level, output io.Writer) {
	if !isValid(lowerLevel) || !isValid(upperLevel) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	levels := levelsBetween(lowerLevel, upperLevel)
	for _, lvl := range levels {
		outputs[lvl] = []io.Writer{output}
	}
}

// AddOutputBetween adds the specified output to the list of outputs for all
// specified levels.
func AddOutputBetween(lowerLevel, upperLevel Level, output io.Writer) {
	if !isValid(lowerLevel) || !isValid(upperLevel) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	levels := levelsBetween(lowerLevel, upperLevel)
	for _, lvl := range levels {
		outputs[lvl] = append(outputs[lvl], output)
	}
}

// SetOutputAbove removes all outputs and replaces them with the specified
// output. This is repeated for every specified level.
func SetOutputAbove(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = []io.Writer{output}
}

// AddOutput adds the specified output to the list of outputs for the specified
// levels.
func AddOutputAbove(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = append(outputs[lvl], output)
}

// SetOutput removes all outputs and replaces them with the specified output.
// This is repeated for every specified level.
func SetOutputBelow(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = []io.Writer{output}
}

// AddOutput adds the specified output to the list of outputs for the specified
// levels.
func AddOutputBelow(lvl Level, output io.Writer) {
	if !isValid(lvl) {
		return
	}

	outputMtx.Lock()
	defer outputMtx.Unlock()

	outputs[lvl] = append(outputs[lvl], output)
}

// writeToOutput writes the given message to all outputs of the level, if errors
// occur, they are collected and returned afterwards. Every output is at least
// attempted.
func writeToOutput(lvl Level, message string) {
	if !isValid(lvl) {
		return
	}

	outputMtx.RLock()
	defer outputMtx.RUnlock()
	var errs []error
	var err error
	for _, out := range outputs[lvl] {
		if isTerminal(out) {
			_, err = out.Write([]byte(message))
		} else {
			_, err = out.Write([]byte(ansi.StripString(message)))
		}
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		// just try to get word about the current error out
		for _, err := range errs {
			for _, out := range outputs[lvl] {
				if isTerminal(out) {
					_, err = out.Write([]byte(getLogLine(ERROR, fmt.Sprintf("cannot write Log: %s. Attempted to write: %s", err, message))))
				} else {
					_, err = out.Write([]byte(ansi.StripString(getLogLine(ERROR, fmt.Sprintf("cannot write Log: %s. Attempted to write: %s", err, message)))))
				}
			}
		}
	}
}

func isTerminal(file io.Writer) bool {
	var output interface{} = &file

	if _, ok := output.(os.FileInfo); !ok {
		return false
	}

	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
