package logrotation

import (
	"io"
	"os"
	"sync"
)

var _ io.Writer = &Rotor{}

type Rotor struct {
	// filepath is the path of the logfile
	filepath string
	// Permissions indicate unix file permissions to apply. By default it's
	// 0o600
	Permissions os.FileMode
	// MaxFileSize is after what time the file is truncated. Default: 32 MiB
	MaxFileSize uint
	// Retention is how many "old" logs are kept. Default: 2. This means you
	// have 3 files in total: file.log, file.log.1.gz, file.log.2.gz
	Retention int
	// KeptPercent is how many % (of lines) are kept when using flowing-mode
	// This value is only used if retention is 0 or below. The last x% of
	// lines are kept and mark the starting point of the new log. Default: 5
	KeptPercent int
	// Compress indicates whether old logfiles are to be compressed. This
	// makes rotating the log slightly longer while at the same time
	// drastically shrinking their filesize.
	Compress bool

	file    *os.File
	fileMtx sync.Mutex
}

func NewRotor(path string) *Rotor {
	return &Rotor{
		filepath:    path,
		Permissions: 0o600,
		MaxFileSize: 32 * 1024 * 1024,
		Retention:   2,
		KeptPercent: 5,
		Compress:    true,
	}
}

func (R *Rotor) Write(bytes []byte) (int, error) {
}
