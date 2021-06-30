package logrotation

import (
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	r := NewRotor("_test_write.log")
	r.MaxFileSize = 100

	err := r.Open()
	if err != nil {
		t.Skipf("cannot open logfile: %v", err)
	}
	t.Cleanup(func() { os.Remove(r.filepath) })
	defer r.Close()

	var data []byte

	for i := 0; i < 99; i++ {
		data = append(data, byte(rand.Intn(255)))
	}

	n, err := r.Write(data)
	if err != nil {
		t.Errorf("write failed: %v", err)
	}

	if n != 99 {
		t.Errorf("incomplete write. wrote %d byte instead of %d", n, 99)
	}

	check, err := os.ReadFile(r.filepath)
	if err != nil {
		t.Errorf("created file not readable: %v", err)
	}

	if !reflect.DeepEqual(data, check) {
		t.Errorf("written data not equal to original")
	}
}

func TestWriteAndRotate(t *testing.T) {
	r := NewRotor("_test_write_and_rotate.log", OptionNoCompression)
	r.MaxFileSize = 50

	err := r.Open()
	if err != nil {
		t.Skipf("cannot open logfile: %v", err)
	}
	t.Cleanup(func() { os.Remove(r.filepath) })
	t.Cleanup(func() { os.Remove(r.filepath + ".1") })
	defer r.Close()

	var data []byte

	for i := 0; i < 99; i++ {
		data = append(data, byte(rand.Intn(255)))
	}

	n, err := r.Write(data)
	if err != nil {
		t.Errorf("write failed: %v", err)
	}

	if n != 99 {
		t.Errorf("incomplete write. wrote %d byte instead of %d", n, 99)
	}

	check, err := os.ReadFile(r.filepath + ".1")
	if err != nil {
		t.Errorf("created file not readable: %v", err)
	}

	if !reflect.DeepEqual(data, check) {
		t.Errorf("written data not equal to original")
	}

	fi, err := os.Stat(r.filepath)
	if err != nil {
		t.Errorf("cannot stat new logfile: %v", err)
	}

	if fi.Size() != 0 {
		t.Error("new logfile is not empty")
	}
}

func BenchmarkWrite(b *testing.B) {
	var data []byte

	// get 2000 MiB of random data
	for i := 0; i < 2000*1024*1024; i++ {
		data = append(data, byte(rand.Intn(255)))
	}

	b.Run("write", func(b *testing.B) {
		b.StopTimer()
		r := NewRotor("_bench_write_sync.log")
		b.StartTimer()

		r.Write(data)

		b.StopTimer()
		r.Close()
		os.Remove("_bench_write_sync.log")
	})

	b.Run("write-no-sync", func(b *testing.B) {
		b.StopTimer()
		r := NewRotor("_bench_write_nosync.log", OptionNoSync)
		b.StartTimer()

		r.Write(data)

		b.StopTimer()
		r.Close()
		os.Remove("_bench_write_nosync.log")
	})
}
