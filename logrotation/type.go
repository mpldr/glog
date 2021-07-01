package logrotation

import (
	"compress/flate"
	"fmt"
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
func NewRotor(path string, opts ...Option) *Rotor {
	R := &Rotor{
		filepath:         path,
		Permissions:      0o600,
		MaxFileSize:      32 * 1024 * 1024,
		Retention:        2,
		KeptPercent:      5,
		compressExt:      ".gz",
		compressionLevel: flate.DefaultCompression,
		fileFlags:        os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC,
	}
	R.compressor = R.gzipCompression

	for _, o := range opts {
		switch o {
		case OptionReportErrors:
			R.errorsEnabled = true
			R.Errors = make(chan error)
		case OptionNoCompression:
			R.compressor = R.noCompression
			R.compressExt = ""
		case OptionGZip:
			R.compressor = R.gzipCompression
			R.compressExt = ".gz"
		case OptionZlib:
			R.compressor = R.zlibCompression
			R.compressExt = ".zip"
		case OptionNoSync:
			R.fileFlags &= ^os.O_SYNC
		case OptionMaxCompression:
			R.compressionLevel = flate.BestCompression
		case OptionMinCompression:
			R.compressionLevel = flate.BestSpeed
		}
	}

	return R
}

func (R *Rotor) Open() error {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	fh, err := os.OpenFile(R.filepath, R.fileFlags, R.Permissions)
	if err != nil {
		return err
	}

	fi, err := fh.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat file: %v", err)
	}

	R.size = uint64(fi.Size())
	R.file = fh
	return nil
}

func (R *Rotor) Close() error {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	return R.file.Close()
}

func (R *Rotor) SetCompressor(cf CompressorFunc, ext string) {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	R.compressor = cf
	R.compressExt = ext
}

func (R *Rotor) Write(bts []byte) (n int, err error) {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	n, err = R.file.Write(bts)

	R.size += uint64(n)

	if R.size >= R.MaxFileSize {
		err = R.rotateInsecure()
		if err != nil && R.errorsEnabled {
			R.Errors <- err
		}
		R.size = 0
	}
	return
}
