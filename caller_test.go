package glog

import (
	"testing"
)

func caller(skipFrames int) string {
	return GetCaller(skipFrames)
}

func TestGetCaller(t *testing.T) {
	sc := ShortCaller
	scl := ShowCallerLine
	t.Cleanup(func() {
		ShowCallerLine = scl
		ShortCaller = sc
	})
	ShowCallerLine = false
	ShortCaller = false

	// without line, not shortened
	if caller(0) != "git.sr.ht/~poldi1405/glog.caller" {
		t.Logf("expected: %s", "git.sr.ht/~poldi1405/glog.caller")
		t.Logf("got:      %s", caller(0))
		t.Fail()
	}
	if caller(1) != "git.sr.ht/~poldi1405/glog.TestGetCaller" {
		t.Logf("expected: %s", "git.sr.ht/~poldi1405/glog.TestGetCaller")
		t.Logf("got:      %s", caller(1))
		t.Fail()
	}

	ShortCaller = true

	// without line, shortened
	if caller(0) != "glog.caller" {
		t.Logf("expected: %s", "glog.caller")
		t.Logf("got:      %s", caller(0))
		t.Fail()
	}
	if caller(1) != "glog.TestGetCaller" {
		t.Logf("expected: %s", "glog.TestGetCaller")
		t.Logf("got:      %s", caller(1))
		t.Fail()
	}

	ShowCallerLine = true

	// with line, shortened
	c := caller(0)
	if c != "caller_test.go:8" {
		t.Logf("expected: %s", "caller_test.go:8")
		t.Logf("got:      %s", c)
		t.Fail()
	}
	c = caller(1)
	if c != "caller_test.go:56" {
		t.Logf("expected: %s", "caller_test.go:56")
		t.Logf("got:      %s", c)
		t.Fail()
	}

	ShortCaller = false

	// with line, not shortened
	c = caller(0)
	if c != "git.sr.ht/~poldi1405/glog/caller_test.go:8" {
		t.Logf("expected: %s", "git.sr.ht/~poldi1405/caller_test.go:8")
		t.Logf("got:      %s", c)
		t.Fail()
	}
	c = caller(1)
	if c != "git.sr.ht/~poldi1405/glog/caller_test.go:72" {
		t.Logf("expected: %s", "git.sr.ht/~poldi1405/glog/caller_test.go:72")
		t.Logf("got:      %s", c)
		t.Fail()
	}
}
