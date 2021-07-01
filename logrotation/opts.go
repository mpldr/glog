package logrotation

type Option uint8

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
)
