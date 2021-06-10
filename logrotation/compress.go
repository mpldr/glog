package logrotation

import (
	"compress/gzip"
	"fmt"
	"io"
)

func deflate(output io.Writer, input io.Reader) error {
	zw, err := gzip.NewWriterLevel(output, 5)
	if err != nil {
		return fmt.Errorf("cannot create deflate writer for log-rotation: %w", err)
	}

	_, err = io.Copy(zw, input)
	if err != nil {
		return fmt.Errorf("unable to compress file: %w", err)
	}

	if err := zw.Close(); err != nil {
		return fmt.Errorf("cannot close deflate writer: %w", err)
	}
	return nil
}
