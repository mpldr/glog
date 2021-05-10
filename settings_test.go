package glog

import (
	"testing"
)

func TestSetShowCaller(t *testing.T) {
	tmp := showCaller
	t.Cleanup(func() {
		showCaller = tmp
	})

	SetShowCaller(TRACE, false)
	if showCaller[TRACE] {
		t.Fail()
	}

	SetShowCaller(WARNING, true)
	if !showCaller[WARNING] {
		t.Fail()
	}

	SetShowCaller(FATAL, true)
	if !showCaller[FATAL] {
		t.Fail()
	}
}

func TestSetShowCallerInvalid(t *testing.T) {
	tmp := showCaller

	SetShowCaller(42, false)
	for k, v := range showCaller {
		if tmp[k] != v {
			t.Fail()
		}
	}

	// reset
	showCaller = tmp
}
