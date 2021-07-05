package logrotation_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"git.sr.ht/~poldi1405/glog/logrotation"
)

var PS = string(os.PathSeparator)

func TestWriteNoCompressionBB(t *testing.T) {
	t.Parallel()

	err := os.Mkdir("_test_write_no_compression_bb", 0o700)
	if err != nil {
		t.Skipf("cannot setup test conditions: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll("_test_write_no_compression_bb") })

	r := logrotation.NewRotor("_test_write_no_compression_bb"+PS+"logfile.log", logrotation.OptionNoCompression)
	r.MaxFileSize = 128
	err = r.Open()
	if err != nil {
		t.Errorf("failed to open file: %v", err)
		return
	}

	for i := byte(0); i < 255; i++ {
		_, err = r.Write([]byte{i})
		if err != nil {
			t.Errorf("failed to write to file: %v", err)
			return
		}
	}

	cnt, err := os.ReadFile("_test_write_no_compression_bb" + PS + "logfile.log")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
		return
	}

	for i := byte(128); i < 255; i++ {
		if cnt[i-128] != i {
			t.Errorf("wrong content in logfile at index: %d", i-128)
			return
		}
	}

	cnt, err = os.ReadFile("_test_write_no_compression_bb" + PS + "logfile.log.1")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
		return
	}

	for i := byte(0); i < 128; i++ {
		if cnt[i] != i {
			t.Errorf("wrong content in logfile at index: %d", i)
			return
		}
	}
}

func TestWriteGzipCompressionBB(t *testing.T) {
	t.Parallel()

	err := os.Mkdir("_test_write_gzip_compression_bb", 0o700)
	if err != nil {
		t.Skipf("cannot setup test conditions: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll("_test_write_gzip_compression_bb") })

	r := logrotation.NewRotor("_test_write_gzip_compression_bb" + PS + "logfile.log")
	r.MaxFileSize = 128
	err = r.Open()
	if err != nil {
		t.Errorf("failed to open file: %v", err)
		return
	}

	for i := byte(0); i < 255; i++ {
		_, err = r.Write([]byte{i})
		if err != nil {
			t.Errorf("failed to write to file: %v", err)
			return
		}
	}

	cnt, err := os.ReadFile("_test_write_gzip_compression_bb" + PS + "logfile.log")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
		return
	}

	for i := byte(128); i < 255; i++ {
		if cnt[i-128] != i {
			t.Errorf("wrong content in logfile at index: %d", i-128)
			return
		}
	}

	cmdoutput := bytes.NewBuffer([]byte{})

	cmd := exec.Command("gzip", "-d", "-c", "_test_write_gzip_compression_bb"+PS+"logfile.log.1.gz")
	cmd.Stdout = cmdoutput
	if err := cmd.Run(); err != nil {
		if err == exec.ErrNotFound {
			t.Skipf("unable to run decompression: %v", err)
		}
		t.Errorf("error while decompressing: %v", err)
		return
	}

	for i := byte(0); i < 128; i++ {
		if cmdoutput.Bytes()[i] != i {
			t.Errorf("wrong content in logfile at index: %d", i)
			return
		}
	}
}

func TestWriteZlibCompressionBB(t *testing.T) {
	t.Parallel()

	err := os.Mkdir("_test_write_zlib_compression_bb", 0o700)
	if err != nil {
		t.Skipf("cannot setup test conditions: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll("_test_write_zlib_compression_bb") })

	r := logrotation.NewRotor("_test_write_zlib_compression_bb"+PS+"logfile.log", logrotation.OptionZlib)
	r.MaxFileSize = 128
	err = r.Open()
	if err != nil {
		t.Errorf("failed to open file: %v", err)
		return
	}

	for i := byte(0); i < 255; i++ {
		_, err = r.Write([]byte{i})
		if err != nil {
			t.Errorf("failed to write to file: %v", err)
			return
		}
	}

	cnt, err := os.ReadFile("_test_write_zlib_compression_bb" + PS + "logfile.log")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
		return
	}

	for i := byte(128); i < 255; i++ {
		if cnt[i-128] != i {
			t.Errorf("wrong content in logfile at index: %d", i-128)
			return
		}
	}

	fileContent, err := os.ReadFile("_test_write_zlib_compression_bb" + PS + "logfile.log.1.zip")

	cmdoutput := bytes.NewBuffer([]byte{})
	cmdinput := bytes.NewBuffer(fileContent)

	cmd := exec.Command("zlib-flate", "-uncompress")
	cmd.Stdin = cmdinput
	cmd.Stdout = cmdoutput
	if err := cmd.Run(); err != nil {
		if err == exec.ErrNotFound {
			t.Skipf("unable to run decompression: %v", err)
		}
		t.Errorf("error while decompressing: %v", err)
		t.Logf("stderr: %v", cmd.Stderr)
		return
	}

	for i := byte(0); i < 128; i++ {
		if cmdoutput.Bytes()[i] != i {
			t.Errorf("wrong content in logfile at index: %d", i)
			return
		}
	}
}
