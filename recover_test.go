package glog

import (
	"os"
	"testing"
)

func TestPanicHandler(t *testing.T) {
	msg := "testpanic"

	defer func() {
		r := recover()
		if r == nil {
			t.FailNow()
		}
		if r != msg {
			t.Fail()
		}
	}()
	defer PanicHandler()

	panic(msg)
}

func TestPanicHandlerNoPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.FailNow()
		}
	}()
	defer PanicHandler()
	t.Cleanup(func() { os.Remove("panic.log") })
}

func TestPanicHandlerFailWrite(t *testing.T) {
	err := os.Mkdir("panic.log", 0o700)
	if err != nil {
		t.Logf("cannot create panic.log: %v", err)
		t.SkipNow()
	}
	t.Cleanup(func() { os.Remove("panic.log") })

	defer func() {
		r := recover()
		if r != nil {
			t.FailNow()
		}
	}()
	defer PanicHandler()
}
