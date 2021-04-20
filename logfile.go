package glog

import (
	"os"
)

// AddLogFile opens the specified files and adds it to the specified levels
func AddLogFile(filepath string, lvls ...Level) (*os.File, error) {
	fh, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_SYNC|os.O_CREATE, 0o600)
	if err != nil {
		metalog("error opening logfile:", err)
		return nil, err
	}

	for _, lvl := range lvls {
		metalog("adding", filepath, "to level", lvl)
		AddOutput(lvl, fh)
	}
	return fh, nil
}
