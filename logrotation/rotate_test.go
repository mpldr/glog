package logrotation

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestRotateNoLogFile(t *testing.T) {
	// no existing file, create one
	r := NewRotor("_test_no_file_exists.log")

	r.Rotate()
	t.Cleanup(func() { os.Remove("_test_no_file_exists.log") })

	if !fileExists("_test_no_file_exists.log") {
		t.Fail()
	}
}

func TestRotateRetention(t *testing.T) {
	r := NewRotor("_test_rotation_retention.log", OptionNoCompression)

	err := os.WriteFile(r.filepath, []byte{0}, 0o600)
	if err != nil {
		t.Skipf("cannot create testfile: %v", err)
	}
	t.Cleanup(func() { os.Remove(r.filepath) })

	for i := 1; i < 3; i++ {
		filename := fmt.Sprintf(r.filepath+".%d", i)
		err := os.WriteFile(filename, []byte{byte(i)}, 0o600)
		if err != nil {
			t.Skipf("cannot create testfile: %v", err)
		}
		t.Cleanup(func() { os.Remove(filename) })
	}

	err = r.Open()
	if err != nil {
		t.Skipf("unable to open logfile: %v", err)
	}

	err = r.Rotate()
	if err != nil {
		t.Errorf("rotating failed: %v", err)
	}

	content, err := os.ReadFile(r.filepath)
	if err != nil || len(content) != 0 {
		if len(content) != 0 {
			err = fmt.Errorf("new file is not empty. contains: %v", content)
		}

		t.Errorf("cannot verify newly created logfile: %v", err)
	}

	content, err = os.ReadFile(r.filepath + ".1")
	if err != nil || len(content) != 1 {
		if len(content) == 0 {
			err = errors.New("rotated file is empty")
		} else if content[0] != 0 {
			err = fmt.Errorf("wrong content. expected %v, but got %v", []byte{0}, content)
		}

		t.Errorf("cannot verify first rotated logfile: %v", err)
	}

	content, err = os.ReadFile(r.filepath + ".2")
	if err != nil {
		if len(content) == 0 {
			err = errors.New("rotated file is empty")
		} else if content[0] != 1 {
			err = fmt.Errorf("wrong content. expected %v, but got %v", []byte{1}, content)
		}

		t.Errorf("cannot verify second rotated logfile: %v", err)
	}
}

func TestRotateGzip(t *testing.T) {
	r := NewRotor("_test_rotation_retention.log", OptionGZip)
	r.Retention = 2

	err := os.WriteFile(r.filepath, []byte{0}, 0o600)
	if err != nil {
		t.Skipf("cannot create testfile: %v", err)
	}
	t.Cleanup(func() { os.Remove(r.filepath) })

	for i := 1; i < 3; i++ {
		buf := bytes.NewBuffer([]byte{})
		trg := bytes.NewBuffer([]byte{})
		buf.Write([]byte{byte(i)})
		r.gzipCompression(trg, buf)

		filename := fmt.Sprintf(r.filepath+".%d.gz", i)
		err := os.WriteFile(filename, trg.Bytes(), 0o600)
		if err != nil {
			t.Skipf("cannot create testfile: %v", err)
		}
		t.Cleanup(func() { os.Remove(filename) })
	}

	err = r.Open()
	if err != nil {
		t.Skipf("unable to open logfile: %v", err)
	}

	err = r.Rotate()
	if err != nil {
		t.Errorf("rotating failed: %v", err)
	}

	content, err := os.ReadFile(r.filepath)
	if err != nil || len(content) != 0 {
		if len(content) != 0 {
			err = fmt.Errorf("new file is not empty. contains: %v", content)
		}

		t.Errorf("cannot verify newly created logfile: %v", err)
	}

	content, err = os.ReadFile(r.filepath + ".1.gz")
	if err != nil || reflect.DeepEqual(content, []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x33, 0x00, 0x00, 0x21, 0xdf, 0xdb, 0xf4, 0x01, 0x00, 0x00, 0x00}) {
		if len(content) == 0 {
			err = errors.New("rotated file is empty")
		} else if content[0] != 0 {
			err = fmt.Errorf("wrong content. expected %v, but got %v", []byte{0}, content)
		}

		t.Errorf("cannot verify first rotated logfile: %v", err)
	}

	content, err = os.ReadFile(r.filepath + ".2.gz")
	if err != nil || reflect.DeepEqual(content, []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x33, 0x04, 0x00, 0xb7, 0xef, 0xdc, 0x83, 0x01, 0x00, 0x00, 0x00}) {
		if len(content) == 0 {
			err = errors.New("rotated file is empty")
		} else if content[0] != 0 {
			err = fmt.Errorf("wrong content. expected %v, but got %v", []byte{0}, content)
		}

		t.Errorf("cannot verify second rotated logfile: %v", err)
	}
}
