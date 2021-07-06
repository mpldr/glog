package logrotation

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"math/rand"
	"reflect"
	"testing"
)

var indata []byte

func init() {
	for i := 0; i < 4*1024*1024; i++ {
		indata = append(indata, byte(rand.Intn(255)))
	}
}

func TestOptionNoCompression(t *testing.T) {
	t.Parallel()
	r := NewRotor("_test_option_no_compression", OptionNoCompression)

	inbuf := bytes.NewBuffer(indata)
	outbuf := bytes.NewBuffer([]byte{})

	err := r.compressor(outbuf, inbuf)
	if err != nil {
		t.Errorf("cannot compress data: %v", err)
	}

	if !reflect.DeepEqual(indata, outbuf.Bytes()) {
		t.Fail()
	}
}

func TestOptionGzipCompression(t *testing.T) {
	t.Parallel()
	r := NewRotor("_test_option_gzip_compression", OptionGZip, OptionMinCompression)

	inbuf := bytes.NewBuffer(indata)
	outbuf := bytes.NewBuffer([]byte{})

	err := r.compressor(outbuf, inbuf)
	if err != nil {
		t.Skipf("cannot compress data: %v", err)
	}

	decompressor, err := gzip.NewReader(outbuf)
	if err != nil {
		t.Skipf("cannot create decompressor: %v", err)
	}
	decompressbuf := bytes.NewBuffer([]byte{})

	_, err = io.Copy(decompressbuf, decompressor)
	if err != nil {
		t.Skipf("cannot decompress data: %v", err)
	}

	if !reflect.DeepEqual(indata, decompressbuf.Bytes()) {
		t.Fail()
	}
}

func TestOptionZlibCompression(t *testing.T) {
	t.Parallel()
	r := NewRotor("_test_option_zlib_compression", OptionZlib)

	inbuf := bytes.NewBuffer(indata)
	outbuf := bytes.NewBuffer([]byte{})

	err := r.compressor(outbuf, inbuf)
	if err != nil {
		t.Skipf("cannot compress data: %v", err)
	}

	decompressor, err := zlib.NewReader(outbuf)
	if err != nil {
		t.Skipf("cannot create decompressor: %v", err)
	}
	decompressbuf := bytes.NewBuffer([]byte{})

	_, err = io.Copy(decompressbuf, decompressor)
	if err != nil {
		t.Skipf("cannot decompress data: %v", err)
	}

	if !reflect.DeepEqual(indata, decompressbuf.Bytes()) {
		t.Fail()
	}
}
