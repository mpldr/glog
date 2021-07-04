package logrotation

import (
	"compress/flate"
	"fmt"
	"io"
	"os"
	"sync"
)

var _ io.Writer = &Rotor{}

// Rotor is the struct containing all Settings, file handlers, and functions
// required for logrotation.
type Rotor struct {
	// filepath is the path of the logfile
	filepath string
	// Permissions indicate unix file permissions to apply. By default it's
	// 0o600
	Permissions os.FileMode
	// fileFlags contains fileFlags that are applied to logfiles on opening
	fileFlags int
	// MaxFileSize is after what time the file is truncated. Default: 32 MiB
	MaxFileSize uint64
	// Retention is how many "old" logs are kept. Default: 2. This means you
	// have 3 files in total: file.log, file.log.1.gz, file.log.2.gz
	Retention int
	// KeptPercent is how many % (of lines) are kept when using flowing-mode
	// This value is only used if retention is 0 or below. The last x% of
	// lines are kept and mark the starting point of the new log. Default: 5
	//
	// Please note that this is an approximation. The percent-position will
	// sought and everything _after_ the next '\n' will be kept for the
	// rotated file.
	KeptPercent int
	// compress indicates whether old logfiles are to be compressed. This
	// makes rotating the log slightly longer while at the same time
	// drastically shrinking their filesize.
	compressor CompressorFunc
	// compressExt contains the file extension for the specified compression
	// method
	compressExt string
	// compressionLevel contains the compression level to use with gzip/zlib
	// compression. It has no effect on custom compressors.
	compressionLevel int
	// Errors is an unbuffered channel to get notified of errors while
	// rotating. It is only enabled when the Rotor was created with
	// OptionReportErrors
	Errors        chan error
	errorsEnabled bool

	file    *os.File
	fileMtx sync.Mutex
	size    uint64
}

// NewRotor generates a new Rotor with the given options overwriting the
// defaults. The defaults are:
// Permissions: 600,
// MaxFileSize: 32 MiB,
// Retention:   2,
// KeptPercent: 5,
// Compression: gzip (Default Compression)
func NewRotor(path string, opts ...uint8) *Rotor {
	r := &Rotor{
		filepath:         path,
		Permissions:      0o600,
		MaxFileSize:      32 * 1024 * 1024,
		Retention:        2,
		KeptPercent:      5,
		compressExt:      ".gz",
		compressionLevel: flate.DefaultCompression,
		fileFlags:        os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC,
	}
	r.compressor = r.gzipCompression

	for _, o := range opts {
		switch o {
		case OptionReportErrors:
			r.errorsEnabled = true
			r.Errors = make(chan error)
		case OptionNoCompression:
			r.compressor = r.noCompression
			r.compressExt = ""
		case OptionGZip:
			r.compressor = r.gzipCompression
			r.compressExt = ".gz"
		case OptionZlib:
			r.compressor = r.zlibCompression
			r.compressExt = ".zip"
		case OptionNoSync:
			r.fileFlags &= ^os.O_SYNC
		case OptionMaxCompression:
			r.compressionLevel = flate.BestCompression
		case OptionMinCompression:
			r.compressionLevel = flate.BestSpeed
		}
	}

	return r
}

func (r *Rotor) isSync() bool {
	return r.fileFlags&os.O_SYNC != 0
}

// Open opens (or creates) the logfile specified when setting up the Rotor.
// Closing this file is the responsibility of the user.
func (r *Rotor) Open() error {
	r.fileMtx.Lock()
	defer r.fileMtx.Unlock()

	fh, err := os.OpenFile(r.filepath, r.fileFlags, r.Permissions)
	if err != nil {
		return err
	}

	fi, err := fh.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat file: %v", err)
	}

	r.size = uint64(fi.Size())
	r.file = fh

	return nil
}

// Close closes the underlying filedescriptor
func (r *Rotor) Close() error {
	r.fileMtx.Lock()
	defer r.fileMtx.Unlock()

	return r.file.Close()
}

// SetCompressor allows setting a custom compression method for use when rotating
func (r *Rotor) SetCompressor(cf CompressorFunc, ext string) {
	r.fileMtx.Lock()
	defer r.fileMtx.Unlock()

	r.compressor = cf
	r.compressExt = ext
}

func (r *Rotor) Write(bts []byte) (n int, err error) {
	r.fileMtx.Lock()
	defer r.fileMtx.Unlock()

	n, err = r.file.Write(bts)

	r.size += uint64(n)

	if r.size >= r.MaxFileSize {
		err = r.rotateInsecure()
		if err != nil && r.errorsEnabled {
			r.Errors <- err
		}
		r.size = 0
	}

	return
}
