package logrotation

import (
	"compress/flate"
	"os"
	"testing"
)

func TestOptionNoSync(t *testing.T) {
	r := NewRotor("_test_option_no_sync.log", OptionNoSync)

	if r.fileFlags&os.O_SYNC != 0 {
		t.Fail()
	}
}

func TestOptionMaxCompression(t *testing.T) {
	r := NewRotor("_test_option_max_compression.log", OptionMaxCompression)

	if r.compressionLevel != flate.BestCompression {
		t.Fail()
	}
}

func TestOptionMinCompression(t *testing.T) {
	r := NewRotor("_test_option_min_compression.log", OptionMinCompression)

	if r.compressionLevel != flate.BestSpeed {
		t.Fail()
	}
}
