package logrotation

import (
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
)

type CompressorFunc func(io.Writer, io.Reader) error

func noCompression(output io.Writer, input io.Reader) error {
	_, err := io.Copy(output, input)
	if err != nil {
		return fmt.Errorf("unable to copy file: %w", err)
	}

	return nil
}

func gzipCompression(output io.Writer, input io.Reader) error {
	zw, err := gzip.NewWriterLevel(output, flate.DefaultCompression)
	if err != nil {
		return fmt.Errorf("cannot create deflate writer for log-rotation: %w", err)
	}
	defer zw.Close()

	_, err = io.Copy(zw, input)
	if err != nil {
		return fmt.Errorf("unable to compress file: %w", err)
	}

	return nil
}

func zlibCompression(output io.Writer, input io.Reader) error {
	zw, err := zlib.NewWriterLevel(output, flate.DefaultCompression)
	if err != nil {
		return fmt.Errorf("cannot create zlib writer for log-rotation: %w", err)
	}
	defer zw.Close()

	_, err = io.Copy(zw, input)
	if err != nil {
		return fmt.Errorf("unable to compress file: %w", err)
	}

	return nil
}
