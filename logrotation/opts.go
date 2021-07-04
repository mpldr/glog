package logrotation

const (
	// OptionReportErrors enables the error reporting channel that informs
	// about errors that happen in the background and are otherwise ignored
	OptionReportErrors = iota
	// OptionNoCompression disabled the compression of Archived logs.
	OptionNoCompression
	// OptionGZip enables compression using gzip
	OptionGZip
	// OptionZlib enables compression using zip
	OptionZlib
	// OptionNoSync disables the synchronous flag when opening logfiles this
	// can yield double speed but may lead to data loss if the system is
	// turned off before a flush to disk happened
	OptionNoSync
	// OptionMaxCompression sets the compression level to maximum, thereby
	// taking longer but achieving better compression
	OptionMaxCompression
	// OptionMinCompression sets the compression level to minimum, thereby
	// speeding up compression drastically while also increasing the rotated
	// filesize substantially
	OptionMinCompression
)
