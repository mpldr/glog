package glog

import (
	"bufio"
	"os"
)

func AddLogFile(filepath string, lvls ...Level) (*os.File, error) {
	fh, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_SYNC|os.O_CREATE, 0o600)
	if err != nil {
		return nil, err
	}

	bufw := bufio.NewWriter(fh)

	for _, lvl := range lvls {
		AddOutput(lvl, bufw)
	}
	return nil, nil
}
